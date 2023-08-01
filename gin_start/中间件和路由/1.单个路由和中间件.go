package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func f0(c *gin.Context) {
	fmt.Println("f0 << in")
	fmt.Println("0")
	c.Next()
	fmt.Println("f0 >> out")
}
func index(c *gin.Context) {
	fmt.Println("f1 << in")
	fmt.Println("1")
	c.JSON(200, gin.H{"msg": "index"})
	c.Next()

	fmt.Println("f1 >> out")
}
func f2(c *gin.Context) {
	fmt.Println("f2 << in")
	fmt.Println("2")
	/*
		c.Abort()拦截，后续的HandlerFunc就不会执行了
		c.Next()，Next前后形成了其他语言中的请求中间件和响应中间件
		如果其中一个中间件响应了c.Abort()，后续中间件将不再执行，直接按照顺序走完所有的响应中间件
		中间件会倒序执行所有响应中间件
	*/
	c.Abort()
	c.Next()
	fmt.Println("f2 >> out")
}
func f3(c *gin.Context) {
	fmt.Println("f3 << in")
	fmt.Println("3")
	fmt.Println("f3 >> out")

}

func main() {
	router := gin.Default()

	router.GET("/", f0, index, f2, f3)

	router.Run(":8080")
}
