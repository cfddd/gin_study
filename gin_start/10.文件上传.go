package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()

	// 单文件上传
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		{
			fmt.Println(file.Filename)
			fmt.Println(file.Header)
			fmt.Println(file.Size) //单位是字节
		}

		// 文件对象  文件路径，注意要从项目根路径开始写
		c.SaveUploadedFile(file, "./gin_start/uploads/cfd.html")

		//打开文件，读取文件中的数据
		readerFile, _ := file.Open()
		data, _ := io.ReadAll(readerFile)
		fmt.Println(string(data))

		//打开文件，在本地创建一个文件，然后拷贝上传文件的内容
		writerFile, _ := os.Create("./gin_start/uploads/cfd.html")
		defer writerFile.Close()
		n, _ := io.Copy(writerFile, readerFile)
		fmt.Println(n)

		c.JSON(200, gin.H{"msg": "上传成功"})
	})

	// 多文件上传
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			c.SaveUploadedFile(file, "./gin_start/uploads/"+file.Filename)
		}
		c.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传 %d 个文件", len(files))})
	})

	router.Run(":8080")
}
