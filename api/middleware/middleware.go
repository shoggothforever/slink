package middleware

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
		var purl model.PauseUrl
		purl.Short = ""
		dao.Getdb().Where("short=?", short).First(&purl)
		if purl.Short != "" {
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "短链接已冻结，请解冻后再试",
			})
		} else {
			dao.Getdb().Where("short=?", short).Find(&urls)
			if len(urls) != 0 {
				c.Redirect(301, urls[0].Origin)
			} else {
				//c.JSON(200, gin.H{"code": 404, "msg": "请输入正确的短链接"})
			}
		}
		c.Next()
	}
}

//使用中间件实现鉴权
func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("Authorization")
		if len(jwt) > 7 { //为Bearer Token去除前七位数据
			jwt = jwt[7:]
		} else {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 403, "msg": "请登录后再试",
			})
			return
		} //如果在PostMan中使用 Bearer Token 会在jwt前加上bearer: 前缀
		var cur model.User
		dao.Getdb().Raw("select * from users where id =(select user_id from cookies where user_id=id)").First(&cur)
		c.Set("user", cur)
		var authjwt string
		dao.Getdb().Raw("select jwt from cookies where user_id = ?", cur.Id).First(&authjwt)
		if authjwt == "" {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 403, "msg": "请登录后再试",
			})
			return
		}
		if jwt != authjwt {
			c.Set("AUthInfo", "Failed!")
			c.AbortWithStatusJSON(200, gin.H{
				"code": 403, "msg": "请输入正确的信息",
			})
			return
		} else {
			c.Set("AuthInfo", "Success!")
			c.Next()
		}

	}
}
