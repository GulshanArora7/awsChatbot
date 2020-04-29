package main

import (
	"github.com/GulshanArora7/awsChatbot/config"
	"github.com/GulshanArora7/awsChatbot/controller"
	_ "github.com/GulshanArora7/awsChatbot/gateway/slackclient"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	g := e.Group("/awschatbot")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/v1/health", controller.CheckHealth)
	g.POST("/v1/slackevent", controller.SlackMessageEvents)
	g.POST("/v1/response_actions", controller.SlackReplyEvents)
	e.Logger.Fatal(e.Start(":" + config.Port))
}
