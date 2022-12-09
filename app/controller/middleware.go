package controller

import (
	"github.com/gin-gonic/gin"
	"shortlink/model"
)

//使用中间件实现重定向
func RedirectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

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
