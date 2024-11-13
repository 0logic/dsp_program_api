package controllers

import (
	"dsp_program_api/models/Table"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 合作媒体广告位的获取
func DspMediaPositionGet(ctx *gin.Context) {
	position_name, _ := ctx.GetQuery("position_name")
	position_type, _ := ctx.GetQuery("position_type")

	pages, _ := ctx.GetQuery("page")
	page, _ := strconv.Atoi(pages)
	page_sizes, _ := ctx.GetQuery("page_size")
	page_size, _ := strconv.Atoi(page_sizes)

	countList := 0
	listQuery := Table.DspMediaPositionDB

	if position_name != "" {
		listQuery = listQuery.Where("position_name LIKE?", fmt.Sprintf("%%%s%%", position_name))
	}

	if position_type != "" {
		listQuery = listQuery.Where("position_type=?", position_type)
	}

	listQuery.Model(&Table.DspMediaPosition{}).Count(&countList)

	offsets := (page - 1) * page_size

	// 获取数据
	var listRes []Table.DspMediaPosition
	listQuery.Offset(offsets).Limit(page_size).Find(&listRes)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    990,
		"message": "成功",
		"count":   countList,
		"data":    listRes,
	})

}

// 合作媒体广告位的创建
func DspMediaPositionPost(ctx *gin.Context) {

	// 获取媒体ID
	var defaultMediaId int64 = 0
	media_id, err := strconv.ParseInt(ctx.PostForm("media_id"), 10, 64)
	if err != nil {
		media_id = defaultMediaId
	}
	position_name, _ := ctx.GetPostForm("position_name")
	position_type, _ := ctx.GetPostForm("position_type")
	position_size, _ := ctx.GetPostForm("position_size")
	connect_type, _ := ctx.GetPostForm("connect_type")
	// 是否精准匹配
	defaultAccurate := 0
	is_accurate, err1 := strconv.Atoi(ctx.PostForm("is_accurate"))
	if err1 != nil {
		is_accurate = defaultAccurate
	}
	//素材类型
	defaultType := 0
	position_material_type, err2 := strconv.Atoi(ctx.PostForm("position_material_type"))
	if err2 != nil {
		position_material_type = defaultType
	}
	//循环播放类型
	defaultRing := 0
	is_ring, err3 := strconv.Atoi(ctx.PostForm("is_ring"))
	if err3 != nil {
		is_ring = defaultRing
	}
	//是否自动播放
	defaultAuto := 1
	is_auto, err4 := strconv.Atoi(ctx.PostForm("is_auto"))
	if err4 != nil {
		is_auto = defaultAuto
	}
	//自动播放间隔时间
	defaultAutoTime := 0
	auto_time, err5 := strconv.Atoi(ctx.PostForm("auto_time"))
	if err5 != nil {
		auto_time = defaultAutoTime
	}
	// 获取创建人
	var defaultCreatedId int64 = 1
	created_id, err6 := strconv.ParseInt(ctx.PostForm("created_id"), 10, 64)
	if err6 != nil {
		created_id = defaultCreatedId
	}

	// 判断媒体广告位是否已提交
	ishave := 0
	Table.DspMediaPositionDB.Model(&Table.DspMediaPosition{}).Where("media_id=?", media_id).Where("position_name=?", position_name).Where("position_type=?", position_type).Where("position_size=?", position_size).Count(&ishave)
	if ishave > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在该媒体广告位类型，请勿重复提交！",
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
				"message": "添加失败，请检查参数！",
			})
		}
	}()
	//检测是否媒体id有效
	var firstModel Table.DspMedia
	if err := Table.DspMediaDB.Where("media_id=?", media_id).First(&firstModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "无效ID",
		})
		return
	}
	// 在这里执行数据库操作
	var insertModel Table.DspMediaPosition
	insertModel.MediaId = uint64(media_id)
	insertModel.PositionName = position_name
	insertModel.PositionType = position_type
	insertModel.ConnectType = connect_type
	insertModel.PositionSize = position_size
	insertModel.IsAccurate = uint(is_accurate)
	insertModel.PositionMaterialType = uint(position_material_type)
	insertModel.IsRing = uint(is_ring)
	insertModel.IsAuto = uint(is_auto)
	insertModel.AutoTime = uint(auto_time)
	insertModel.CreatedId = uint64(created_id)

	if err := Table.DspMediaPositionDB.Create(&insertModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "添加失败！",
		})
		return
	}
	// 媒体表广告位+1
	firstModel.MediaAdCount += 1
	if err := Table.DspMediaDB.Model(&Table.DspMedia{}).Where("media_id=?", firstModel.MediaId).Update(firstModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "添加失败！",
		})
		return
	}

}

// 合作媒体广告位的删除
func DspMediaPositionDelete(ctx *gin.Context) {

}

// 合作媒体广告位的更新
func DspMediaPositionPut(ctx *gin.Context) {
	var defaultMediaPositionId int64 = 0
	media_position_id, err := strconv.ParseInt(ctx.PostForm("media_position_id"), 10, 64)
	if err != nil {
		media_position_id = defaultMediaPositionId
	}

	var defaultMediaId int64 = 0
	media_id, err0 := strconv.ParseInt(ctx.PostForm("media_id"), 10, 64)
	if err0 != nil {
		media_id = defaultMediaId
	}

	position_name, _ := ctx.GetPostForm("position_name")
	position_type, _ := ctx.GetPostForm("position_type")
	position_size, _ := ctx.GetPostForm("position_size")
	connect_type, _ := ctx.GetPostForm("connect_type")

	// 是否精准匹配
	defaultAccurate := 0
	is_accurate, err1 := strconv.Atoi(ctx.PostForm("is_accurate"))
	if err1 != nil {
		is_accurate = defaultAccurate
	}
	//素材类型
	defaultType := 0
	position_material_type, err2 := strconv.Atoi(ctx.PostForm("position_material_type"))
	if err2 != nil {
		position_material_type = defaultType
	}
	//循环播放类型
	defaultRing := 0
	is_ring, err3 := strconv.Atoi(ctx.PostForm("is_ring"))
	if err3 != nil {
		is_ring = defaultRing
	}
	//是否自动播放
	defaultAuto := 1
	is_auto, err4 := strconv.Atoi(ctx.PostForm("is_auto"))
	if err4 != nil {
		is_auto = defaultAuto
	}
	//自动播放间隔时间
	defaultAutoTime := 0
	auto_time, err5 := strconv.Atoi(ctx.PostForm("auto_time"))
	if err5 != nil {
		auto_time = defaultAutoTime
	}

	//状态
	defaultStatus := 0
	status, err6 := strconv.Atoi(ctx.PostForm("status"))
	if err6 != nil {
		status = defaultStatus
	}

	//获取对应的媒体类型及广告位
	var firstModel Table.DspMediaPosition
	if err := Table.DspMediaPositionDB.Where("media_position_id=?", media_position_id).First(&firstModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "无效ID",
		})
		return
	}

	if position_name != "" {
		firstModel.PositionName = position_name
	}
	if position_type != "" {
		firstModel.PositionType = position_type
	}
	if position_size != "" {
		firstModel.PositionSize = position_size
	}
	if connect_type != "" {
		firstModel.ConnectType = connect_type
	}

	default_media_id := firstModel.MediaId
	firstModel.MediaId = uint64(media_id)
	firstModel.IsAccurate = uint(is_accurate)
	firstModel.PositionMaterialType = uint(position_material_type)
	firstModel.IsRing = uint(is_ring)
	firstModel.IsAuto = uint(is_auto)
	firstModel.AutoTime = uint(auto_time)
	firstModel.Status = uint(status)

	// 判断是否重复
	isAnother := 0
	Table.DspMediaPositionDB.Model(&Table.DspMediaPosition{}).Where("media_position_id!=?", firstModel.MediaPositionId).Where("media_id=?", firstModel.MediaId).Where("position_name=?", firstModel.PositionName).Where("position_type=?", firstModel.PositionType).Where("position_size=?", firstModel.PositionSize).Count(&isAnother)
	if isAnother > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "已存在该广告位！",
		})
		return
	}

	// 更新数据
	if err := Table.DspMediaPositionDB.Model(&Table.DspMediaPosition{}).Where("media_position_id=?", firstModel.MediaPositionId).Update(firstModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "更新失败",
		})
		return
	}
	// 媒体ID不一致时，原来的广告位减少，新的增加
	var defaultMediaModel Table.DspMedia
	if err := Table.DspMediaDB.Where("media_id=?", default_media_id).First(&defaultMediaModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "无效ID",
		})
		return
	}
	var NewMediaModel Table.DspMedia
	if err := Table.DspMediaDB.Where("media_id=?", firstModel.MediaId).First(&NewMediaModel).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    999,
			"message": "无效ID",
		})
		return
	}
	if default_media_id != firstModel.MediaId {
		defaultMediaModel.MediaAdCount -= 1
		if err := Table.DspMediaDB.Model(&Table.DspMedia{}).Where("media_id=?", default_media_id).Update(defaultMediaModel).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    999,
				"message": "添加失败！",
			})
			return
		}

		NewMediaModel.MediaAdCount += 1
		if err := Table.DspMediaDB.Model(&Table.DspMedia{}).Where("media_id=?", firstModel.MediaId).Update(NewMediaModel).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    999,
				"message": "添加失败！",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    990,
		"message": "成功",
	})

}

func DspMediaPosition(ctx *gin.Context) {
	/**
	1、Media position 的相关操作，按照rest api方式
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		DspMediaPositionGet(ctx)
		break
	case "POST":
		DspMediaPositionPost(ctx)
		break
	case "PUT":
		DspMediaPositionPut(ctx)
		break
	case "DELETE":
		DspMediaPositionDelete(ctx)
		break
	default:
		break
	}
}
