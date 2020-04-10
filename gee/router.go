package gee

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

//router roots => key,eg: roots['GET'],roots['POST']
//handlers key eg: handlers['GET-/p/:lang/doc'],handlers['POST-/p/book']
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

//newRouter 实例化
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//Only one * is allowed
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

//addRoute 注册路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

//getRoute 获取路由
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
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

//handle 路由处理
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found :%s\n", c.Path)
	}

}

func newTestRoute() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

//TestParsePattern 测试解析pattern
func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test paresPattern failed")
	}
}

//TestGetRoute 测试获取路由
func TestGetRoute(t *testing.T) {
	r := newTestRoute()
	n, ps := r.getRoute("GET", "/hello/caoayu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if ps["name"] != "caoayu" {
		t.Fatal("name should be equal to 'caoayu'")
	}
	fmt.Printf("matched path: %s ,params['name']:%s\n", n.pattern, ps["name"])
}
