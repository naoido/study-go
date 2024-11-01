package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	hellopb "naoido.com/study-go/pkg/grpc"
	"net"
	"os"
	"os/signal"
)

const (
	PORT int = 8080
)

var (
	users = map[int]string{
		1: "naoido",
		2: "なおいど",
		3: "一般男性",
	}
)

func NewMyServer() *MyServer {
	return &MyServer{}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	hellopb.RegisterGrpcServiceServer(s, NewMyServer())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", PORT)
		err := s.Serve(listener)
		if err != nil {
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

type MyServer struct {
	hellopb.UnimplementedGrpcServiceServer
}

func (s *MyServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := ""
	switch user := req.GetUser().(type) {
	case *hellopb.HelloRequest_Id:
		name = users[int(user.Id)]
	case *hellopb.HelloRequest_Name:
		name = user.Name
	}
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}, nil
}
