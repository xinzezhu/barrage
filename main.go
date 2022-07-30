package main

import (
	"context"
	"fmt"
	"net"
	"time"

	read_proc "github.com/xinzezhu/barrage_protocol/read_proc"   // 引入编译生成的包
	write_proc "github.com/xinzezhu/barrage_protocol/write_proc" // 引入编译生成的包
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	ServerAddress   = "127.0.0.1:20000"
	ReadProcAddress = "127.0.0.1:30000"
)

func main() {
	// 连接
	conn, err := grpc.Dial(ReadProcAddress, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := read_proc.NewReadProcClient(conn)

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				body := read_proc.Body{}
				last_sec_time := time.Now().Unix() - 1
				if v, ok := BarrageSecMap.Load(last_sec_time); ok {
					roomIdMap := v.(map[uint32][]string)
					var barrge read_proc.BarragePerSecond
					for k, v := range roomIdMap {
						barrge.RoomId = k
						for i := 0; i < len(v); i++ {
							barrge.Content = append(barrge.Content, v[i])
						}
						body.Barrage = append(body.Barrage, &barrge)
					}

					// 同步弹幕
					body.TimeStamp = uint64(last_sec_time)
					req := &read_proc.Request{Body: &body}
					c.SyncBarrage(context.Background(), req)
					fmt.Println(req)
				} else {
					fmt.Printf("time:%d,这一秒没有任何弹幕\n", last_sec_time)
				}

			}
		}
	}()

	listen, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	write_proc.RegisterWirteProcServer(s, HandleWriteProc)

	fmt.Println("Listen on " + ServerAddress)
	s.Serve(listen)
}
