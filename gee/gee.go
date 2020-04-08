//

package gee

import (
	"fmt"
	"net/http"
)

//HandlerFunc 定义handerfunc
type HandlerFunc func(http.ResponseWriter, *http.Request)

//Engine 实现ServeHTTP接口
type Engine struct {
	router map[string]HandlerFunc
}

//New 实例化一个Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

//addRoute 添加路由
func (engine *Engine) addRoute(method string, patten string, handler HandlerFunc) {
	key := method + "-" + patten
	engine.router[key] = handler
}

//GET方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//ServeHTTP 实现
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handle, ok := engine.router[key]; ok {
		handle(w, req)
	} else {
		fmt.Fprintf(w, "404 not found %s\n", req.URL)
	}
}

//Run 运行
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
