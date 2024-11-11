package controllers

import (
	"dsp_program_api/models/Table"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 合作客户的获取
func GuestGet(ctx *gin.Context) {

}

// 合作客户的创建
func GuestPost(ctx *gin.Context) {
	guest_name, _ := ctx.GetPostForm("name")
	guest_email, _ := ctx.GetPostForm("email")
	guest_phone, _ := ctx.GetPostForm("phone")
	guest_company, _ := ctx.GetPostForm("company")
	country, _ := ctx.GetPostForm("country")
	reason, _ := ctx.GetPostForm("reason") // 处理原因长度过长
	if len(reason) > 250 {
		reason = reason[0:248]
	}

	guest_ip := ctx.ClientIP()

	// 同一个用户当日不能超过2次提交
	var count int

	now := time.Now()
	year, month, day := now.Date()
	dawn := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	dawn.Format("2006-01-02 15:04:05")

	Table.AdvertiserGuestDB.Model(&Table.AdvertiserGuest{}).Where("guest_name=?", guest_name).Where("created_at>?", dawn).Count(&count)
	if count > 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    990,
			"message": "已提交，请勿重复提交",
		})
		return
	}
	defer func() {
		e := recover()
		if e == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    990,
				"message": "成功",
			})
		}
	}()
	Table.AdvertiserGuestDB.Create(&Table.AdvertiserGuest{
		GuestName:    guest_name,
		GuestEmail:   guest_email,
		GuestPhone:   guest_phone,
		GuestCompany: guest_company,
		Country:      country,
		Reason:       reason,
		GuestIp:      guest_ip,
	})
}

// 合作客户的删除
func GuestDelete(ctx *gin.Context) {

}

// 合作客户的更新
func GuestPut(ctx *gin.Context) {

}

func Guest(ctx *gin.Context) {
	/**
	1、guest 的相关操作，按照rest api方式
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		GuestGet(ctx)
		break
	case "POST":
		GuestPost(ctx)
		break
	case "PUT":
		GuestPut(ctx)
		break
	case "DELETE":
		GuestDelete(ctx)
		break
	default:
		break
	}
}
