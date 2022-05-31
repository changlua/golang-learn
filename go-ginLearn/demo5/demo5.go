package demo5

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context)  {
	defer func() {
		if r := recover(); r != nil{
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			//封装通用JSON返回
			c.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg": errorToString(r),
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续的插件及代码执行
	c.Next()
}
//错误转字符串
func errorToString(r interface{}) string {
	switch v := r.(type) {
		case error:
			return v.Error()
		default:
			return r.(string)
	}
}

func Main()  {
	r := gin.Default()
	//使用全局异常捕捉插件:Recover 要尽量放在第一个被加载
	r.Use(Recover)
	r.GET("/test", func(c *gin.Context) {
		// 无意抛出 panic
		var slice = []int{1, 2, 3, 4, 5}
		slice[6] = 6
	})

	r.Run()
}
