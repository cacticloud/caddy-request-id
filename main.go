package caddyrequestid

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
	caddy.RegisterModule(RequestID{})
}

type RequestID struct{}

// CaddyModule 返回 Caddy 模块的信息。
func (RequestID) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.request_id",
		New: func() caddy.Module { return new(RequestID) },
	}
}

func (reqID RequestID) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	uniqueID := generateUniqueID()         // 生成唯一 ID
	r.Header.Set("X-Request-ID", uniqueID) // 设置请求头
	return next.ServeHTTP(w, r)            // 继续执行下一个处理器
}

// generateUniqueID 使用当前时间和一个随机数生成一个唯一的 ID。
func generateUniqueID() string {
	now := time.Now()
	return fmt.Sprintf("%d%06d", now.UnixNano(), rand.Intn(999999))
}

var (
	_ caddyhttp.MiddlewareHandler = (*RequestID)(nil)
)
