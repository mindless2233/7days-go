package cin

import (
	"log"
	"net/http"
)

//将和路由相关的方法和结构提取了出来，放到了一个新的文件中router.go，
//方便我们下一次对 router 的功能进行增强，例如提供动态路由的支持。
//router 的 handle 方法作了一个细微的调整，即 handler 的参数，变成了 Context。

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handle, ok := r.handlers[key]; ok {
		handle(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%S\n", c.Path)
	}

}
