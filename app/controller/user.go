package controller

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shortlink/dao"
	"shortlink/model"
	"time"
)

func SaveUser(u *model.User) error {

	if err := dao.Db.Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		return gorm.ErrNotImplemented
	}
	return nil
}
func SaveUrl(u *model.UrlInfo) error {
	var data []model.User
	dao.Db.Where("id", u.UserId).Find(&data)
	if len(data) == 0 {
		logrus.Info("UserId invalid,Please use this service after Login")
	}
	u.StartTime = time.Now()
	u.ExpireTime = time.Now().Add(time.Hour * 24)
	u.UserId = model.CurrentUser.GetId()
	u.Short = ""
	if err := dao.Db.Omit("id").Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
		return gorm.ErrNotImplemented
	}
	return nil
}
func SaveLogin(l *model.LoginInfo) error {
	l.LoginAt = time.Now()
	l.UserId = model.CurrentUser.Id
	if err := dao.Db.Omit("id").Create(&l).Error; err != nil {
		logrus.Error("插入数据失败", err)
		return gorm.ErrNotImplemented
	}
	return nil
}
