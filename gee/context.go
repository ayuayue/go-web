//

package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//H 返回Data数据结构
type H map[string]interface{}

//Context 上下文,包含http.responseWriter,http.request,path,method,code
type Context struct {
	// 原始对象
	Writer http.ResponseWriter
	Req    *http.Request
	//请求信息
	Path   string
	Method string
	//相应信息
	StatusCode int
	//请求参数
	Params   map[string]string
	handlers []HandlerFunc
	index    int
	engine   *Engine
}

//newContext 实例化返回一个Context实例
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

//Next 处理中间件逻辑
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

//Param 参数处理
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

//PostForm 处理Post方式传name的值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//Query 处理get方式url参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//Status  返回response status code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//SetHeader 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//String 返回字符串响应内容
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text-plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//JSON 返回json响应内容
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//Data 返回状态码和数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

//HTML 返回html响应信息
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

//Fail 中间件测试
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}
