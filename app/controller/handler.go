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
	SaveUser(&user)
}

// login
func Login(c *gin.Context) {

}

// logout
func Logout(c *gin.Context) {

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
