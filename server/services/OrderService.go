package services

import (
	"fmt"
	"golang.org/x/net/context"
)

type OrderService struct {

}

func (o *OrderService)NewOrder(ctx context.Context, in *OrderRequest) (*OrderResponse, error){
	fmt.Println(in)
	if err:=in.OrderMain.Validate();err!=nil{
		return &OrderResponse{
			Message: "failed",
			Status: err.Error(),
		},nil
	}
	res := OrderResponse{
		Message: "worked",
		Status: "success",
	}
	return &res,nil
}