package controllers

import (
	"dsp_program_api/models/Table"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 合作媒体的获取
func DspMediaGet(ctx *gin.Context) {
	media_name, _ := ctx.GetQuery("media_name")
	media_campaign, _ := ctx.GetQuery("media_campaign")

	pages, _ := ctx.GetQuery("page")
	page, _ := strconv.Atoi(pages)
	page_sizes, _ := ctx.GetQuery("page_size")
	page_size, _ := strconv.Atoi(page_sizes)

	countList := 0
	listQuery := Table.DspMediaDB

	if media_name != "" {
		listQuery = listQuery.Where("media_name=?", media_name)
	}

	if media_campaign != "" {
		listQuery = listQuery.Where("media_campaign=?", media_campaign)
	}

	listQuery.Model(&Table.DspMedia{}).Count(&countList)

	offsets := (page - 1) * page_size

	// 获取数据
	var listRes []Table.DspMedia
	listQuery.Offset(offsets).Limit(page_size).Find(&listRes)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    990,
		"message": "成功",
		"count":   countList,
		"data":    listRes,
	})

}

// 合作媒体的创建
func DspMediaPost(ctx *gin.Context) {

	media_name, _ := ctx.GetPostForm("media_name")
	media_industry, _ := ctx.GetPostForm("media_industry")
	media_type, _ := ctx.GetPostForm("media_type")
	media_campaign, _ := ctx.GetPostForm("media_campaign")
	defaultStatus := 0
	is_support, err1 := strconv.Atoi(ctx.PostForm("is_support"))
	if err1 != nil {
		is_support = defaultStatus
	}

	// 获取创建人
	var defaultCreatedId int64 = 1
	created_id, err := strconv.ParseInt(ctx.PostForm("created_id"), 10, 64)
	if err != nil {
		created_id = defaultCreatedId
	}

	// 判断媒体类型及广告位是已提交
	ishave := 0
	Table.DspMediaDB.Model(&Table.DspMedia{}).Where("media_name=?", media_name).Where("media_campaign=?", media_campaign).Where("media_type=?", media_type).Count(&ishave)
	if ishave > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在该类型的媒体，请勿重复提交！",
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
	Table.DspMediaDB.Create(&Table.DspMedia{
		MediaName:     media_name,
		MediaIndustry: media_industry,
		MediaType:     media_type,
		MediaCampaign: media_campaign,
		IsSupport:     uint(is_support),
		CreatedId:     uint64(created_id),
	})

}

// 合作媒体的删除
func DspMediaDelete(ctx *gin.Context) {

}

// 合作媒体的更新
func DspMediaPut(ctx *gin.Context) {
	var defaultMediaId int64 = 0
	media_id, err := strconv.ParseInt(ctx.PostForm("media_id"), 10, 64)
	if err != nil {
		media_id = defaultMediaId
	}

	media_name, _ := ctx.GetPostForm("media_name")
	media_type, _ := ctx.GetPostForm("media_type")
	media_campaign, _ := ctx.GetPostForm("media_campaign")
	media_industry, _ := ctx.GetPostForm("media_industry")
	defaultStatus := 0
	is_support, err1 := strconv.Atoi(ctx.PostForm("is_support"))
	if err1 != nil {
		is_support = defaultStatus
	}

	//获取对应的媒体类型及广告位
	var firstModel Table.DspMedia
	if err := Table.DspMediaDB.Where("media_id=?", media_id).First(&firstModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "无效ID",
		})
		return
	}

	if media_name != "" {
		firstModel.MediaName = media_name
	}
	if media_campaign != "" {
		firstModel.MediaCampaign = media_campaign
	}
	if media_type != "" {
		firstModel.MediaType = media_type
	}

	// 判断是否重复
	isAnother := 0
	Table.DspMediaDB.Model(&Table.DspMedia{}).Where("media_id!=?", firstModel.MediaId).Where("media_name=?", firstModel.MediaName).Where("media_campaign=?", firstModel.MediaCampaign).Where("media_type=?", firstModel.MediaType).Count(&isAnother)
	if isAnother > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在该类型的媒体，请勿重复提交！",
		})
		return
	}

	// 更新数据
	firstModel.MediaIndustry = media_industry
	firstModel.IsSupport = uint(is_support)

	if err := Table.DspMediaDB.Model(&Table.DspMedia{}).Where("media_id=?", firstModel.MediaId).Update(firstModel).Error; err != nil {
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

func DspMedia(ctx *gin.Context) {
	/**
	1、Media 的相关操作，按照rest api方式
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		DspMediaGet(ctx)
		break
	case "POST":
		DspMediaPost(ctx)
		break
	case "PUT":
		DspMediaPut(ctx)
		break
	case "DELETE":
		DspMediaDelete(ctx)
		break
	default:
		break
	}
}
