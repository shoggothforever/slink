package dao

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shortlink/api/utils"
	"shortlink/model"
	"time"
)

func SaveUser(u *model.User) error {
	var user []model.User
	Getdb().Where("name=?", u.Name).Find(&user)
	if len(user) != 0 {
		logrus.Info("该用户名已存在")
		*u = user[0]
		return gorm.ErrRegistered
	}
	Getdb().Where("email=?", u.Email).Find(&user)
	if len(user) != 0 {
		logrus.Info("该邮箱地址已存在")
		*u = user[0]
		return gorm.ErrRegistered
	}
	/*
		给用户密码进行MD5加密
	*/
	u.Pwd = utils.Messagedigest5(u.Pwd)
	Lock.Lock()
	if err := Getdb().Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		defer Lock.Unlock()
		return gorm.ErrNotImplemented
	} else {
		defer Lock.Unlock()
	}
	return nil
}

/*
向数据库中保存链接转换信息，数据不能重复
*/
func SaveUrl(u *model.UrlInfo, userid int) error {
	var dupl []model.UrlInfo
	Getdb().Where("user_id=? and origin=? and short=?", userid, u.Origin, u.Short).First(&dupl)
	if len(dupl) != 0 {
		logrus.Info("数据已存在")
		*u = dupl[0]
		return nil
	}
	var data []model.User
	Getdb().Where("id", u.UserId).Find(&data)
	if len(data) == 0 {
		logrus.Info("UserId invalid,Please use this service after Login")
	}
	u.StartTime = time.Now().In(time.Local)
	u.ExpireTime = time.Now().In(time.Local).Add(time.Hour * 24) //暂时没有自定义短链接过期时间
	u.UserId = userid
	Lock.Lock()
	if err := Getdb().Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		defer Lock.Unlock()
		return gorm.ErrNotImplemented
	} else {
		defer Lock.Unlock()
	}
	return nil
}

/*
向数据库中保存用户登录信息
*/
func SaveLogin(l *model.LoginInfo, userid int) error {
	l.LoginAt = time.Now().In(time.Local)
	l.UserId = userid
	Lock.Lock()
	if err := Getdb().Omit("id").Create(&l).Error; err != nil {
		logrus.Error("插入数据失败", err)
		defer Lock.Unlock()
		return gorm.ErrNotImplemented
	} else {
		defer Lock.Unlock()
	}

	return nil
}

func SaveJwt(id int, name string) {
	utils.ExpireTime = 4 * time.Hour
	model.AuthJwt, _ = utils.GenerateJwt(id, name)
	model.AuthClaims, _ = utils.ParseToken(model.AuthJwt)
	var cookie model.Cookie
	cookie.UserId = id
	cookie.Jwt = model.AuthJwt
	cookie.CreatedAt = time.Now().In(time.Local)
	var del model.Cookie
	Getdb().Model(&cookie).Where("user_id=?", id).Delete(&del)
	Getdb().Model(&cookie).Create(&cookie)
}
