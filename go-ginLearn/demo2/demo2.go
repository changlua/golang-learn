package demo2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Main()  {
	r := gin.Default()
	//方式一：使用JSON工具来进行序列化
	r.POST("/user", func(c *gin.Context) {
		//1、定义一个Map来进行接收
		requestMap := make(map[string]string)
		//2、进行JSON序列化（使用JSON工具类）
		if err := json.NewDecoder(c.Request.Body).Decode(&requestMap); err != nil{
			panic(err)
		}
		c.JSON(200, gin.H{
			"code": 200,
			"result": requestMap,
		})
	})
	//方式二：使用gin自带的bind方法来进行（底层会根据对应的类型来判断进行序列化，建议用框架带的更方便）
	r.POST("/user2", func(c *gin.Context) {
		//1、定义一个user对象
		user := User{}
		//2、使用gin的bind来进行序列化
		if err := c.BindJSON(&user); err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"code": 200,
			"result": user,
		})
	})
	r.Run()
}