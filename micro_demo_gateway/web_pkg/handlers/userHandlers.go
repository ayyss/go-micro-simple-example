package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo_gateway/services"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func UserLogin(ctx *gin.Context) {
	//将请求参数与rpc请求绑定
	var rpcReq services.UserRequest
	PanicIfError(ctx.Bind(&rpcReq))
	//从上下文中获取rpc服务
	userMicroService := ctx.Keys["userRpcService"].(services.UserService)
	//调用rpc服务
	rpcResp, err := userMicroService.UserLogin(context.Background(), &rpcReq)
	PanicIfError(err)
	ctx.JSON(200, gin.H{"data": rpcResp})
}

func UserRegister(ctx *gin.Context) {
	var rpcReq services.UserRequest
	PanicIfError(ctx.Bind(&rpcReq))
	userMicroService := ctx.Keys["userRpcService"].(services.UserService)
	rpcResp, err := userMicroService.UserRegister(context.Background(), &rpcReq)
	PanicIfError(err)
	ctx.JSON(200, gin.H{"data": rpcResp})
}
