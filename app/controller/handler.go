package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shortlink/dao"
	"shortlink/model"
	"strconv"
	"time"
)

/*
handler for /user
*/
func Register(c *gin.Context) {
	var user model.User
	//var login model.LoginInfo
	user.Name = c.PostForm("name")
	if user.Name == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入用户名"})
	}
	user.Email = c.PostForm("email")
	if user.Name == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入邮箱"})
	}
	user.Pwd = c.PostForm("pwd")
	if user.Name == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入密码"})
	}
	if user.Name == "" || user.Email == "" || user.Pwd == "" {
		return
	}
	if err := SaveUser(&user); err == nil {
		c.JSON(200, model.RegisterResponse{
			model.Response{200, "注册成功"},
			user,
		})
	} else if err == gorm.ErrRegistered {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "该用户或者邮箱已存在",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "注册失败",
		})
	}
}

// login
func Login(c *gin.Context) {
	var user model.User
	var login model.LoginInfo
	user.Name = c.PostForm("name")
	if user.Name == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入用户名"})
	}
	user.Pwd = messagedigest5(c.PostForm("pwd"))

	if user.Name == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入密码"})
	}
	if user.Name == "" || user.Pwd == "" {
		return
	}
	var data []model.User
	dao.Getdb().Where("name=? AND pwd=?", user.Name, user.Pwd).First(&data)
	if len(data) == 0 {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "登录失败，请输入正确的账户名和密码",
		})
	} else {
		if err := SaveLogin(&login, data[0].Id); err == nil {
			cur := data[0]
			SaveJwt(cur.Id, cur.Name)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "登陆成功",
				"id":   cur.GetId(),
				"jwt":  model.AuthJwt,
			})
		}
	}
}

// logout
func Logout(c *gin.Context) {
	var cur model.User
	cur.Id = -1
	c.Set("user", cur)
	c.JSON(200, model.Response{
		200, "退出成功",
	})
}

// info
func GetInfo(c *gin.Context) {
	cur, _ := getcuruser(c)
	c.JSON(200, model.InfoResponse{
		model.Response{200, "当前用户信息:"},
		cur.GetId(),
		cur.Name,
		cur.Email,
	})

}

// record/get
func GetLoginInfo(c *gin.Context) {
	var infos []model.LoginRecord
	cur, _ := getcuruser(c)
	id := cur.GetId()
	page := 0
	page, _ = strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	page = (page - 1) * 10
	dao.Getdb().Raw("select id,user_id,login_at from login_infos where user_id=? limit ?,10", id, page).Find(&infos)
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
	cur, _ := getcuruser(c)
	var urls []model.UrlInfo
	id := cur.GetId()
	page := 0
	page, _ = strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	page = (page - 1) * 10
	dao.Getdb().Raw("select * from url_infos where user_id=(select id from users where id=?) limit ?,10", id, page).Find(&urls)
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
	cur, _ := getcuruser(c)
	var url model.UrlInfo
	url.Origin = c.PostForm("origin")
	if url.Origin == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入需要转换的链接"})
		return
	}
	url.Short = GenShort(c.PostForm("short"))
	url.Comment = c.PostForm("comment")
	if err := SaveUrl(&url, cur.Id); err == nil {
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
	cur, _ := getcuruser(c)
	var url []model.UrlInfo
	id := c.PostForm("id")
	user_id := cur.Id
	if id == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入需要查找的链接ID"})
		return
	}
	//dao.Getdb().Raw("select * from url_infos where url_infos.user_id=(select id from users where id=?)", id).First(&urls)
	dao.Getdb().Model(&model.UrlInfo{}).Where("user_id=? and id=?", user_id, id).First(&url)
	if len(url) == 0 {
		logrus.Info("no such document")
		c.JSON(200, model.Response{
			400,
			"没有此链接",
		})
	} else {
		c.JSON(200, model.QueryResponse{
			model.Response{200, "查找到信息"},
			url,
		})
	}
}
func Update(c *gin.Context) {
	cur, _ := getcuruser(c)
	var url model.UrlInfo
	url.Origin = ""
	userid := cur.GetId()
	id := c.PostForm("id")
	if id == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入需要查找的链接ID"})
		return
	}
	url.Short = GenShort(c.PostForm("newshort"))
	url.Comment = c.PostForm("comment")
	dao.Getdb().Model(model.UrlInfo{}).Where("user_id=? and id=?", userid, id).First(&url)
	if url.Origin != "" {
		dao.Getdb().Model(url).Updates(map[string]interface{}{
			"short":       url.Short,
			"comment":     url.Comment,
			"start_time":  time.Now().In(time.Local),
			"expire_time": time.Now().Add(24 * time.Hour).In(time.Local),
		})
		c.JSON(200, model.UpdateResponse{
			model.Response{200, "更新成功"},
		})
	} else {
		c.JSON(200, model.UpdateResponse{
			model.Response{404, "该链接不存在"},
		})
	}

}
func Delete(c *gin.Context) {
	cur, _ := getcuruser(c)
	var url []model.UrlInfo
	userid := cur.GetId()
	id := c.PostForm("id")
	if id == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入需要删除的链接ID"})
		return
	}
	dao.Getdb().Where("user_id=? and id=?", userid, id).Find(&url)
	if len(url) == 0 {
		c.JSON(200, model.Response{
			404, "表中没有该数据",
		})
	} else {
		dao.Getdb().Model(model.UrlInfo{}).Delete(&url)
		c.JSON(200, model.UpdateResponse{
			model.Response{200, "删除成功"},
		})
	}
}
func Pause(c *gin.Context) {
	cur, _ := getcuruser(c)
	id := c.PostForm("id")
	if id == "" {
		c.AbortWithStatusJSON(200, gin.H{"code": 403, "msg": "请输入需要冻结的链接ID"})
		return
	}
	user_id := cur.Id
	var url model.UrlInfo
	var purl model.PauseUrl
	url.Short = ""
	purl.Short = ""
	dao.Getdb().Model(&model.UrlInfo{}).Where("id=? and user_id=?", id, user_id).First(&url)
	dao.Getdb().Model(&model.PauseUrl{}).Where("url_id=? and user_id=?", id, user_id).First(&purl)
	if url.Short != "" && purl.Short == "" {
		var p model.PauseUrl
		p.UrlId, _ = strconv.Atoi(id)
		p.UserId = user_id
		p.Short = url.Short
		dao.Getdb().Model(&model.PauseUrl{}).Create(p)
		c.JSON(200, gin.H{"msg": "短链接暂停成功"})
	} else if url.Short != "" && purl.Short != "" {
		dao.Getdb().Model(&model.PauseUrl{}).Where("url_id=? and user_id=?", id, user_id).Delete(&purl)
		c.JSON(200, gin.H{"msg": "短链接重新启用"})
	} else {
		c.JSON(200, gin.H{"msg": "请输入正确的短链接编号"})
	}
}
