package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortlink/app/controller"
)

func Router() {
	r := gin.Default()
	/*
		添加前端需求的方法
		r.SetFuncMap(template.FuncMap{

		})

	*/
	//r.LoadHTMLGlob("htmlFilePath")htmlFilePath="templates/*"解析模板
	//r.Static("/statics/html/xxx", "./statics/html") //处理静态文件
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	userRoute := r.Group("/user")
	{
		userRoute.POST("/register", controller.Register)
		userRoute.POST("/login", controller.Login)
		userRoute.POST("/logout", controller.Logout)
		userRoute.POST("/info", controller.GetInfo)
		userRoute.POST("/record/get", controller.GetLoginInfo)
		userRoute.POST("/url/get", controller.GetUrl)

	}
	urlRoute := r.Group("/url")
	{
		urlRoute.POST("/create", controller.Create)
		urlRoute.POST("/query", controller.Query)
		urlRoute.PUT("/update", controller.Update)
		urlRoute.DELETE("/delete", controller.Delete)
		urlRoute.POST("/pause", controller.Pause)
		urlRoute.POST("/shorten", controller.Shorten)
	}
	go controller.Clean()
	r.Run(":9090")
}
