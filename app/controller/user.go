package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shortlink/dao"
	"shortlink/model"
	"time"
)

/*
向数据库中保存用户，用户名和邮箱地址不能重复
*/
func SaveUser(u *model.User) error {
	var user []model.User
	dao.Getdb().Where("name=?", u.Name).Find(&user)
	if len(user) != 0 {
		logrus.Info("该用户名已存在")
		*u = user[0]
		return gorm.ErrRegistered
	}
	dao.Getdb().Where("email=?", u.Email).Find(&user)
	if len(user) != 0 {
		logrus.Info("该邮箱地址已存在")
		*u = user[0]
		return gorm.ErrRegistered
	}
	/*
		给用户密码进行MD5加密
	*/
	u.Pwd = messagedigest5(u.Pwd)
	dao.Lock.Lock()
	if err := dao.Getdb().Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		dao.Lock.Unlock()
		return gorm.ErrNotImplemented
	} else {
		dao.Lock.Unlock()
	}
	return nil
}
func SaveUrl(u *model.UrlInfo, userid int) error {
	var dupl []model.UrlInfo
	dao.Getdb().Where("user_id=? and origin=? and short=?", userid, u.Origin, u.Short).First(&dupl)
	if len(dupl) != 0 {
		logrus.Info("数据已存在")
		*u = dupl[0]
		return nil
	}
	var data []model.User
	dao.Getdb().Where("id", u.UserId).Find(&data)
	if len(data) == 0 {
		logrus.Info("UserId invalid,Please use this service after Login")
	}
	u.StartTime = time.Now().In(time.Local)
	u.ExpireTime = time.Now().In(time.Local).Add(time.Hour * 24) //暂时没有自定义短链接过期时间
	u.UserId = userid
	dao.Lock.Lock()
	if err := dao.Getdb().Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		dao.Lock.Unlock()
		return gorm.ErrNotImplemented
	} else {
		dao.Lock.Unlock()
	}
	return nil
}
func SaveLogin(l *model.LoginInfo, userid int) error {
	l.LoginAt = time.Now().In(time.Local)
	l.UserId = userid
	dao.Lock.Lock()
	if err := dao.Getdb().Omit("id").Create(&l).Error; err != nil {
		logrus.Error("插入数据失败", err)
		dao.Lock.Unlock()
		return gorm.ErrNotImplemented
	} else {
		dao.Lock.Unlock()
	}

	return nil
}
func CleanUrl() {
	st := time.Now().Unix()
	if exist := dao.Getdb().Migrator().HasTable("url_infos"); exist == true {
		for {
			ed := time.Now().Unix() - st
			if ed >= 10 {
				dao.Lock.Lock()
				dao.Getdb().Exec("delete from url_infos where datediff(NOW(),url_infos.start_time)>=1")
				st = ed
				dao.Lock.Unlock()
			}
		}
	}
}
func CleanJwt() {
	st := time.Now().Unix()
	if exist := dao.Getdb().Migrator().HasTable("cookies"); exist == true {
		for {
			//fmt.Println(st)
			ed := time.Now().Unix() - st
			if ed >= 10 {
				dao.Lock.Lock()
				dao.Getdb().Exec("delete from cookies where datediff(NOW(),created_at)>=1")
				st = ed
				dao.Lock.Unlock()
			}
		}
	}
}
func CleanLogin() {
	st := time.Now().Unix()
	if exist := dao.Getdb().Migrator().HasTable("login_infos"); exist == true {
		for {
			//fmt.Println(st)
			ed := time.Now().Unix() - st
			if ed >= 10 {
				dao.Lock.Lock()
				dao.Getdb().Exec("delete from login_infos where datediff(NOW(),login_at)>=30")
				st = ed
				dao.Lock.Unlock()
			}
		}
	}
}
func getcuruser(c *gin.Context) (model.User, bool) {
	tmpuser, ok := c.Get("user")
	if ok != false {
		return tmpuser.(model.User), ok
	}
	return model.User{}, ok
}
