package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(http.StatusOK, "返回txt")
}

func _json1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func _json2(c *gin.Context) {
	// You also can use a struct
	type Msg struct {
		Name    string `json:"user"`
		Message string `json:"message"`
		Number  int    `json:"number"`
		passwd  string `json:"-"` //ignore json
	}
	msg := Msg{"CFD", "hey", 21, "123456"}
	// 注意 msg.Name 变成了 "user" 字段
	// 以下方式都会输出 :   {"user": "hanru", "Message": "hey", "Number": 123}
	c.JSON(http.StatusOK, msg)
}

func _json3(c *gin.Context) {
	userMap := map[string]string{
		"userName": "CFD",
		"age":      "100",
	}
	c.JSON(http.StatusOK, userMap)
}

func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "CFD", "message": "hey", "status": http.StatusOK, "data": gin.H{"gender": "male"}})
}

func _yaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "CFD", "message": "hey", "status": http.StatusOK, "data": gin.H{"gender": "male"}})
}

func _html(c *gin.Context) {
	//传递"hhh"的值"cfd"给html文件里面使用
	c.HTML(http.StatusOK, "index.html", gin.H{"hhh": "cfd"})
}

func _html1(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.html", gin.H{
		"title": "Posts",
	})
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"title": "Users",
	})
}

func _redirect(c *gin.Context) {
	//c.Redirect(301,"https://www.cfd.com")         //301 Moved Permanently
	c.Redirect(302, "http://www.cfd.com") //302 Found
	//更多状态码，网上搜
}

func main() {
	// 创建一个默认的路由
	router := gin.Default()

	//响应字符串
	router.GET("/txt", _string)

	//响应json
	{
		// 直接响应json
		router.GET("/json", _json1)
		// 结构体转json
		router.GET("/moreJSON", _json2)
		// Map转json
		router.GET("/mapJSON", _json3)
	}

	// 响应xml
	router.GET("/xmlJSON", _xml)

	//响应yaml
	router.GET("/yamlJSON", _yaml)

	//响应html
	//router.LoadHTMLGlob("gin_start/template/*")           //加载目录下所有文件
	router.LoadHTMLFiles("gin_start/template/index.html") //加载一个文件
	router.GET("/index", _html)

	//响应文件
	//goland中没有相对当前文件的路径，只有相对当前项目根路径的路径
	//router.StaticFile(访问路径，文件路径)
	router.StaticFile("/static/CF", "gin_start/static/CFD.jpg")
	//router.StaticFS(访问路径，后缀文件路径)
	//使用方法为访问路径+后缀文件路径中的具体文件
	//因为可以为空，所以不允许有其他相同的”访问路径“，即前缀不能相同
	//路径之上的文件是无法访问的
	router.StaticFS("/gin_start/static", http.Dir("gin_start/static/static"))

	//redirect
	router.GET("/cfd", _redirect)

	// 启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	router.Run("0.0.0.0:8080")
	// 用原生http服务的方式， router.Run本质就是http.ListenAndServe的进一步封装
	http.ListenAndServe(":8080", router)
}
