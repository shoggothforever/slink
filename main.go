package main

import (
	"shortlink/dao"
	"shortlink/router"
)

func init() {
	dao.Init()
}
func main() {
	router.Router()
}
