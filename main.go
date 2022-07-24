package main

import (
	"fmt"

	pb "github.com/xinzezhu/protocol/barrage/read_proc" // 引入编译生成的包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewReadProcClient(conn)

	// 调用方法
	header := pb.Header{}
	header.RoomId = 1
	req := &pb.Request{Header: &header}
	res, err := c.SyncBarrage(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(res.Header.RoomId)
}
