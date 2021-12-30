package cin

import (
	"log"
	"net/http"
)

//形参由(http.ResponseWriter, *http.Request)改为(*Context)
//type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router.addRouter(method, pattern, handlerFunc)
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	log.Printf("GET-%s", pattern)
	engine.addRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	log.Printf("POST-%s", pattern)
	engine.addRoute("POST", pattern, handlerFunc)
}
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//key := req.Method + "-" + req.URL.Path
	//if handler, ok := engine.router[key]; ok {
	//	handler(w, req)
	//} else {
	//	fmt.Fprint(w, "404 NOT FOUND: %s\n", req.URL)
	//}
	c := newContext(w, req)
	engine.router.handle(c)
}
