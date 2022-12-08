package controller

import "github.com/gin-gonic/gin"

//使用中间件实现重定向
func RedirectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//使用中间件实现赋值
func ADDInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
