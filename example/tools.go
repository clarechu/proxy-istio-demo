package example

//go:generate protoc -I=. --proto_path=./proto  --go_out=plugins=grpc,paths=source_relative:./proto demo.proto
//go:generate protoc -I=. --proto_path=./proto  --go_out=plugins=grpc,paths=source_relative:./proto health.proto
