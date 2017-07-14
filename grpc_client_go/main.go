package main

import (
	"flag"
	pb "github.com/fukajun/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
)

var (
	addrFlag = flag.String("addr", "0.0.0.0:5000", "server address host:post")
)

func main() {
	conn, err := grpc.Dial(*addrFlag, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	//接続は最後に必ず閉じる
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	stream, err := c.Watch(context.Background(), &pb.Channel{Name: "stream fukajun"})
	if err != nil {
		log.Printf("greeting")
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("fatal")
		}
		log.Printf("%#v", feature.Name)
	}

	//サーバーに対してリクエストを送信する
	//for value := 0; ; value++ {
	//resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "fukajun" + fmt.Sprintf("%d", value)})
	//if err != nil {
	//log.Fatalf("RPC error: %v", err)
	//}
	//log.Printf("Greeting: %s %d", resp.Message, value)
	//}
}
