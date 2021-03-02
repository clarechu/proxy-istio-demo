package pkg

import (
	"github.com/ClareChu/proxy-istio-demo/example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func Get(conn *grpc.ClientConn, message string) *proto.DemoResponse {
	client := proto.NewDemoInterfaceClient(conn)
	req := &proto.DemoRequest{
		Message: message,
	}
	resp, err := client.Get(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response --> %+v", resp)
	return resp
}
