# KellyFramework

## TL;DR

kellyframework是一个golang的HTTP JSON API框架. 简而言之, 它可以很容易地把一个golang的函数或者对象方法转换成HTTP JSON API, 使用者只需要
考虑该函数或方法应当对应哪个URL路径, 而不用关心json编解码/url query string解析/url pattern解析等各种琐事.

## Example

它用起来就像这样:
```go
package main

import "os"
import "net/http"
import "github.com/abadcafe/kellyframework"

type userInfo struct {
    Age int
    Address string
}

type userName struct {
    Name string
}

type user struct {
    userName
    userInfo
}

func addUser(ctx *kellyframework.ServiceMethodContext, user *user) error {
    // do add user or return error
    return nil
}

func getUser(ctx *kellyframework.ServiceMethodContext, name *userName) interface{} {
    // return userInfo or return error
    return &userInfo{}
}

func main() {
    accessLogFile, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }

    routes := []*kellyframework.Route{
        {Method: "POST", Path: "/user/:Name", Function: addUser},
        {Method: "POST", Path: "/user/", Function: addUser},
        {Method: "GET", Path: "/user/:Name", Function: getUser, BypassRequestBody: true},
    }
    loggingServiceRouter, err := kellyframework.NewLoggingHTTPRouter(routes, nil, accessLogFile)
    if err != nil {
        panic(err)
    }

    http.Handle("/", loggingServiceRouter)
    http.ListenAndServe(":8080", nil)
}
```
就这么几行, 就得到了一个带access log功能的HTTP JSON API服务器.

可以用curl与之交互:
```shell
# 添加用户, name这个字段放在url pattern里, 优先级最高
curl -H "Content-Type: application/json" -d '{"Age": 18, "Address": "beijing"}' http://127.0.0.1:8080/user/test
# 或者把name字段直接放在json body里, 优先级次高
curl -H "Content-Type: application/json" -d '{"Name": "test", "Age": 18, "Address": "beijing"}' http://127.0.0.1:8080/user/
# 或者把name字段放在query string里, 优先级最低
curl -H "Content-Type: application/json" -d '{"Age": 18, "Address": "beijing"}' http://127.0.0.1:8080/user/?name=test

# 获取用户信息
curl -G http://127.0.0.1:8080/user/test
```
**如果函数返回error或者panic了, 该框架会自动封装成HTTP的500错误.**

## FAQ

### kellyframework对函数原型是否有要求?

是的, 只支持一种函数原型:

`func(*ServiceMethodContext, *struct) anything`

框架会把url pattern/json/query string解析成struct, 并把第一个返回值给json encode之后放在response body里输出.

如果返回的是error类型,则按以下格式输出:
```json
{
  "code": 500,
  "msg": "service method error", 
  "data": "returned error string"
}
```

### 我想返回自定义的错误码怎么办?

你可以在你的函数中, 返回一个`*kellyframework.FormattedResponse`结构体, 在其中你可以填写你想要的code, msg和data字段内容, 并且会把code字段赋值给http状态码返回.

### 我的函数并不关心输入的参数怎么办?

直接这样定义函数就好了:
```
func getSomething(_ *kellyframework.ServiceMethodContext, _ *struct{}) anything {
    return anything
}
```

### 看上去不错, 但那些struct的字段的有效性验证起来很麻烦.

kellyframework集成了[validator](https://godoc.org/gopkg.in/go-playground/validator.v9), 可以使用validator的struct tag语法为
struct添加各种约束.

例如:
```
type userInfo struct {
    Age int        `validate:"max=140,min=0"`
    Address string
}
```

一旦输入的字段不符合约束, http框架即会返回HTTP的400错误.

### `kellyframework.ServiceMethodContext`是干吗用的?

这个结构体包含以下字段:
```
type ServiceMethodContext struct {
	Context            context.Context // 该HTTP请求的context
	RemoteAddr         string // 发起请求的远端地址, 有可能是反向代理服务器的IP地址.
	RequestHeader      http.Header // 请求头, http.Header类型, 各字段都可以读.
	RequestBodyReader  io.ReadCloser // request body
	ResponseHeader     http.Header // 响应头, http.Header类型, 可以往里添各种http头的字段.
	ResponseBodyWriter io.Writer // response body
}
```
这些字段都可以随便使用.

但需要注意的是, 如果你想自己处理request body, 那么你就应当把`Route.BypassRequestBody`设为true, 这样框架就会忽略request body, 留给你自
己的函数来解析.

而如果你想自己返回response body, 那么你就应当把`Route.BypassResponseBody`设为true, 这样框架就不会尝试按json格式去encode你返回的结构体.
注意: 如果你返回的response body是`kellyframework.FormattedResponse`类型或者`error`类型, 框架仍然会按json格式去encode你返回的结构体.

### access log是否支持自动切分?

日志切分这个事情并不应当由HTTP JSON API框架来完成, 你可以用[autosplitfile](https://github.com/abadcafe/autosplitfile)来替代普通的
os.File传进kellyframework.NewLoggingHTTPRouter()方法里, 这样你的access log就是自动切分的了.
