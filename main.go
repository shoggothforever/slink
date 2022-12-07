package main

import (
	"fmt"
	"shortlink/app/controller"
	"shortlink/dao"
	"shortlink/model"
)

func main() {
	dao.Init()
	//router.Router()
	for {
		var u model.User
		fmt.Println("输入姓名，邮箱和密码")
		fmt.Scanf("%s %s %s", &u.Name, &u.Email, &u.Pwd)
		//var url model.UrlInfo
		//fmt.Println("输入长链接和评论")
		//fmt.Scanf("%s %s", &url.Origin, &url.Comment)
		//u.Url = url
		//fmt.Println("输入完毕")
		controller.SaveUser(&u)
	}
}
