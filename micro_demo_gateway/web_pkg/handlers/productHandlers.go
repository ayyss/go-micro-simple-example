package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_demo_gateway/services"
)

//func ProductQuery(ctx *gin.Context) {
//	var rpcReq services.ProductRequest
//	PanicIfError(ctx.BindQuery(&rpcReq))
//	productRpcService := ctx.Keys["productRpcService"].(services.ProductService)
//	rpcResp, err := productRpcService.ProductQuery(context.Background(), &rpcReq)
//	PanicIfError(err)
//	ctx.JSON(200, gin.H{"data": rpcResp})
//}

func ProductPush(ctx *gin.Context) {
	var rpcReq services.ProductRequest
	PanicIfError(ctx.Bind(&rpcReq))
	productRpcService := ctx.Keys["productRpcService"].(services.ProductService)
	rpcResp, err := productRpcService.ProductPush(context.Background(), &rpcReq)
	PanicIfError(err)
	ctx.JSON(200, gin.H{"data": rpcResp})
}

//func ProductPop(ctx *gin.Context) {
//	var rpcReq services.ProductRequest
//	PanicIfError(ctx.Bind(&rpcReq))
//	productRpcService := ctx.Keys["productRpcService"].(services.ProductService)
//	rpcResp, err := productRpcService.ProductPop(context.Background(), &rpcReq)
//	PanicIfError(err)
//	ctx.JSON(200, gin.H{"data": rpcResp})
//}
