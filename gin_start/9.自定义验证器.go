package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type User struct {
	//在binding中添加了sign验证器，调用自己的判断函数
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"`
	Age  int    `json:"age" binding:"required" msg:"年龄校验失败"`
}

func _GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传obj的指针
	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}

// 如果用户名等于nameSt就校验失败
// 自己的判断函数
func signValid(fl validator.FieldLevel) bool {
	var nameSt string = "cfddfc"
	name := fl.Field().Interface().(string)
	if name == nameSt {
		return false
	}
	return true

}

func main() {
	router := gin.Default()

	//自定义验证器
	//这里添加了一个sign
	//在上面的结构体中binding里面有相关信息
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.POST("/", func(c *gin.Context) {

		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
		return
	})
	router.Run(":8080")
}
