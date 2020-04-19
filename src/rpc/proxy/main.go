package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)
// 实现反向代理
func proxyFunc(w http.ResponseWriter, r *http.Request) {
	u,_ := url.Parse("http://127.0.0.1:4040")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w,r)
}

func main() {

	http.HandleFunc("/", proxyFunc)
	err := http.ListenAndServe(":5050",nil)
	if err != nil {
		fmt.Println("isten 5050")
	}
}
