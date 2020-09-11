package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpcLearn/client/helper"
	"grpcLearn/client/services"
	"io"
	"log"
	"time"
)

func main() {
	//creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "sensetime.com")
	//if err != nil {
	//	log.Fatalf("加载客户端证书失败, err: %v\n", err)
	//}


	//cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	//if err != nil {
	//	log.Fatalf("加载客户端证书失败, err: %v\n", err)
	//}
	//
	//certPool := x509.NewCertPool()
	//ca, err := ioutil.ReadFile("cert/ca.pem")
	//if err != nil {
	//	log.Fatalf("读取公钥文件失败: %v\n", err)
	//}
	//
	//certPool.AppendCertsFromPEM(ca)
	//
	//creds := credentials.NewTLS(&tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	ServerName:   "localhost",
	//	RootCAs:      certPool,
	//})


	// 连接 8081 (不使用任何证书)
	//conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()
	// 返回定义的客户端
	prodClient := services.NewProductServiceClient(conn)
	// 客户端发送请求
	//prodRes, err := prodClient.GetProductStock(context.Background(),
	//	&services.ProdRequest{ProdId: 12,ProdArea: services.ProdAreas_C}, )
	prodRes, err := prodClient.GetProductInfo(context.Background(),
		&services.ProdRequest{ProdId: 12}, )
	// 打印返回的结果
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(prodRes)


	//ctx := context.Background()
	//prodRes, err := prodClient.GetProdStocks(ctx,&services.QuerySize{Size: 12},)
	//if err != nil {
	//	log.Fatalf("请求GRPC服务端失败 %v\n", err)
	//}
	//fmt.Println(prodRes.Prodres[2].ProdStock)



	t := timestamp.Timestamp{
		Seconds: time.Now().Unix(),
	}
	orderClient := services.NewOrderServiceClient(conn)
	orderRes, err := orderClient.NewOrder(context.Background(),
		&services.OrderRequest{OrderMain: &services.OrderMain{
			Order: 32,
			OrderNo: "1111",
			OrderMoney: -2,
			OrderTime: &t,
		}},
		)
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(orderRes)


	//
	userClient := services.NewUserServiceClient(conn)
	users := make([]*services.UserInfo, 0)
	var i int32 = 0
	for i = 0; i < 18; i++ {
		user := &services.UserInfo{UserId: i + 1}
		users = append(users, user)
	}

	stream, err := userClient.GetUserScoreByServerStream(context.Background(),
		&services.UserScoreRequest{Users: users},
	)

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	for {
		userRes, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("读取服务端流失败 err: %v\n", err.Error())
		}
		fmt.Println(userRes.Users)
	}

}

