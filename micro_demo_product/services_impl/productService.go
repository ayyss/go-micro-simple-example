package services_impl

import (
	"context"
	"micro_demo_product/services"
)

var productMap = make(map[string]uint32)

type ProductService struct {
}

//func (*ProductService) ProductQuery(ctx context.Context, req *services.ProductRequest, resp *services.ProductResponse) error {
//	if _, ok := productMap[req.ProductName]; !ok {
//		resp.Code = 1
//		resp.Msg = "产品不存在"
//	} else {
//		resp.Code = 0
//		resp.Data = &services.ProductModel{ProductName: req.ProductName, Number: productMap[req.ProductName]}
//	}
//	return nil
//}

func (*ProductService) ProductPush(ctx context.Context, req *services.ProductRequest, resp *services.ProductResponse) error {
	if _, ok := productMap[req.ProductName]; !ok {
		productMap[req.ProductName] = req.Number
	} else {
		productMap[req.ProductName] += req.Number
	}
	resp.Code = 0
	resp.Data = &services.ProductModel{ProductName: req.ProductName, Number: productMap[req.ProductName]}
	return nil
}

//func (*ProductService) ProductPop(ctx context.Context, req *services.ProductRequest, resp *services.ProductResponse) error {
//	if _, ok := productMap[req.ProductName]; !ok {
//		resp.Code = 1
//		resp.Msg = "产品不存在"
//	} else if productMap[req.ProductName] < req.Number {
//		resp.Code = 1
//		resp.Msg = "产品库存不足"
//	} else {
//		productMap[req.ProductName] -= req.Number
//		resp.Code = 0
//		resp.Data = &services.ProductModel{ProductName: req.ProductName, Number: productMap[req.ProductName]}
//	}
//	return nil
//}
