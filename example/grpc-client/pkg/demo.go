package pkg

import (
	"github.com/ClareChu/proxy-istio-demo/example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func Get(conn *grpc.ClientConn) {
	client := proto.NewDemoInterfaceClient(conn)
	req := &proto.DemoRequest{
		Message: "hello world",
	}
	resp, err := client.Get(context.TODO(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", resp)
}
