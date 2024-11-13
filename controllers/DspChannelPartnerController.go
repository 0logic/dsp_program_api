package controllers

import (
	"dsp_program_api/models/Table"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 合作媒体的获取
func DspChannelPartnerGet(ctx *gin.Context) {
	channel_partner_name, _ := ctx.GetQuery("channel_partner_name")
	channel_partner_type, _ := ctx.GetQuery("channel_partner_type")

	status := ctx.GetInt("status")
	channel_partner_open := ctx.GetInt("channel_partner_open")

	pages, _ := ctx.GetQuery("page")
	page, _ := strconv.Atoi(pages)
	page_sizes, _ := ctx.GetQuery("page_size")
	page_size, _ := strconv.Atoi(page_sizes)

	countList := 0
	listQuery := Table.DspChannelPartnerDB

	if channel_partner_name != "" {
		listQuery = listQuery.Where("channel_partner_name LIKE?", fmt.Sprintf("%%%s%%", channel_partner_name))
	}

	if channel_partner_type != "" {
		listQuery = listQuery.Where("channel_partner_type=?", channel_partner_type)
	}

	if channel_partner_open > 0 {
		listQuery = listQuery.Where("channel_partner_open=?", channel_partner_open)
	}

	if status > 0 {
		listQuery = listQuery.Where("status=?", status)
	}

	listQuery.Model(&Table.DspChannelPartner{}).Count(&countList)

	offsets := (page - 1) * page_size

	// 获取数据
	var listRes []Table.DspChannelPartner
	listQuery.Offset(offsets).Limit(page_size).Find(&listRes)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    990,
		"message": "成功",
		"count":   countList,
		"data":    listRes,
	})

}

// 合作媒体的创建
func DspChannelPartnerPost(ctx *gin.Context) {

	channel_partner_name, _ := ctx.GetPostForm("channel_partner_name")
	channel_partner_type, _ := ctx.GetPostForm("channel_partner_type")

	defaultSpecialistId := 0
	media_specialist_id, err1 := strconv.Atoi(ctx.PostForm("media_specialist_id"))
	if err1 != nil {
		media_specialist_id = defaultSpecialistId
	}

	defaultStatus := 0
	status, err2 := strconv.Atoi(ctx.PostForm("status"))
	if err2 != nil {
		status = defaultStatus
	}

	defaultPartnerOpen := 0
	channel_partner_open, err3 := strconv.Atoi(ctx.PostForm("channel_partner_open"))
	if err3 != nil {
		channel_partner_open = defaultPartnerOpen
	}

	// 获取创建人
	var defaultCreatedId int64 = 1
	created_id, err := strconv.ParseInt(ctx.PostForm("created_id"), 10, 64)
	if err != nil {
		created_id = defaultCreatedId
	}

	// 判断是否已存在
	ishave := 0
	Table.DspChannelPartnerDB.Model(&Table.DspChannelPartner{}).Where("channel_partner_name=?", channel_partner_name).Where("channel_partner_type=?", channel_partner_type).Count(&ishave)
	if ishave > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在该渠道主！",
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
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    999,
				"message": "失败",
			})
		}
	}()
	Table.DspChannelPartnerDB.Create(&Table.DspChannelPartner{
		ChannelPartnerName: channel_partner_name,
		ChannelPartnerType: channel_partner_type,
		ChannelPartnerOpen: uint(channel_partner_open),
		MediaSpecialistId:  uint64(media_specialist_id),
		Status:             uint(status),
		CreatedId:          uint64(created_id),
	})

}

// 合作媒体的删除
func DspChannelPartnerDelete(ctx *gin.Context) {

}

// 合作媒体的更新
func DspChannelPartnerPut(ctx *gin.Context) {
	var defaultPartnerId int64 = 0
	channel_partner_id, err := strconv.ParseInt(ctx.PostForm("channel_partner_id"), 10, 64)
	if err != nil {
		channel_partner_id = defaultPartnerId
	}

	channel_partner_name, _ := ctx.GetPostForm("channel_partner_name")
	channel_partner_type, _ := ctx.GetPostForm("channel_partner_type")

	//获取对应的媒体类型及广告位
	var firstModel Table.DspChannelPartner
	if err := Table.DspChannelPartnerDB.Where("channel_partner_id=?", channel_partner_id).First(&firstModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "无效ID",
		})
		return
	}

	defaultStatus := firstModel.Status
	status, err1 := strconv.Atoi(ctx.PostForm("status"))
	if err1 != nil {
		status = int(defaultStatus)
	}

	defaultMediaId := firstModel.MediaSpecialistId
	media_specialist_id, err2 := strconv.Atoi(ctx.PostForm("seller_id"))
	if err2 != nil {
		media_specialist_id = int(defaultMediaId)
	}

	defaultOpen := firstModel.ChannelPartnerOpen
	channel_partner_open, err3 := strconv.Atoi(ctx.PostForm("channel_partner_open"))
	if err3 != nil {
		channel_partner_open = int(defaultOpen)
	}

	if channel_partner_name != "" {
		firstModel.ChannelPartnerName = channel_partner_name
	}
	if channel_partner_type != "" {
		firstModel.ChannelPartnerType = channel_partner_type
	}

	// 判断是否重复
	isAnother := 0
	Table.DspChannelPartnerDB.Model(&Table.DspChannelPartner{}).Where("channel_partner_id!=?", firstModel.ChannelPartnerId).Where("channel_partner_name=?", firstModel.ChannelPartnerName).Where("channel_partner_type=?", firstModel.ChannelPartnerType).Count(&isAnother)
	if isAnother > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在渠道主，请勿重复提交！",
		})
		return
	}

	// 更新数据
	if firstModel.Status != uint(status) {
		firstModel.Status = uint(status)
	}
	if firstModel.MediaSpecialistId != uint64(media_specialist_id) {
		firstModel.MediaSpecialistId = uint64(media_specialist_id)
	}
	if firstModel.ChannelPartnerOpen != uint(channel_partner_open) {
		firstModel.ChannelPartnerOpen = uint(channel_partner_open)
	}

	if err := Table.DspChannelPartnerDB.Model(&Table.DspChannelPartner{}).Where("channel_partner_id=?", firstModel.ChannelPartnerId).Update(firstModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    990,
		"message": "成功",
	})

}

func DspChannelPartner(ctx *gin.Context) {
	/**
	1、渠道客户 的相关操作，按照rest api方式
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		DspChannelPartnerGet(ctx)
		break
	case "POST":
		DspChannelPartnerPost(ctx)
		break
	case "PUT":
		DspChannelPartnerPut(ctx)
		break
	case "DELETE":
		DspChannelPartnerDelete(ctx)
		break
	default:
		break
	}
}
