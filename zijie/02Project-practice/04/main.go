package main

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
	"zijie/02Project-practice/04/controller"
)

func main()  {
	//初始化数据源

	//gin初始化（web框架）
	r := gin.Default()
	r.GET("/topic/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		pageInfo := controller.QueryPageInfo(topicId)
		c.JSON(200, pageInfo)  //JSON序列化返回
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func initMemory()  {
	//暂不进行初始化
}