package main

import (
	"shortlink/dao"
	"shortlink/router"
)

func main() {
	dao.Init()
	router.Router()
}
