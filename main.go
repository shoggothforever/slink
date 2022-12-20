package main

import (
	"shortlink/dao"
	"shortlink/router"
)

//	@title			shortLink API
//	@version		1.0
//	@description	This is a sample shortLink server.

//	@contact.name	Kalun
//	@contact.url	124.220.190.203
//	@contact.email	1337231450@qq.com

//	@host		localhost:3000
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth
func init() {
	dao.Init()
}
func main() {
	router.Router()
}
