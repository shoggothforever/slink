package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"shortlink/dao"
	"shortlink/model"
)

/*
handler for /user
*/
func Register(c *gin.Context) {
	var user model.User
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Pwd = c.PostForm("pwd")
	if err := SaveUser(&user); err == nil {
		c.JSON(200, model.RegisterResponse{
			model.Response{200, "注册成功"},
			user,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "注册失败",
		})
	}
}

// login
func Login(c *gin.Context) {
	var user model.User
	user.Name = c.PostForm("name")
	user.Pwd = c.PostForm("pwd")
	var data []model.User
	dao.Db.Where("name=? AND pwd=?", user.Name, user.Pwd).First(&data)
	if len(data) == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "登录失败，请输入正确的账户名和密码",
		})
	} else {
		model.CurrentUser = data[0]
		c.JSON(200, model.LoginResponse{
			model.Response{200, "登陆成功"},
			model.CurrentUser,
		})
	}
}

// logout
func Logout(c *gin.Context) {
	model.CurrentUser = model.User{}
	c.JSON(200, model.LoginResponse{
		model.Response{200, "退出成功"},
		model.CurrentUser,
	})
}

// info
func GetInfo(c *gin.Context) {

}

// record/get
func GetLoginInfo(c *gin.Context) {

}

// url/get
func GetUrl(c *gin.Context) {

}

/*
handler for /url
*/
func Create(c *gin.Context) {
	var url model.UrlInfo
	url.Origin = c.PostForm("origin")
	url.Short = c.PostForm("short")
	url.Comment = c.PostForm("comment")
	if err := SaveUrl(&url); err == nil {
		c.JSON(200, gin.H{
			"code":    200,
			"msg":     "链接信息存储成功",
			"urlinfo": url,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "链接信息存储失败",
		})
	}
}
func Querry(c *gin.Context) {
	var urls []model.UrlInfo
	id := c.PostForm("id")
	dao.Db.Raw("select * from url_infos where url_infos.user_id=(select id from users where id=?)", id).First(&urls)
	if len(urls) == 0 {
		logrus.Info("no such document")
		c.JSON(200, model.Response{
			400,
			"没有此链接",
		})
	} else {
		c.JSON(200, model.QueryResponse{
			model.Response{200, "查找到信息"},
			urls,
		})
	}
}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
func Pause(c *gin.Context) {

}
func Shorten(c *gin.Context) {

}
