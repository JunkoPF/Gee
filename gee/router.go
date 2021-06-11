package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	trees    map[string]*tree
	handlers map[string]HandlerFunc
}

type params map[string]string

func newRouter() *router {
	return &router{
		trees:    make(map[string]*tree),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	strs := strings.Split(pattern, "/")
	parts := strs[1:]
	for ind, part := range parts {
		if part == "" {
			return parts[:ind]
		} else if part == "*" {
			return parts[:ind+1]
		}
	}
	return parts
}

func unparsePattern(values []string) string {
	return "/" + strings.Join(values, "/")
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	values := parsePattern(pattern)
	if _, ok := r.trees[method]; !ok {
		r.trees[method] = makeTree()
	}
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	r.trees[method].insert(values)
	r.handlers[key] = handler

}

func (r *router) getRoute(method string, path string) (*node, params) {
	values := parsePattern(path)
	tr, ok := r.trees[method]
	if !ok {
		return nil, nil
	}

	if nd := tr.search(values); nd != nil {
		if nd.pattern != nil {
			ps := make(params)
			for i := 0; i < len(nd.pattern); i++ {
				switch nd.pattern[i][0] {
				case ':':
					{
						ps[nd.pattern[i][1:]] = values[i]
					}
				case '*':
					{
						ps[nd.pattern[i][1:]] = unparsePattern(values[i:])
					}
				}
			}
			return nd, ps
		}
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, ps := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + unparsePattern(n.pattern)
		c.Params = ps
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		return
	}
	c.Next()
}
