package router

import (
	"fmt"
	"io"
	log "k8s.io/klog/v2"
	"net"
	"net/http"
	"os"
	"strings"
)

//ComponentHandler
// @ID ComponentHandler
// @Summary 反向代理所有第三方路由
// @Description 反向代理所有第三方路由
// @Accept  json
// @Tags Component
// @Success 200 {object} object  "success"
// @Router / [get]
func ComponentHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Received request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)
	transport := http.DefaultTransport
	outReq := new(http.Request)
	*outReq = *r // this only does shallow copies of maps
	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}
	outReq.Host = fmt.Sprintf("%s:%s", os.Getenv("PROXY_IP"), r.URL.Port())
	log.Infof("proxy url --> %s", outReq.Host)
	outReq.URL.Host = outReq.Host
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer res.Body.Close()
	for key, value := range res.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}
