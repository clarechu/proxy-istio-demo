package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handle).Methods(http.MethodGet)
	srv := &http.Server{
		// new router
		Handler: r,
		Addr:    fmt.Sprintf(":%d", 8080),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("started on port(s): 8080 (http) with context path ''")
	log.Fatal(srv.ListenAndServe())
}

type Health struct {
	UP bool `json:"up"`
}

func handle(resp http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("request url = %v", req.RequestURI)
	fmt.Println(s)
	h := &Health{UP: true}
	b, _ := json.Marshal(h)
	fmt.Printf("proxy: ---> %v", req.Header.Get("proxy"))
	//resp.Header().Set("hello world", "demo")
	_, _ = resp.Write(b)
}
