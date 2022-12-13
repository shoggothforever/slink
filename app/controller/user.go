package controller

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
	"shortlink/dao"
	"shortlink/model"
	"time"
)

/*
向数据库中保存用户，用户名和邮箱地址不能重复
*/
func SaveUser(u *model.User) error {
	var user []model.User
	dao.Db.Where("name=?", u.Name).Find(&user)
	if len(user) != 0 {
		logrus.Info("该用户名已存在")
		*u = user[0]
		return gorm.ErrRegistered
	}
	dao.Db.Where("email=?", u.Email).Find(&user)
	if len(user) != 0 {
		logrus.Info("该邮箱地址已存在")
		*u = user[0]
		return gorm.ErrRegistered
	}
	if err := dao.Db.Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		return gorm.ErrNotImplemented
	}

	return nil
}
func SaveUrl(u *model.UrlInfo) error {
	var dupl []model.UrlInfo
	dao.Db.Where("user_id=? and origin=? and short=?", model.CurrentUser.GetId(), u.Origin, u.Short).First(&dupl)
	if len(dupl) != 0 {
		logrus.Info("数据已存在")
		*u = dupl[0]
		return nil
	}
	var data []model.User
	dao.Db.Where("id", u.UserId).Find(&data)
	if len(data) == 0 {
		logrus.Info("UserId invalid,Please use this service after Login")
	}
	u.StartTime = time.Now().In(time.Local)
	u.ExpireTime = time.Now().In(time.Local).Add(time.Hour * 24)
	u.UserId = model.CurrentUser.GetId()
	if err := dao.Db.Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		return gorm.ErrNotImplemented
	}
	return nil
}
func SaveLogin(l *model.LoginInfo) error {
	l.LoginAt = time.Now().In(time.Local)
	l.UserId = model.CurrentUser.Id
	if err := dao.Db.Omit("id").Create(&l).Error; err != nil {
		logrus.Error("插入数据失败", err)
		return gorm.ErrNotImplemented
	}
	return nil
}
func CleanUrl() {
	st := time.Now().Unix()
	if exist := dao.Db.Migrator().HasTable("url_infos"); exist == true {
		for {
			ed := time.Now().Unix() - st
			if ed >= 10 {
				dao.Lock.Lock()
				dao.Db.Exec("delete from url_infos where datediff(NOW(),url_infos.start_time)>=1")
				st = ed
				dao.Lock.Unlock()
			}
		}
	}
}
func CleanJwt() {
	st := time.Now().Unix()
	if exist := dao.Db.Migrator().HasTable("cookies"); exist == true {
		for {
			//fmt.Println(st)
			ed := time.Now().Unix() - st
			if ed >= 10 {
				dao.Lock.Lock()
				dao.Db.Exec("delete from cookies where datediff(NOW(),created_at)>=1")
				st = ed
				dao.Lock.Unlock()
			}
		}
	}
}
func CleanLogin() {
	st := time.Now().Unix()
	if exist := dao.Db.Migrator().HasTable("login_infos"); exist == true {
		for {
			//fmt.Println(st)
			ed := time.Now().Unix() - st
			if ed >= 10 {
				dao.Lock.Lock()
				dao.Db.Exec("delete from login_infos where datediff(NOW(),login_at)>=30")
				st = ed
				dao.Lock.Unlock()
			}
		}
	}
}
func GenShort(short string) string {
	if short != "" {
		return short
	}
	ts := time.Now().UnixNano()
	rand.Seed(ts)
	return Encode(int(ts))
}
