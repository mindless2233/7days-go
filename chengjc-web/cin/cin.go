package cin

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

//形参由(http.ResponseWriter, *http.Request)改为(*Context)
//type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)
type Engine struct {
	router *router
	*RouterGroup
	groups        []*RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

//新增分则结构体
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func New() *Engine {

	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
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

	//2022/1/6注释
	//c := newContext(w, req)
	//engine.router.handle(c)

	//2022/1/6新增
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}
func (group *RouterGroup) Group(prefix string) *RouterGroup {

	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
func (group *RouterGroup) addRoute(method string, comp string, handlerFunc HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRouter(method, pattern, handlerFunc)
}
func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("GET", pattern, handlerFunc)
}
func (group *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	group.addRoute("POST", pattern, handlerFunc)
}
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServe := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(context *Context) {
		file := context.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			context.Status(http.StatusNotFound)
			return
		}
		fileServe.ServeHTTP(context.Writer, context.Req)
	}
}
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))

}
