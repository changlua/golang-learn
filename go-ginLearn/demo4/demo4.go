package demo4

import "github.com/gin-gonic/gin"

func Main()  {
	r := gin.Default()
	//v1的API
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			commonResponse(c, "/v1/login")
		})
		v1.POST("/register",func(c *gin.Context) {
			commonResponse(c, "/v1/register")
		})
	}

	//v2的API
	v2 := r.Group("v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			commonResponse(c, "/v2/login")
		})
		v2.POST("/register", func(c *gin.Context) {
			commonResponse(c, "/v2/register")
		})
	}

	//其他无对应路由时走的接口
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": "404",
			"message": "page not found",
		})
	})

	r.Run()
}

func commonResponse(c *gin.Context, url string)  {
	c.JSON(200, gin.H{
		"code": 200,
		"url": url,
	})
}