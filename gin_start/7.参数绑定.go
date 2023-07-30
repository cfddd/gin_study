package main

import (
	"github.com/gin-gonic/gin"
)

type SignUserInfo struct {
	/*
	   常用验证器
	   // 不能为空，并且不能没有这个字段
	   required： 必填字段，如：binding:"required"

	   // 针对字符串的长度
	   min 最小长度，如：binding:"min=5"
	   max 最大长度，如：binding:"max=10"
	   len 长度，如：binding:"len=6"

	   // 针对数字的大小
	   eq 等于，如：binding:"eq=3"
	   ne 不等于，如：binding:"ne=12"
	   gt 大于，如：binding:"gt=10"
	   gte 大于等于，如：binding:"gte=10"
	   lt 小于，如：binding:"lt=10"
	   lte 小于等于，如：binding:"lte=10"

	   // 针对同级字段的
	   eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
	   nefield 不等于其他字段的值


	   - 忽略字段，如：binding:"-"

	*/
	Name       string `json:"name" binding:"required,max=6,min=4" ` //必填的选项
	Age        int    `json:"age" binding:"gt=0,lt=250"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"eqfield=Password"`

	/*
		// 枚举  只能是red 或green
		oneof=red green

		// 字符串
		contains=fengfeng  // 包含fengfeng的字符串
		excludes // 不包含
		startswith  // 字符串前缀
		endswith  // 字符串后缀

		// 数组
		dive  // dive后面的验证就是针对数组中的每一个元素

		// 网络验证
		ip
		ipv4
		ipv6
		uri
		url
		// uri 在于I(Identifier)是统一资源标示符，可以唯一标识一个资源。
		// url 在于Locater，是统一资源定位符，提供找到该资源的确切路径

		// 日期验证  1月2号下午3点4分5秒在2006年
		datetime=2006-01-02
	*/
	Sex      string   `json:"sex" binding:"oneof=man woman"`
	LikeList []string `json:"like_list" binding:"required,dive,startswith=like"`
	IP       string   `json:"ip" binding:"required,ip"'`
	Url      string   `json:"url" binding:"url"'` //url是uri的子集
	Uri      string   `json:"uri" binding:"uri"`
	Date     string   `json:"date" time_format:"2020-01-02 15:04:05"`
}

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {

		var user SignUserInfo

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user})
	})

	router.Run(":8080")
}
