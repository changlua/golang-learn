package demo1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main()  {
	//使用默认中间件创建一个gin路由器
	r := gin.Default()
	//1、常见的方法
	//get
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//get-1：获取路径中的参数【注意：无法匹配/user/,/user这类路径】 //示例：/hello/changlu
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")  //获取url中的相应参数
		c.String(http.StatusOK, "hello %s", name)
	})
	//get-2：获取get参数  //示例：/welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest") //若是没有获取到，可获取到默认值。【底层还是走query方法】
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})
	//post
	r.POST("/ping", commonResponse)
	//post-1：获取form表单
	r.POST("/form_post", func(c *gin.Context) {
		//获取表单，和get的query类似
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	//put
	r.PUT("/ping", commonResponse)
	//其他方法：delete、patch、head、options
	//2、启动服务
	//r.Run()  //默认启动服务以及监听：127.0.0.1::8080
	r.Run(":3000")  //指定端口
}

func commonResponse(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}