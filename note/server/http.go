package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	// 新建响应头
	header := res.Header()
	header["test"] = []string{"test"}
	header["Content-Type"] = []string{"application/json"}
	// 写入请求头
	res.WriteHeader(http.StatusBadRequest)
	switch req.Method {
	case "GET":
		res.Write([]byte("get请求"))
	case "POST":
		a, _ := io.ReadAll(req.Body)
		res.Write(a)
	}
	// 打印请求头
	fmt.Println(req.Header)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", mux)
}
