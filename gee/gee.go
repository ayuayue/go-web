//

package gee

import (
	"net/http"
)

//HandlerFunc 定义handerfunc
type HandlerFunc func(*Context)

//Engine 实现ServeHTTP接口
type Engine struct {
	router *router
}

//New 实例化一个Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

//addRoute 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}
//GET 方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//POST 方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//ServeHTTP 实现
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

//Run 运行
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
