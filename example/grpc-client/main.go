package main

import (
	"encoding/json"
	"flag"
	"github.com/ClareChu/proxy-istio-demo/example/grpc-client/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

var addr = flag.String("addr", "localhost", "The address of the server to connect to")
var port = flag.String("port", "7575", "The port to connect to")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(
		net.JoinHostPort(*addr, *port),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" Serving web on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", RegistryHandle(conn)))
}

func RegistryHandle(conn *grpc.ClientConn) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", &demoHandler{
		conn: conn,
	})
	return mux
}

type demoHandler struct {
	conn *grpc.ClientConn
}

func (d *demoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")
	log.Printf("get message data--> %v", message)
	b, err := json.Marshal(pkg.Get(d.conn, message))
	if err != nil {
		log.Printf("%+v", err)
	}
	w.Write(b)
}
