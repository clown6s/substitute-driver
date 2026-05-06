package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// Body 统一响应结构
type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response 带 data 的统一响应结构
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Success 成功响应
func Success(w http.ResponseWriter, resp any) {
	httpx.OkJson(w, Response{
		Code:    0,
		Message: "ok",
		Data:    resp,
	})
}

// Ok 成功响应（无 data）
func Ok(w http.ResponseWriter) {
	httpx.OkJson(w, Response{
		Code:    0,
		Message: "ok",
	})
}

// Error 错误响应
func Error(w http.ResponseWriter, code int, msg string) {
	httpx.WriteJson(w, http.StatusOK, Body{
		Code:    code,
		Message: msg,
	})
}

// ErrorDefault 默认错误响应（code=500）
func ErrorDefault(w http.ResponseWriter, msg string) {
	httpx.WriteJson(w, http.StatusInternalServerError, Body{
		Code:    500,
		Message: msg,
	})
}

// ErrorFromGrpc 从 gRPC 错误中提取 code 和 message，返回统一响应
func ErrorFromGrpc(w http.ResponseWriter, err error) {
	st, ok := status.FromError(err)
	if !ok {
		ErrorDefault(w, "服务内部错误")
		return
	}
	Error(w, int(st.Code()), st.Message())
}
