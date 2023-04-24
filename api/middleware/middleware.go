package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shortlink/dal"
	"shortlink/model"
)

// 使用中间件实现重定向
func RedirectShort() gin.HandlerFunc {
	return func(c *gin.Context) {
		short := c.Request.URL.String()
		fmt.Println(short)
		short = short[1:]
		var urls []model.UrlInfo
		var purl model.PauseUrl
		purl.Short = ""
		err := dal.Getdb().Where("short=?", short).First(&purl).Error
		if err == nil {
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "短链接已冻结，请解冻后再试",
			})
		} else if err != nil && err != gorm.ErrRecordNotFound {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "重定向服务失败",
			})
		} else {
			dal.Getdb().Where("short=?", short).Find(&urls)
			if len(urls) != 0 {
				c.Redirect(302, urls[0].Origin)
			} else {
				//c.JSON(200, gin.H{"code": 404, "msg": "请输入正确的短链接"})
			}
		}
		c.Next()
	}
}

// 使用中间件实现鉴权
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
		dal.Getdb().Raw("select * from users where id =(select user_id from cookies where user_id=id)").First(&cur)
		c.Set("user", cur)
		var authjwt string
		dal.Getdb().Raw("select jwt from cookies where user_id = ?", cur.Id).First(&authjwt)
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
			c.Set("user", cur)
			c.Next()
		}

	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400")
			c.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
