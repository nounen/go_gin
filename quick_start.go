package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// quick start
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// parameters in path
	// 路由正则更多案例: `"/user/:name/*action"`
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.String(http.StatusOK, "hello %s", name)
	})

	// querystring parameters
	// eg: http://localhost:8080/welcome?firstname=lin&lastname=moumou
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// Multipart/Urlencoded Form (建议使用 postman 把数据发过来)
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// Upload files
	/*
		curl -X POST http://localhost:8080/upload \
			-F "file=@./README.md" \
			-H "Content-Type: multipart/form-data"
	*/
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		log.Println(file.Filename)

		c.SaveUploadedFile(file, "upload_README.md")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// Multiple files
	// 多文件上传. 略

	// Grouping routes
	v1 := router.Group("/v1")

	v1.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.String(http.StatusOK, "v1 route group: %s", name)
	})

	// 控制器 -- 数据解析绑定
	// 带标签的结构体
	type Login struct {
		User     string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	// 绑定JSON的例子 ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login

		// 根据请求头中 content-type 自动推断: c.Bind(&form)

		if c.BindJSON(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	// 响应 -- JSON/XML/YAML响应
	router.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user" xml:"user"`
			Message string
			Number  int
		}

		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"user": "Lena", "Message": "hey", "Number": 123}
		// c.JSON(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		// c.XML(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		// c.YAML(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		c.JSON(http.StatusOK, msg)
		// c.XML(http.StatusOK, msg)
		// c.YAML(http.StatusOK, msg)
	})

	// 1. isten and serve on 0.0.0.0:8080
	router.Run()

	// 2. 除了默认服务器中 router.Run() 的方式外，还可以用 http.ListenAndServe()，比如
	// http.ListenAndServe(":8080", router)

	// 2. 或者自定义 HTTP 服务器的配置：
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

	// 3. HTTP 服务器替换方案 想无缝重启、停机吗? 以下有几种方式：
	// 可以使用 fvbock/endless 来替换默认的 ListenAndServe https://github.com/fvbock/endless
	// endless.ListenAndServe(":4242", router)
}
