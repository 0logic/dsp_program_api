package main

import (
	"dsp_program_api/controllers"
	"dsp_program_api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	/**
	*1、if have diff media api then add middlewares
	*2、inside bidding
	*3、return campaign info
	 */

	// start
	r := gin.Default()
	// proxy cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	//corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Token"}
	corsConfig.ExposeHeaders = []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// api
	apiGroup := r.Group("/api")

	// global add middlewares

	adminGroup := apiGroup.Group("/admin")
	adminGroup.Use(middlewares.AddRequestLog("程序化后台", "请求日志"))
	//adminGroup.Use(middlewares.CheckAuth("程序化后台", "权限校验"))
	{
		//adminGroup.GET("/campaign", controllers.Campaign)
		//adminGroup.POST("/material", controllers.Material)
		//adminGroup.POST("/account", controllers.Account)
		adminGroup.Any("/guest", controllers.Guest)
		adminGroup.Any("/media", controllers.DspMedia)
		adminGroup.Any("/mediaPosition", controllers.DspMediaPosition)
		adminGroup.Any("/adPartner", controllers.DspAdPartner)
		adminGroup.Any("/channelPartner", controllers.DspChannelPartner)
	}

	//mojiApp request
	mojiGroup := apiGroup.Group("/moji")
	mojiGroup.Use(middlewares.AddRequestLog("墨迹天气"))
	{
		//get media request
		//mojiGroup.GET("/callback", controllers.GetCall)
	}

	// run
	r.Run(":8090")

}
