package server

import (
	"context"
	v1 "grpc_demo/api/proto/v1"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)
	c := make(chan os.Signal, 1)
	// goroutine捕获系统中断信号
	go func() {
		for range c {
			log.Println("shutting down grpc server")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()
	log.Println("starting grpc server")
	return server.Serve(listener)

}
