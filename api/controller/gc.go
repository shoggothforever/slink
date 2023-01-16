package controller

import (
	"github.com/gin-gonic/gin"
	"shortlink/dao"
	"shortlink/model"
	"time"
)

/*
向数据库中保存用户信息，用户名和邮箱地址不能重复


清除过期URL信息
*/
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

/*
清除过期JWT信息
*/
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

/*
清除久远登录信息
*/
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

/*
从上下文中获取当前登录用户信息
*/
func getcuruser(c *gin.Context) (model.User, bool) {
	tmpuser, ok := c.Get("user")
	if ok != false {
		return tmpuser.(model.User), ok
	}
	return model.User{}, ok
}
