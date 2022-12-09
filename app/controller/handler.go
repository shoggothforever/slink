package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shortlink/dao"
	"shortlink/model"
	"time"
)

/*
handler for /user
*/
func Register(c *gin.Context) {
	var user model.User
	//var login model.LoginInfo
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Pwd = c.PostForm("pwd")
	if err := SaveUser(&user); err == nil {
		c.JSON(200, model.RegisterResponse{
			model.Response{200, "注册成功"},
			user,
		})
		model.CurrentUser = user
	} else if err == gorm.ErrRegistered {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "该用户或者邮箱已存在",
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
	var login model.LoginInfo
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
		if err := SaveLogin(&login); err == nil {

			c.JSON(200, model.LoginResponse{
				model.Response{200, "登陆成功"},
				model.CurrentUser.GetId(),
			})
		} else {
			model.CurrentUser = data[0]
			c.JSON(200, model.LoginResponse{
				model.Response{201, "登陆成功,写入数据失败"},
				model.CurrentUser.GetId(),
			})
		}
		SaveJwt(model.CurrentUser.Name, model.CurrentUser.Pwd)
		fmt.Println(model.AuthClaims)
		fmt.Println(model.AuthToken)
	}
}

// logout
func Logout(c *gin.Context) {
	model.CurrentUser = model.User{}
	model.CurrentUser.Id = -1
	c.JSON(200, model.Response{
		200, "退出成功",
	})
}

// info
func GetInfo(c *gin.Context) {
	c.JSON(200, model.InfoResponse{
		model.Response{200, "当前用户信息:"},
		model.CurrentUser.GetId(),
		model.CurrentUser.Name,
		model.CurrentUser.Email,
	})

}

// record/get
func GetLoginInfo(c *gin.Context) {
	var infos []model.LoginRecord
	id := model.CurrentUser.GetId()
	dao.Db.Raw("select id,user_id,login_at from login_infos where user_id=? limit 0,10", id).Find(&infos)
	if len(infos) == 0 {
		logrus.Info("no login document")
		c.JSON(200, model.Response{
			400,
			"没有此链接",
		})
	} else {
		c.JSON(200, model.LoginInfoResponse{
			model.Response{200, "查找到信息"},
			infos,
		})
	}
}

// url/get
func GetUrl(c *gin.Context) {
	var urls []model.UrlInfo
	id := model.CurrentUser.GetId()
	dao.Db.Raw("select * from url_infos where user_id=(select id from users where id=?)", id).Find(&urls)
	if len(urls) == 0 {
		logrus.Info("no such document")
		c.JSON(200, model.Response{
			400,
			"该用户没有链接记录",
		})
	} else {
		c.JSON(200, model.QueryResponse{
			model.Response{200, "查找到信息"},
			urls,
		})
	}
}

/*
handler for /url
*/
func Create(c *gin.Context) {
	var url model.UrlInfo
	url.Origin = c.PostForm("origin")
	url.Short = GenShort(c.PostForm("short"))
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
func Query(c *gin.Context) {
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
	var url model.UrlInfo
	id := model.CurrentUser.GetId()
	url.Origin = c.PostForm("origin")
	short := c.PostForm("oldshort")
	url.Short = GenShort(c.PostForm("newshort"))
	url.Comment = c.PostForm("comment")
	dao.Db.Model(model.UrlInfo{}).Where("user_id=? and origin=? and short=?", id, url.Origin, short).Updates(map[string]interface{}{
		"short":       url.Short,
		"comment":     url.Comment,
		"start_time":  time.Now().In(time.Local),
		"expire_time": time.Now().Add(24 * time.Hour).In(time.Local),
	})
	c.JSON(200, model.UpdateResponse{
		model.Response{200, "更新成功"},
	})
}
func Delete(c *gin.Context) {

	var url []model.UrlInfo
	id := model.CurrentUser.GetId()
	origin := c.PostForm("origin")
	short := c.PostForm("short")
	dao.Db.Where("user_id=? and origin=? and short=?", id, origin, short).Find(&url)
	if len(url) == 0 {
		c.JSON(200, model.Response{
			404, "表中没有该数据",
		})
	} else {
		dao.Db.Model(model.UrlInfo{}).Delete(&url)
		c.JSON(200, model.UpdateResponse{
			model.Response{200, "删除成功"},
		})
	}
}
func Pause(c *gin.Context) {

}
