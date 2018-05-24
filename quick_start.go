package main

import (
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

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
