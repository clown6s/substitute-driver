package middleware

import (
	"context"
	"net/http"

	"substitute-driver/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// SetupErrorHandler 全局错误处理
// 替换 httpx 默认的 ErrorCtx 处理逻辑，将 gRPC status error 转为结构化 JSON 响应
func SetupErrorHandler() {
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		st, ok := status.FromError(err)
		if ok {
			return http.StatusOK, response.Body{
				Code:    int(st.Code()),
				Message: st.Message(),
			}
		}
		return http.StatusInternalServerError, response.Body{
			Code:    500,
			Message: "服务内部错误",
		}
	})
}
