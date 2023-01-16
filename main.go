package main

import (
	"shortlink/model"
	"shortlink/router"
)

func init() {
	model.Init()
}
func main() {
	router.Router()
}
