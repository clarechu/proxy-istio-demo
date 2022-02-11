package main

import (
	"flag"
	"github.com/ClareChu/proxy-istio-demo/example/grpc-server/pkg"
	"github.com/ClareChu/proxy-istio-demo/pkg/tracing"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

var kafkaAddress = flag.String("kafka", "10.10.13.110:9092", "The kafka address of the server to connect to")

const (
	SamplingServerURL  = ""
	LocalAgentHostPort = ""
	ServerAddr         = ":7575"
)

func main() {
	flag.Parse()
	tracer, closer := tracing.NewTracing(SamplingServerURL, LocalAgentHostPort)
	defer closer.Close()
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
		}),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)),
	)
	lis, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	go func() {
		pkg.Consume(*kafkaAddress)
	}()
	RegistryGrpc(server)
	log.Printf("Serving gRPC on grpc://localhost%v", ServerAddr)
	server.Serve(lis)
}

func RegistryGrpc(server *grpc.Server) {
	pkg.NewDemo(server)
	pkg.NewMail(server)
}
