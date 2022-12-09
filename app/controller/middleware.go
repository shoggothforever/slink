package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shortlink/dao"
	"shortlink/model"
)

//使用中间件实现重定向
func RedirectShort() gin.HandlerFunc {
	return func(c *gin.Context) {
		short := c.Request.URL.String()
		fmt.Println(short)
		short = short[1:]
		var urls []model.UrlInfo
		dao.Db.Where("short=?", short).Find(&urls)
		if len(urls) != 0 {
			c.Redirect(301, urls[0].Origin)
		} else {
			fmt.Println("不存在对应的短链接")
		}
		c.Next()
	}
}

//使用中间件实现赋值
func ADDInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

//使用中间件实现鉴权
func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if model.CurrentUser.Id == model.NOTLOGIN {
			c.AbortWithStatusJSON(404, model.Response{404, "请登录后再试"})
		} else {
			//if(c.PostForm("token")!=model.AuthToken){
			//	c.JSON(200, model.Response{404, "验证错误"})
			//	c.AbortWithStatus(404)
			//}

			c.Next()
		}
	}
}

//func CookieMiddleWare() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		cookieValue, err := c.Cookie("Jwt")
//		if err != nil {
//			//c.AbortWithStatusJSON(403, gin.H{"message": "get cookie failed..."})
//			fmt.Printf("err happened :%v\n", err)
//			code := _
//			c.SetCookie("Jwt", code, 3600*24, "/", "http://localhost:9090", false, false)
//		}
//		fmt.Println(cookieValue)
//		c.Next()
//	}
//}
