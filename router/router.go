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
	"shortlink/api/middleware"
	"shortlink/api/views"
	"shortlink/model"
	"time"
)

func Router() {
	f, _ := os.Create("sl.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	r.Use(middleware.RedirectShort())
	r.GET("/", func(c *gin.Context) {
		c.Set("userid", model.NOTLOGIN)
		time.Sleep(2 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
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
		userRoute.POST("/register", views.Register)
		userRoute.POST("/login", views.Login)
		userRoute.POST("/logout", middleware.AuthJwt(), views.Logout)
		userRoute.GET("/info", middleware.AuthJwt(), views.GetInfo)
		userRoute.GET("/record/get", middleware.AuthJwt(), views.GetLoginInfo)
		userRoute.GET("/url/get", middleware.AuthJwt(), views.GetUrl)
	}
	urlRoute := r.Group("/url", middleware.AuthJwt())
	{
		urlRoute.POST("/create", views.Create)
		urlRoute.POST("/query", views.Query)
		urlRoute.PUT("/update", views.Update)
		urlRoute.DELETE("/delete", views.Delete)
		urlRoute.POST("/pause", views.Pause)
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
