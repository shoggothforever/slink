package controller

import (
	"github.com/sirupsen/logrus"
	"shortlink/dao"
	"shortlink/model"
)

func SaveUser(u *model.User) {
	if err := dao.Db.Create(u).Error; err != nil {
		logrus.Error("插入数据失败", err)
	}

}
