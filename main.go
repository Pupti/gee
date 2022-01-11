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
	"log"
	"net/http"
)

type Engine struct{}

// http.ResponseWriter 是一个接口， 因为接口有其本身的类型，在使用中会带来陷阱
// http.Request 是指针类型，这样可以节省内存
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Head[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND\n", req.URL)
	}
}

func main() {
	engine := new(Engine)

	// http.listenAndServe 启动服务器
	// 该函数的第二个参数是一个接口，如果这个接口实现serveHTTP函数，
	// 只要传入实现serveHTTP接口的实例，所有的http请求就由该实例处理
	log.Fatal(http.ListenAndServe(":9999", engine))
}
