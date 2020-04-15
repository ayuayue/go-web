//

package main

import (
	"gee"
	"net/http"
)

//onlyForV2 测试日志中间件
// func onlyForV2() gee.HandlerFunc {
// 	return func(c *gee.Context) {
// 		t := time.Now()
// 		c.Fail(500, "Internal Server Error")
// 		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
// 	}
// }

// type student struct {
// 	Name string
// 	Age  int
// }

// func formatAsData(t time.Time) string {
// 	year, month, day := t.Date()
// 	return fmt.Sprintf("%d~$02d-%02d", year, month, day)
// }
func main() {
	r := gee.New()
	// r.Static("/assets", "./static")
	// r.Use(gee.Logger())
	// r.SetFuncMap(template.FuncMap{
	// 	"formatAsDate": formatAsData,
	// })
	// r.LoadHTMLGlob("templates/*")
	// r.Static("/assets", "./static")
	// stu1 := &student{Name: "caoayu", Age: 18}
	// stu2 := &student{Name: "ayu", Age: 18}
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "css.tmpl", nil)
	// })
	// r.GET("/students", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "arr.tmpl", gee.H{
	// 		"title":  "caoayu",
	// 		"stuArr": [2]*student{stu1, stu2},
	// 	})
	// })
	// r.GET("/date", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
	// 		"title": "caoayu",
	// 		"now":  time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
	// 	})
	// })
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// })
	// r.GET("/hello", func(c *gee.Context) {
	// 	// expect /hello?name=geektutu
	// 	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	// })

	// r.POST("/login", func(c *gee.Context) {
	// 	c.JSON(http.StatusOK, gee.H{
	// 		"username": c.PostForm("username"),
	// 		"password": c.PostForm("password"),
	// 	})
	// })
	// r.GET("/hello/:name", func(c *gee.Context) {
	// 	c.String(http.StatusOK, "hello %s ,you're at %s\n", c.Param("name"), c.Path)
	// })
	// r.GET("/assets/*filepath", func(c *gee.Context) {
	// 	c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	// })
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/", func(c *gee.Context) {
	// 		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	// 	})
	// 	v1.GET("/hello", func(c *gee.Context) {
	// 		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
	// 	})
	// }
	// v2 := r.Group("/v2")
	// v2.Use(onlyForV2())
	// {
	// 	v2.GET("/hello/:name", func(c *gee.Context) {
	// 		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
	// 	})
	// 	v2.POST("/login", func(c *gee.Context) {
	// 		c.JSON(http.StatusOK, gee.H{
	// 			"username": c.PostForm("username"),
	// 			"password": c.PostForm("password"),
	// 		})
	// 	})
	// }
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello caoayu\n")
	})
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"caoayu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
