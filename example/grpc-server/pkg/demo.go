package pkg

import (
	"context"
	"github.com/ClareChu/proxy-istio-demo/example/proto"
	"google.golang.org/grpc"
)

func NewDemo(server *grpc.Server) {
	proto.RegisterDemoInterfaceServer(server, &Demo{})
}

type Demo struct {
}

func (d *Demo) Get(ctx context.Context, req *proto.DemoRequest) (*proto.DemoResponse, error) {
	resp := &proto.DemoResponse{
		Code:    200,
		Message: "success",
		Data:    req.Message,
	}
	return resp, nil
}
