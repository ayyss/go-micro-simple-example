package web_pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// InitMiddleware 拦截请求将rpc服务绑定到上下文中
func InitMiddleware(microServices []interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["userRpcService"] = microServices[0]
		ctx.Keys["productRpcService"] = microServices[1]
		ctx.Next()
	}
}

// ErrorMiddleware 错误统一处理
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(404, gin.H{"code": 404, "msg": fmt.Sprintf("%s", r)})
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
