package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func m10(c *gin.Context) {
	fmt.Println("m10 <<<<<< in")
	c.Next()
	fmt.Println("m10 <<<<<< out")
}

func m11(c *gin.Context) {
	c.Set("age", 100)
	fmt.Println("m11 <<<<<< in")
	c.Next()
	fmt.Println("m11 <<<<<< out")
}

func m12(c *gin.Context) {
	fmt.Println("m12 <<<<<< in")
	c.Set("name", "CFD")
	c.Set("user", User{Name: "CFD", Age: 20})
}

func main() {
	router := gin.Default()

	//全局中间件
	router.Use(m10, m11)

	// 使用全局中间件
	router.GET("/m10", func(c *gin.Context) {
		fmt.Println("index <<< in")
		c.JSON(200, gin.H{"msg": "m10"})
		c.Next()
		fmt.Println("index <<< out")
	})

	// 使用abort的顺序
	router.GET("/m11", func(c *gin.Context) {
		fmt.Println("index <<< in")
		c.JSON(200, gin.H{"msg": "m11 abort()吞掉了你的响应，但没全局响应还在，说明全局中间件在最先执行"})
		c.Abort()
		c.Next()
		fmt.Println("index <<< out")
	})

	// 中间件之间的消息传递
	// 先执行了才有消息，全局变量可以先有消息
	router.GET("/m12", m12, func(c *gin.Context) {
		name, _ := c.Get("name")
		user, _ := c.Get("user")
		c.JSON(200, gin.H{"msg": name, "msg1": user})

		//user.Name//错误，因为user是any类型
		//可以断言转换
		_user := user.(User)
		fmt.Println(_user.Age)
	})

	router.Run(":8080")
}
