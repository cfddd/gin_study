package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

// 查询参数
func _query(c *gin.Context) {
	//http://localhost:8080/query?user=%E5%B8%B8%E5%87%A4%E8%BF%AA&user=cfd&user=23&tiem=key
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user"))          // 拿到多个相同的查询参数
	fmt.Println(c.DefaultQuery("addr", "四川省")) //默认参数
	//http://localhost:8080/query?user=%E5%B8%B8%E5%87%A4%E8%BF%AA&addr=apple

}

// 动态参数
func _param(c *gin.Context) {
	//http://localhost:8080/param/CFD/LH
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

// 表单参数
func _form(c *gin.Context) {
	//使用postman
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "四川省")) // 如果用户没传，就使用默认值
	forms, err := c.MultipartForm()               // 接收所有的form参数，包括文件
	fmt.Println(forms, err)
}

// 原始参数
func _raw(c *gin.Context) {

	type User struct {
		Name string `json:"name"`
		Addr string `json:"addr"`
		Age  int    `json:"age"`
	}
	var user User
	err := bindJson(c, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	//body, _ := c.GetRawData()
	//fmt.Println(string(body))
	////
	//contentType := c.GetHeader("Content-Type")
	//switch contentType {
	//case "application/json":
	//
	//	// json解析到结构体
	//	type User struct {
	//		Name string `json:"name"`
	//		Age  int    `json:"age"`
	//	}
	//	var user User
	//	err := json.Unmarshal(body, &user)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(user)
	//}
}

func bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func main() {
	router := gin.Default()

	router.GET("/query", _query)

	router.GET("/param/:user_id/:book_id", _param)

	router.POST("/form", _form)

	router.POST("/raw", _raw)

	router.Run(":8080")
}
