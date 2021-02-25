package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Println("started on port(s): 8080 (http) with context path ''")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Health struct {
	UP bool `json:"up"`
}

func handle(resp http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("request url = %v", req.RequestURI)
	fmt.Println(s)
	h := Health{UP: true}
	b, _ := json.Marshal(h)
	fmt.Printf("proxy: ---> %v", req.Header.Get("proxy"))
	resp.Header().Set("hello world", "demo")
	_, _ = resp.Write(b)
}
