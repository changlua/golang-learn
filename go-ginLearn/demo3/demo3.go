package demo3

import (
	"github.com/gin-gonic/gin"
	"go-ginLearn/common"
	"log"
	"mime/multipart"
)

var (
	dst = "C:\\Users\\93997\\Desktop\\upload\\755b7cdd-ed6c-43c3-b62f-14a5c9e167a1.png"
)

func Main()  {
	r := gin.Default()
	//1、上传单个文件
	//表单限制上传大小（默认32MiB）
	//r.MaxMultipartMemory = 8 << 20  // 8Mib
	r.POST("/upload", func(c *gin.Context) {
		//单文件
		file, _ := c.FormFile("file") //直接取出key为file的文件
		log.Println(file.Filename)
		//保存文件
		saveFile(file, c)
		c.JSON(200, gin.H{
			"code": 200,
			"message": "上传成功",
		})
	})
	//2、上传多个文件
	r.POST("/uploads", func(c *gin.Context) {
		//1、取出文件数组
		form, _ := c.MultipartForm()
		files := form.File["file"]  //根据上传的key来取出对应的文件数组
		//2、遍历保存
		for _ , file := range files {
			saveFile(file, c)
		}
		c.JSON(200, gin.H{
			"code": 200,
			"message": "上传成功",
		})
	})
	r.Run()
}

//保存文件
func saveFile(file *multipart.FileHeader, c *gin.Context)  {
	//上传文件到制定目录
	targetPath := dst + common.GenerateFileName(file.Filename)
	if err := c.SaveUploadedFile(file, targetPath); err != nil {
		panic(err)
	}
}
