module github.com/ClareChu/proxy-istio-demo

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/MicahParks/keyfunc v0.3.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/segmentio/kafka-go v0.4.28
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.6.1
	github.com/tetratelabs/proxy-wasm-go-sdk v0.1.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	google.golang.org/grpc v1.16.0
	k8s.io/klog/v2 v2.60.1
)

replace github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible => github.com/clarechu/email v4.0.1-0.20220128100025-d1a37658465a+incompatible
