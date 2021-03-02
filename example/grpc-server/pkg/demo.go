package pkg

import (
	"context"
	"github.com/ClareChu/proxy-istio-demo/example/proto"
	"github.com/uber/jaeger-client-go/crossdock/log"
	"google.golang.org/grpc"
)

func NewDemo(server *grpc.Server) {
	proto.RegisterDemoInterfaceServer(server, &Demo{})
}

type Demo struct {
}

func (d *Demo) Get(ctx context.Context, req *proto.DemoRequest) (*proto.DemoResponse, error) {
	log.Printf("request message ---> %v", req.Message)
	resp := &proto.DemoResponse{
		Code:    200,
		Message: "success",
		Data:    req.Message,
	}
	return resp, nil
}
