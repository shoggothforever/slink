package main

import (
	"shortlink/dao"
	"shortlink/router"
)

func main() {
	dao.Init()
	//for i := 0; i <= 1; i++ {
	//	var u model.User
	//	u.Name = "dsm"
	//	u.Pwd = "123"
	//	u.Email = "133"
	//	var ur model.UrlInfo
	//	ur.Origin = "www.baidu.com"
	//	ur.Comment = "dut"
	//	ur.ExpireTime = time.Now().Add(24 * time.Hour)
	//	u.Url = ur
	//	dao.Db.Create(&u)
	//}
	router.Router()

}
