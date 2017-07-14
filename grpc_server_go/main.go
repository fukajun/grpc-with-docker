package main

import (
	"flag"
	pb "github.com/fukajun/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	addrFlag = flag.String("addr", ":5000", "Address host:post")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//log.Printf("New Reqest: %v", in.String())
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func (s *server) Watch(channel *pb.Channel, stream pb.Greeter_WatchServer) error {
	log.Printf("Request receive channel" + channel.Name)
	log.Printf("%#v", stream)
	nums := []int{1, 2, 3, 4, 5}
	for _, v := range nums {
		stream.Send(&pb.Notification{Name: strconv.Itoa(v)})
		time.Sleep(3 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", *addrFlag)

	if err != nil {
		log.Fatalf("boo")
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
