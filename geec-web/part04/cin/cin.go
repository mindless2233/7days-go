package cin

import (
	"net/http"
)

// HandlerFunc 只要是入参符合这个形式的统统可以叫做 HandlerFunc 型的函数
type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//key := method + "-" + pattern
	//e.router[key] = handler
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

// 结构体Engine实现了ServeHTTP方法。
// 第一个参数是 ResponseWriter，利用 ResponseWriter 可以构造针对该请求的响应。
// 第二个参数是 Request，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//组装 Engine
	//key := req.Method + "-" + req.URL.Path
	//if handler, ok := e.router[key]; ok {
	//	handler(w, req)
	//} else {
	//	fmt.Fprintf(w, "404 Not Found:%s", req.URL)
	//}

	c := newContext(w, req)
	e.router.handle(c)
}
