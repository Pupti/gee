package main

// curl http://localhost:9999/
// URL.Path = "/"
//  curl http://localhost:9999/hello
// Head["User-Agent"] = ["curl/7.63.0"]
// Head["Accept"] = ["*/*"]
//  curl http://localhost:9999/world
// 404 NOT FOUND%!(EXTRA *url.URL=/world)

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.path = %q", req.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q", k, v)
		}
	})

	// http.listenAndServe 启动服务器
	// 该函数的第二个参数是一个接口，如果这个接口实现serveHTTP函数，
	// 只要传入实现serveHTTP接口的实例，所有的http请求就由该实例处理
	r.Run(":9999")
}
