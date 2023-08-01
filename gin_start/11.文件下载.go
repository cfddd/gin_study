package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 直接响应一个路径下的文件下载
	router.GET("/download", func(c *gin.Context) {
		//直接下载
		c.Header("Content-Type", "application/octet-stream")                 // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Disposition", "attachment; filename="+"README.md") // 用来指定下载下来的文件名
		//c.Header("Content-Transfer-Encoding", "binary")                   // 表示传输过程中的编码形式，乱码问题可能就是因为它

		//有些响应，比如图片，浏览器就会显示这个图片，而不是下载，所以我们需要使浏览器唤起下载行为
		c.File("gin_start/uploads/README.md")

	})

	router.Run(":8080")
}
