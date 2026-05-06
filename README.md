# substitute-driver 代驾系统

基于 [go-zero](https://github.com/zeromicro/go-zero) 的代驾微服务项目，采用 gRPC 通信、MySQL 持久化、Etcd 服务发现、Redis 缓存。

## 技术栈

| 组件 | 版本 | 说明 |
|---|---|---|
| Go | 1.26+ | 开发语言 |
| go-zero | v1.10.1 | 微服务框架 |
| goctl | v1.10.1 | 代码生成工具 |
| gRPC | protobuf v3 | 微服务间通信 |
| MySQL | 8.0+ | 数据持久化 |
| Redis | - | CRUD 缓存 |
| Etcd | - | 服务注册与发现 |
| protoc | v34.1 | Proto 编译器 |

## 项目结构

```
substitute-driver/
├── go.work                          # Go workspace 多模块管理
├── substitute-driver.sql            # 建库建表 DDL
│
├── common/                          # 公共模块
│   ├── errorx/errorx.go             # 统一 gRPC 错误码
│   └── model/                       # goctl 生成的 MySQL CRUD
│       ├── userModel.go             # 用户 Model 接口（可扩展）
│       ├── userModel_gen.go         # 用户 Model 实现（自动生成）
│       ├── driverModel.go           # 司机 Model 接口（可扩展）
│       ├── driverModel_gen.go       # 司机 Model 实现（自动生成）
│       └── vars.go
│
├── gateway/                         # HTTP 网关 (:8888)
│   ├── gateway.api                  # API 路由定义
│   ├── gateway.go                   # 入口
│   ├── etc/gateway-api.yaml         # 配置（Etcd 地址）
│   └── internal/
│       ├── config/config.go
│       ├── svc/serviceContext.go    # 注入三个 RPC 客户端
│       ├── handler/                 # HTTP Handler（goctl 生成）
│       │   ├── user/
│       │   ├── driver/
│       │   └── mapmodule/
│       ├── logic/                   # 业务逻辑（转发到 RPC）
│       │   ├── user/
│       │   ├── driver/
│       │   └── mapmodule/
│       └── types/types.go
│
├── user/                            # 用户服务 (:8080)
│   ├── user.proto                   # 服务 Proto 定义
│   ├── user.go                      # 入口
│   ├── etc/user.yaml                # 配置
│   ├── client/user/                 # gRPC 客户端
│   └── internal/
│       ├── config/config.go
│       ├── svc/serviceContext.go
│       ├── server/user/             # gRPC Server
│       └── logic/user/              # 业务逻辑
│           ├── registerLogic.go     # 用户注册（MD5 密码加密）
│           ├── loginLogic.go        # 用户登录（密码校验）
│           ├── getUserInfoLogic.go  # 查询用户信息
│           └── updateUserLogic.go   # 更新用户信息
│
├── driver/                          # 司机服务 (:8081)
│   ├── driver.proto
│   ├── driver.go
│   ├── etc/driver.yaml
│   ├── client/driver/
│   └── internal/
│       ├── config/config.go
│       ├── svc/serviceContext.go
│       ├── server/driver/
│       └── logic/driver/
│           ├── registerLogic.go         # 司机注册
│           ├── getDriverInfoLogic.go    # 查询司机信息
│           ├── updateLocationLogic.go   # 更新司机位置
│           ├── updateStatusLogic.go     # 更新司机状态
│           └── getNearbyDriversLogic.go # 查询附近司机（骨架）
│
└── map/                             # 地图服务 (:8082)
    ├── map.proto
    ├── mapservice.go
    ├── etc/mapservice.yaml
    ├── client/mapservice/
    └── internal/
        ├── config/config.go
        ├── svc/serviceContext.go
        ├── server/mapservice/
        └── logic/mapservice/
            ├── getDistanceLogic.go      # 计算距离（Haversine）
            ├── getRouteLogic.go         # 路线规划（mock）
            ├── geocodeLogic.go          # 地理编码（mock）
            └── reverseGeocodeLogic.go   # 逆地理编码（mock）
```

## HTTP API（Gateway :8888）

Gateway 将所有 HTTP 请求转发到对应的 gRPC 服务，对外统一暴露 REST API。

### 用户模块 `/api/v1/user`

| 方法 | 路径 | 请求体 / 参数 | 响应 | 说明 |
|---|---|---|---|---|
| POST | `/api/v1/user/register` | `{username, password, phone}` | `{id}` | 用户注册，密码 MD5 加密 |
| POST | `/api/v1/user/login` | `{username, password}` | `{id, token}` | 用户登录，返回 token |
| GET | `/api/v1/user/:id` | path: `id` | `{id, username, phone, avatar, createdAt}` | 查询用户信息 |
| PUT | `/api/v1/user/:id` | path: `id`，body: `{username?, phone?, avatar?}` | - | 更新用户信息（按需更新） |

### 司机模块 `/api/v1/driver`

| 方法 | 路径 | 请求体 / 参数 | 响应 | 说明 |
|---|---|---|---|---|
| POST | `/api/v1/driver/register` | `{name, phone, licenseNo, carNo}` | `{id}` | 司机注册，默认休息状态 |
| GET | `/api/v1/driver/:id` | path: `id` | `{driver:{id,name,phone,licenseNo,carNo,status,longitude,latitude}}` | 查询司机信息 |
| PUT | `/api/v1/driver/location` | `{id, longitude, latitude}` | - | 更新司机实时位置 |
| PUT | `/api/v1/driver/status` | `{id, status}` | - | 更新状态（0休息/1接单/2服务中） |
| GET | `/api/v1/driver/nearby` | query: `longitude, latitude, radius?` | `{drivers:[]}` | 查询附近司机（骨架） |

### 地图模块 `/api/v1/map`

| 方法 | 路径 | 请求体 | 响应 | 说明 |
|---|---|---|---|---|
| POST | `/api/v1/map/distance` | `{fromLng, fromLat, toLng, toLat}` | `{distance(km), duration(s)}` | Haversine 直线距离 |
| POST | `/api/v1/map/route` | `{fromLng, fromLat, toLng, toLat}` | `{distance, duration, polyline[]}` | 路线规划（mock） |
| POST | `/api/v1/map/geocode` | `{address}` | `{lng, lat, formattedAddress}` | 地址转坐标（mock） |
| POST | `/api/v1/map/reverse-geocode` | `{lng, lat}` | `{formattedAddress}` | 坐标转地址（mock） |

## gRPC 接口

### user（用户服务 :8080）

| RPC 方法 | 请求字段 | 响应字段 | 说明 |
|---|---|---|---|
| `Register` | username, password, phone | id | 用户注册，密码 MD5 加密 |
| `Login` | username, password | id, token | 用户登录，返回 mock token |
| `GetUserInfo` | id | id, username, phone, avatar, createdAt | 查询用户信息 |
| `UpdateUser` | id, username, phone, avatar | - | 更新用户信息（按需更新） |

### driver（司机服务 :8081）

| RPC 方法 | 请求字段 | 响应字段 | 说明 |
|---|---|---|---|
| `Register` | name, phone, licenseNo, carNo | id | 司机注册，默认休息状态 |
| `GetDriverInfo` | id | id, name, phone, licenseNo, carNo, status, lng, lat | 查询司机信息 |
| `UpdateLocation` | id, longitude, latitude | - | 更新司机实时位置 |
| `UpdateStatus` | id, status | - | 更新状态（0休息/1接单/2服务中） |
| `GetNearbyDrivers` | longitude, latitude, radius | drivers[] | 查询附近司机（骨架） |

### map（地图服务 :8082）

| RPC 方法 | 请求字段 | 响应字段 | 说明 |
|---|---|---|---|
| `GetDistance` | fromLng, fromLat, toLng, toLat | distance(km), duration(s) | Haversine 直线距离 |
| `GetRoute` | fromLng, fromLat, toLng, toLat | distance, duration, polyline[] | 路线规划（mock） |
| `Geocode` | address | lng, lat, formattedAddress | 地址转坐标（mock） |
| `ReverseGeocode` | lng, lat | formattedAddress | 坐标转地址（mock） |

## 数据库设计

### user 表

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint PK AUTO_INCREMENT | 用户 ID |
| username | varchar(64) UNIQUE | 用户名 |
| password | varchar(255) | MD5 加密密码 |
| phone | varchar(20) UNIQUE | 手机号 |
| avatar | varchar(512) | 头像 URL |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

### driver 表

| 字段 | 类型 | 说明 |
|---|---|---|
| id | bigint PK AUTO_INCREMENT | 司机 ID |
| name | varchar(64) | 姓名 |
| phone | varchar(20) UNIQUE | 手机号 |
| license_no | varchar(32) | 驾照号 |
| car_no | varchar(16) | 车牌号 |
| status | tinyint DEFAULT 0 | 状态：0-休息 1-接单中 2-服务中 |
| longitude | double | 经度 |
| latitude | double | 纬度 |
| created_at | timestamp | 创建时间 |
| updated_at | timestamp | 更新时间 |

## 快速开始

### 环境要求

- Go 1.26+
- MySQL 8.0+
- Redis
- Etcd
- goctl `go install github.com/zeromicro/go-zero/tools/goctl@latest`
- protoc + protoc-gen-go + protoc-gen-go-grpc

### 1. 启动基础设施（Docker）

```bash
# Etcd
docker run -d --name etcd \
  -p 2379:2379 -p 2380:2380 \
  -e ALLOW_NONE_AUTHENTICATION=yes \
  bitnami/etcd:latest

# MySQL
docker run -d --name mysql \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=your_password \
  -e MYSQL_DATABASE=substitute_driver \
  mysql:8.0

# Redis
docker run -d --name redis -p 6379:6379 redis:latest
```

### 2. 初始化数据库

```bash
mysql -u root -p substitute_driver < substitute-driver.sql
```

### 3. 修改配置

编辑各服务 `etc/*.yaml`，修改 MySQL 密码、Redis 地址、Etcd 地址：

```yaml
# user/etc/user.yaml 示例
Name: user.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: user.rpc
Mysql:
  DataSource: root:your_password@tcp(127.0.0.1:3306)/substitute_driver?charset=utf8mb4&parseTime=true&loc=Local
CacheRedis:
  - Host: 127.0.0.1:6379
```

> **注意**：`gateway/etc/gateway-api.yaml` 中三个 RPC 的 Etcd Hosts 同样需要改为 `127.0.0.1:2379`。

### 4. 启动服务（顺序：RPC 服务先于 Gateway）

```bash
# 终端 1 — 用户服务
cd user && go run user.go -f etc/user.yaml

# 终端 2 — 司机服务
cd driver && go run driver.go -f etc/driver.yaml

# 终端 3 — 地图服务
cd map && go run mapservice.go -f etc/mapservice.yaml

# 终端 4 — HTTP 网关（三个 RPC 服务全部注册到 Etcd 后再启动）
cd gateway && go run gateway.go -f etc/gateway-api.yaml
```

### 5. 接口测试（curl）

```bash
# 注册用户
curl -X POST http://localhost:8888/api/v1/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","phone":"13800000000"}'

# 登录
curl -X POST http://localhost:8888/api/v1/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'

# 注册司机
curl -X POST http://localhost:8888/api/v1/driver/register \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","phone":"13900000001","licenseNo":"A12345","carNo":"云A88888"}'

# 计算距离（北京 → 上海）
curl -X POST http://localhost:8888/api/v1/map/distance \
  -H "Content-Type: application/json" \
  -d '{"fromLng":116.397,"fromLat":39.908,"toLng":121.473,"toLat":31.230}'
```

### 6. 接口测试（gRPCurl，直连 RPC 服务）

```bash
# 注册用户
grpcurl -plaintext -d '{"username":"test","password":"123456","phone":"13800000000"}' \
  127.0.0.1:8080 user.user/Register

# 登录
grpcurl -plaintext -d '{"username":"test","password":"123456"}' \
  127.0.0.1:8080 user.user/Login

# 计算距离
grpcurl -plaintext -d '{"fromLng":116.397,"fromLat":39.908,"toLng":121.473,"toLat":31.230}' \
  127.0.0.1:8082 mapservice.mapservice/GetDistance
```

## 代码生成

项目使用 goctl 生成服务骨架和 Model，生成命令如下：

### 生成 gRPC 服务

```bash
# 在各服务目录下执行
cd user && goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero -m
cd driver && goctl rpc protoc driver.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero -m
cd map && goctl rpc protoc map.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero -m
```

### 生成 HTTP Gateway

```bash
cd gateway && goctl api go -api gateway.api -dir . --style=goZero
```

### 生成 MySQL Model

```bash
goctl model mysql ddl --src substitute-driver.sql --dir common/model --style=goZero -c
```

## 错误码

| 错误码 | 说明 |
|---|---|
| 10001 | 用户已存在 |
| 10002 | 用户不存在 |
| 10003 | 密码错误 |
| 10004 | 参数错误 |
| 10005 | 数据库错误 |
| 10006 | 服务内部错误 |
| 20001 | 司机已存在 |
| 20002 | 司机不存在 |
| 20003 | 司机不在线 |
| 30001 | 地图服务异常 |

## 服务依赖关系

```
Client (HTTP)
     │
     ▼
 Gateway :8888  ──── Etcd :2379 ────┐
     │                              │
     ├── UserRpc ──────────── user :8080 ──── MySQL / Redis
     ├── DriverRpc ─────────── driver :8081 ── MySQL / Redis
     └── MapRpc ────────────── map :8082
```

> Gateway 通过 Etcd 服务发现找到各 RPC 服务地址，**必须先启动 Etcd 和三个 RPC 服务，再启动 Gateway**。

## 待办

- [x] HTTP API Gateway（gateway 服务，:8888）
- [ ] 集成 JWT 认证（登录签发真实 token，Gateway 鉴权中间件）
- [ ] map 服务接入高德/腾讯地图 SDK
- [ ] driver.GetNearbyDrivers 接入 Redis GEO 实现附近司机搜索
- [ ] 添加订单服务（order）
- [ ] 补充单元测试
