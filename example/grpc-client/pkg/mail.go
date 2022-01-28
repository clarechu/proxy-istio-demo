package pkg

import (
	"context"
	"github.com/ClareChu/proxy-istio-demo/example/proto"
	"google.golang.org/grpc"
	"log"
)

func Mail(conn *grpc.ClientConn, message string) *proto.MailResponse {
	client := proto.NewMailInterfaceClient(conn)
	req := &proto.MailRequest{
		Message: message,
	}
	resp, err := client.Get(context.TODO(), req)
	if err != nil {
		log.Println(err)
	}
	log.Printf("response --> %+v", resp)
	return resp
}
