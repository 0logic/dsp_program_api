package controllers

import (
	"dsp_program_api/models/Table"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 合作媒体的获取
func DspAdPartnerGet(ctx *gin.Context) {
	ad_partner_name, _ := ctx.GetQuery("ad_partner_name")
	ad_partner_style, _ := ctx.GetQuery("ad_partner_style")
	ad_partner_type, _ := ctx.GetQuery("ad_partner_type")
	status := ctx.GetInt("status")

	pages, _ := ctx.GetQuery("page")
	page, _ := strconv.Atoi(pages)
	page_sizes, _ := ctx.GetQuery("page_size")
	page_size, _ := strconv.Atoi(page_sizes)

	countList := 0
	listQuery := Table.DspAdPartnerDB

	if ad_partner_name != "" {
		listQuery = listQuery.Where("ad_partner_name LIKE?", fmt.Sprintf("%%%s%%", ad_partner_name))
	}

	if ad_partner_style != "" {
		listQuery = listQuery.Where("ad_partner_style=?", ad_partner_style)
	}

	if ad_partner_type != "" {
		listQuery = listQuery.Where("ad_partner_type=?", ad_partner_type)
	}

	if status > 0 {
		listQuery = listQuery.Where("status=?", status)
	}

	listQuery.Model(&Table.DspAdPartner{}).Count(&countList)

	offsets := (page - 1) * page_size

	// 获取数据
	var listRes []Table.DspAdPartner
	listQuery.Offset(offsets).Limit(page_size).Find(&listRes)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    990,
		"message": "成功",
		"count":   countList,
		"data":    listRes,
	})

}

// 合作媒体的创建
func DspAdPartnerPost(ctx *gin.Context) {

	ad_partner_name, _ := ctx.GetPostForm("ad_partner_name")
	ad_partner_short_name, _ := ctx.GetPostForm("ad_partner_short_name")
	ad_partner_type, _ := ctx.GetPostForm("ad_partner_type")
	ad_partner_style, _ := ctx.GetPostForm("ad_partner_style")
	defaultSellerId := 0
	seller_id, err1 := strconv.Atoi(ctx.PostForm("seller_id"))
	if err1 != nil {
		seller_id = defaultSellerId
	}

	defaultStatus := 0
	status, err2 := strconv.Atoi(ctx.PostForm("status"))
	if err2 != nil {
		status = defaultStatus
	}

	// 获取创建人
	var defaultCreatedId int64 = 1
	created_id, err := strconv.ParseInt(ctx.PostForm("created_id"), 10, 64)
	if err != nil {
		created_id = defaultCreatedId
	}

	// 判断是否已存在
	ishave := 0
	Table.DspAdPartnerDB.Model(&Table.DspAdPartner{}).Where("ad_partner_name=?", ad_partner_name).Where("ad_partner_type=?", ad_partner_type).Where("ad_partner_style=?", ad_partner_style).Count(&ishave)
	if ishave > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在该广告主！",
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
	Table.DspAdPartnerDB.Create(&Table.DspAdPartner{
		AdPartnerName:      ad_partner_name,
		AdPartnerShortName: ad_partner_short_name,
		AdPartnerType:      ad_partner_type,
		AdPartnerStyle:     ad_partner_style,
		SellerId:           uint64(seller_id),
		Status:             uint(status),
		CreatedId:          uint64(created_id),
	})

}

// 合作媒体的删除
func DspAdPartnerDelete(ctx *gin.Context) {

}

// 合作媒体的更新
func DspAdPartnerPut(ctx *gin.Context) {
	var defaultPartnerId int64 = 0
	ad_partner_id, err := strconv.ParseInt(ctx.PostForm("ad_partner_id"), 10, 64)
	if err != nil {
		ad_partner_id = defaultPartnerId
	}

	ad_partner_name, _ := ctx.GetPostForm("ad_partner_name")
	ad_partner_short_name, _ := ctx.GetPostForm("ad_partner_short_name")
	ad_partner_style, _ := ctx.GetPostForm("ad_partner_style")
	ad_partner_type, _ := ctx.GetPostForm("ad_partner_type")

	//获取对应的媒体类型及广告位
	var firstModel Table.DspAdPartner
	if err := Table.DspAdPartnerDB.Where("ad_partner_id=?", ad_partner_id).First(&firstModel).Error; err != nil {
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

	defaultSellerId := firstModel.SellerId
	seller_id, err1 := strconv.Atoi(ctx.PostForm("seller_id"))
	if err1 != nil {
		seller_id = int(defaultSellerId)
	}

	if ad_partner_name != "" {
		firstModel.AdPartnerName = ad_partner_name
	}
	if ad_partner_short_name != "" {
		firstModel.AdPartnerShortName = ad_partner_short_name
	}
	if ad_partner_style != "" {
		firstModel.AdPartnerStyle = ad_partner_style
	}
	if ad_partner_type != "" {
		firstModel.AdPartnerType = ad_partner_type
	}

	// 判断是否重复
	isAnother := 0
	Table.DspAdPartnerDB.Model(&Table.DspAdPartner{}).Where("ad_partner_id!=?", firstModel.AdPartnerId).Where("ad_partner_name=?", firstModel.AdPartnerName).Where("ad_partner_style=?", firstModel.AdPartnerStyle).Where("ad_partner_type=?", firstModel.AdPartnerType).Count(&isAnother)
	if isAnother > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在广告主，请勿重复提交！",
		})
		return
	}

	// 更新数据
	if firstModel.Status != uint(status) {
		firstModel.Status = uint(status)
	}
	if firstModel.SellerId != uint64(seller_id) {
		firstModel.SellerId = uint64(seller_id)
	}

	if err := Table.DspAdPartnerDB.Model(&Table.DspAdPartner{}).Where("ad_partner_id=?", firstModel.AdPartnerId).Update(firstModel).Error; err != nil {
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

func DspAdPartner(ctx *gin.Context) {
	/**
	1、广告主 的相关操作，按照rest api方式
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		DspAdPartnerGet(ctx)
		break
	case "POST":
		DspAdPartnerPost(ctx)
		break
	case "PUT":
		DspAdPartnerPut(ctx)
		break
	case "DELETE":
		DspAdPartnerDelete(ctx)
		break
	default:
		break
	}
}
