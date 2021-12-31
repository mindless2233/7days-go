package cin

import (
	"net/http"
	"strings"
)

//主要是对路由的相关操作
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
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

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	//log.Printf("Route %4s --- %s", method, pattern)
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

func (r *router) getRoute(method string, pattern string) (*node, map[string]string) {
	serachParts := parsePattern(pattern)
	parms := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.seach(serachParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				parms[part[1:]] = serachParts[index]
			}
			if part[0] == '*' || len(part) > 1 {
				parms[part[1:]] = strings.Join(serachParts[index:], "/")
				break
			}
		}
		return n, parms
	}
	return nil, nil

}
func (r *router) handle(c *Context) {
	//key := c.Method + "-" + c.Path
	//if handler, ok := r.handlers[key]; ok {
	//	handler(c)
	//} else {
	//	c.String(http.StatusNotFound, "NOT FOUND:&s\n", c.Path)
	//}
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Parms = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
