package controllers

import (
	"github.com/gin-gonic/gin"
)

// 广告的获取
func CampaignGet(ctx *gin.Context) {

}

// 广告的创建
func CampaignPost(ctx *gin.Context) {

}

// 广告的删除
func CampaignDelete(ctx *gin.Context) {

}

// 广告的更新
func CampaignPut(ctx *gin.Context) {

}

func Campaign(ctx *gin.Context) {
	/**
	1、campaign 的相关操作，按照rest api方式
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		CampaignGet(ctx)
		break
	case "POST":
		CampaignPost(ctx)
		break
	case "PUT":
		CampaignPut(ctx)
		break
	case "DELETE":
		CampaignDelete(ctx)
		break
	default:
		break
	}
}
