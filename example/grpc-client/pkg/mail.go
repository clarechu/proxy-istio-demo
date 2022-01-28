package pkg

import (
	"context"
	"github.com/ClareChu/proxy-istio-demo/example/proto"
	"google.golang.org/grpc"
	"log"
)

func Mail(conn *grpc.ClientConn, message string) error {
	client := proto.NewMailInterfaceClient(conn)
	req := &proto.MailRequest{
		Message: message,
	}
	resp, err := client.Get(context.TODO(), req)
	if err != nil {
		return err
	}
	log.Printf("response --> %+v", resp)
	return err
}
