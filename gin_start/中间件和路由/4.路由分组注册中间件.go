package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Rse struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type _User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

//func middle(c *gin.Context, usermsg string) {
//	token := c.GetHeader("token")
//	if token == "1234" {
//		//成功了就继续走后面的数据，然后走回来的时候再在这里退出
//		c.Next()
//		return
//	}
//	//失败了就返回
//	c.JSON(200, Rse{1001, nil, "权限验证失败"})
//	c.Abort()
//}

func _middle(msg string) gin.HandlerFunc {
	//闭包,非常实用
	//多接一层，可以给中间件函数传递参数
	fmt.Println("这里的代码立即执行")

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "1234" {
			//成功了就继续走后面的数据，然后走回来的时候再在这里退出
			c.Next()
			return
		}
		//失败了就返回
		c.JSON(200, Rse{1001, nil, msg + "权限验证失败"})
		c.Abort()
	}

}
func UseUserRouter(router *gin.RouterGroup) {
	//增加统一中间件
	userHandle := router.Group("/user").Use(_middle("用户CFD"))
	userHandle.GET("/CFD", func(c *gin.Context) {
		c.JSON(200, Rse{1001, []_User{
			{"CFD", 500, "male"},
			{"YMT", 5000, "female"},
		}, "权限验证失败"})
	})
}

func main() {
	router := gin.Default()
	api := router.Group("/api")

	api.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "1234"})
	})

	UseUserRouter(api)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
