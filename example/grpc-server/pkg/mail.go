package pkg

import (
	"context"
	"github.com/ClareChu/proxy-istio-demo/example/grpc-server/service"
	"github.com/ClareChu/proxy-istio-demo/example/proto"
	"github.com/uber/jaeger-client-go/crossdock/log"
	"google.golang.org/grpc"
)

func NewMail(server *grpc.Server) {
	proto.RegisterMailInterfaceServer(server, &Mail{})
}

type Mail struct {
}

func (d *Mail) Get(ctx context.Context, req *proto.MailRequest) (*proto.MailResponse, error) {
	log.Printf("request message ---> %v", req.Message)
	err := service.SendMailByServer(req.Message)
	return &proto.MailResponse{}, err
}
