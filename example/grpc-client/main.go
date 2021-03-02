package main

import (
	"fmt"
	"github.com/ClareChu/proxy-istio-demo/example/grpc-client/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("aa")
	addr := "localhost"
	port := "7575"
	conn, err := grpc.Dial(
		net.JoinHostPort(addr, port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	pkg.Get(conn)
}
