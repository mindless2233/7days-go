package cin

import (
	"net/http"
	"strings"
)

//将和路由相关的方法和结构提取了出来，放到了一个新的文件中router.go，
//方便我们下一次对 router 的功能进行增强，例如提供动态路由的支持。
//router 的 handle 方法作了一个细微的调整，即 handler 的参数，变成了 Context。

// Trie 树的插入与查找都成功实现了，接下来我们将 Trie 树应用到路由中去吧 。
// 我们使用 roots 来存储每种请求方式的Trie 树根节点。
// 使用 handlers 存储每种请求方式的 HandlerFunc 。
//  getRoute 函数中，还解析了:和*两种匹配符的参数，返回一个 map 。
// 例如/p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}，
// /static/css/geektutu.css匹配到/static/*filepath，解析结果为{filepath: "css/geektutu.css"}。

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}

	return parts
}
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	//log.Printf("Route %4s - %s", method, pattern)
	//key := method + "-" + pattern
	//r.handlers[key] = handler

	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}
func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)

		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}

		}
		return n, params
	}
	return nil, nil
}
func (r *router) handle(c *Context) {
	//key := c.Method + "-" + c.Path
	//if handle, ok := r.handlers[key]; ok {
	//	handle(c)
	//} else {
	//	c.String(http.StatusNotFound, "404 NOT FOUND:%S\n", c.Path)
	//}

	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%S\n", c.Path)
	}

}
