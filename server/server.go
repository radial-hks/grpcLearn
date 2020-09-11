package main

import (
	"google.golang.org/grpc"
	"grpcLearn/server/helper"
	"grpcLearn/server/services"
	"log"
	"net"
)

func main() {
	// 加载证书文件
	//creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	//if err != nil {
	//	log.Fatalf("加载服务端证书和Key失败, err: %v\n", err)
	//}

	//cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	//if err != nil {
	//	log.Fatalf("加载服务端证书失败, err: %v\n", err)
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
	//	ClientAuth:   tls.RequireAndVerifyClientCert,
	//	ClientCAs:    certPool,
	//})

	//rpcServer := grpc.NewServer(grpc.Creds(creds))

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredentials()))
	// 新建grpc服务
	//rpcServer := grpc.NewServer()
	// 注册服务
	services.RegisterProductServiceServer(rpcServer, new(services.ProdService))
	services.RegisterOrderServiceServer(rpcServer,new(services.OrderService))
	services.RegisterUserServiceServer(rpcServer,new(services.UserService))
	// 启动监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
	//mux:= http.NewServeMux()
	//mux.HandleFunc("/",func(writer http.ResponseWriter,request *http.Request){
	//	fmt.Println(request.Header)
	//	fmt.Println(request.Proto)
	//	rpcServer.ServeHTTP(writer,request)
	//})
	//httpServer:=&http.Server{
	//	Addr: ":8081",
	//	Handler: mux,
	//}
	//httpServer.ListenAndServe()
	// http://localhost:8081/
	//httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_passwd.key")
}
