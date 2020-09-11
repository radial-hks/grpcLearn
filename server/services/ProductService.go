package services

import (
	"fmt"
	"golang.org/x/net/context"
)

type ProdService struct{}

func (*ProdService) GetProductStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0;
	if req.ProdArea == ProdAreas_A{
		stock = 40;
	} else if req.ProdArea == ProdAreas_B {
		stock = 30;
	} else {
		stock = 100;
	}
	return &ProdResponse{
		ProdStock: stock,
	}, nil
}

func (p *ProdService)GetProdStocks(ctx context.Context, req *QuerySize) (*ProdResponseList, error){
	fmt.Println(req)
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 100,},
		&ProdResponse{ProdStock: 101,},
		&ProdResponse{ProdStock: 102,},
		&ProdResponse{ProdStock: 103,},
	}
	return &ProdResponseList{Prodres: Prodres},nil
}

func (p *ProdService)GetProductInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error){
	fmt.Println(in)
	ret:= ProdModel{
		ProdName: "test product",
		ProdPrice: 14,
		Id: 1,
	}
	return &ret,nil
}


