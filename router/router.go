package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shortlink/api/controller"
	"shortlink/model"
	"time"
)

func Router() {
	htmlFilePath := "templates/*.html"
	f, _ := os.Create("sl.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	/*
		添加前端需求的方法
		r.SetFuncMap(template.FuncMap{

		})
	*/
	r.LoadHTMLGlob(htmlFilePath)                         //htmlFilePath="templates/*"解析模板
	r.Static("static/html/static", "./templates/static") //处理静态文件
	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.Use(controller.RedirectShort())
	r.GET("/", func(c *gin.Context) {
		c.Set("userid", model.NOTLOGIN)
		c.HTML(200, "index.html", nil)
		//time.Sleep(5 * time.Second)
		//c.String(http.StatusOK, "Welcome Gin Server")
	})
	r.GET("/exit", func(c *gin.Context) {
		srv.Shutdown(context.Background())
	})
	r.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}

		c.JSON(200, fmt.Sprintf("Cookie value: %s \n", cookie))
	})
	userRoute := r.Group("/user")
	{
		userRoute.POST("/register", controller.Register)
		userRoute.POST("/login", controller.Login)
		userRoute.POST("/logout", controller.AuthJwt(), controller.Logout)
		userRoute.GET("/info", controller.AuthJwt(), controller.GetInfo)
		userRoute.GET("/record/get", controller.AuthJwt(), controller.GetLoginInfo)
		userRoute.GET("/url/get", controller.AuthJwt(), controller.GetUrl)
	}
	urlRoute := r.Group("/url", controller.AuthJwt())
	{
		urlRoute.POST("/create", controller.Create)
		urlRoute.POST("/query", controller.Query)
		urlRoute.PUT("/update", controller.Update)
		urlRoute.DELETE("/delete", controller.Delete)
		urlRoute.POST("/pause", controller.Pause)
	}

	//平滑地关机
	go controller.CleanUrl()
	go controller.CleanJwt()
	go controller.CleanLogin()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
