package web_pkg

import (
	"github.com/gin-gonic/gin"
	"micro_demo_gateway/web_pkg/handlers"
)

func InitRouter(microServices ...interface{}) *gin.Engine {
	engine := gin.Default()
	//绑定中间件
	engine.Use(InitMiddleware(microServices), ErrorMiddleware())
	//api接口分组
	userGroup := engine.Group("/api/user")
	{
		userGroup.Handle("POST", "login", handlers.UserLogin)
		userGroup.Handle("POST", "register", handlers.UserRegister)
	}
	productGroup := engine.Group("/api/product")
	{
		//productGroup.Handle("GET", "query", handlers.ProductQuery)
		//productGroup.Handle("GET", "pop", handlers.ProductPop)
		productGroup.Handle("POST", "push", handlers.ProductPush)
	}

	return engine
}
