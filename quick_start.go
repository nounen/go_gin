package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()

	// quick start
	router.GET("/ping", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	// parameters in path
	// 路由正则更多案例: `"/user/:name/*action"`
	router.GET("/user/:name", func (c *gin.Context)  {
		name := c.Param("name")

		c.String(http.StatusOK, "hello %s", name)
	})


	// querystring parameters 
	// eg: http://localhost:8080/welcome?firstname=lin&lastname=moumou
	router.GET("/welcome", func (c *gin.Context)  {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})


	// Multipart/Urlencoded Form (建议使用 postman 把数据发过来)
	router.POST("/form_post", func (c *gin.Context)  {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	// Upload files
	/*
	curl -X POST http://localhost:8080/upload \
		-F "file=@./README.md" \
		-H "Content-Type: multipart/form-data"
	*/
	router.POST("/upload", func (c *gin.Context)  {
		file, _ := c.FormFile("file")

		log.Println(file.Filename)

		c.SaveUploadedFile(file, "upload_README.md")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
