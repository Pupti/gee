## Day1 web框架初步搭建

从直接使用net/http官方标准库，到实现engine实例，在engine实例中实现ServeHTTP方法，这样就可以拦截所有的HTTP请求，有了统一入口，这样可以添加一些处理逻辑。

### 需要注意的

1. 在http.ListenAndServe(":9999", engine)中，并没有直接调用engine的ServeHTTP方法，但是会执行，这里是因为http.ListenAndServe的第二个参数是接口类型http.Handler，在Go语言中，实现了接口方法的struct都可以强制转换为接口类型，可以写为
```golang
    handle := (http.Handle)(engine)
    http.ListenAndServe(":9999", handler)
```
然后ListenAndServe会调用handle.ServeHTTP方法。但是这么写是多余的，传参时会自动进行参数转换。

最终的版本是按照gin框架的雏形设定的，实例化engine变量，统一调度服务器资源，从路由GET、POST添加，到ServeHTTP的统一调度以及服务器的启动，都有engine完成，调用者只需要实例化engine，再调用对应的方法请求，同时定义对应的处理方法。

2. http.ServeHTTP的参数有两个，一个是应答包（接口类型），一个是请求包，其中后者为指针，为了节省空间，前者因为是接口类型，不能通过指针调用，容易引起bug。
3. 子模块的gomod，创建一个子模块后，需要自行`go mod init mokuai-name`。 以及其它各种使用

