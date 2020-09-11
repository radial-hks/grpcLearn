package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpcLearn/server/helper"
	"grpcLearn/server/services"
	"log"
	"net/http"
)

func main(){
	// 使用Gateway启动HTTP Server
	gwmux := runtime.NewServeMux()
	//
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}

	err := services.RegisterProductServiceHandlerFromEndpoint(context.Background(),
		gwmux, "localhost:8081", opt)
	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}

	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(),
		gwmux, "localhost:8081", opt)

	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()

}
