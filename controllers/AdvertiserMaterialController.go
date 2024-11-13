package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path"
)

// 素材的获取
func MaterialGet(ctx *gin.Context) {

}

// 素材的创建
func MaterialPost(ctx *gin.Context) {
	// 单文件上传
	file, _ := ctx.FormFile("file")

	materialName := file.Filename
	ext := path.Ext(materialName)
	extName := ext[1:]

	// 将文件保存到服务器上，并检查错误
	err := ctx.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}

	filesys, err := os.Open(file.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}
	defer filesys.Close()
	// 获取md5
	hash := md5.New()
	if _, err := io.Copy(hash, filesys); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    401,
			"message": "上传失败",
		})
		return
	}

	hashStr := fmt.Sprintf("%x", hash.Sum(nil))

	img, _, err := image.Decode(filesys)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	ctx.JSON(http.StatusOK, gin.H{
		"code":      200,
		"message":   "success",
		"width":     width,
		"height":    height,
		"size":      file.Size,
		"ext_name":  extName,
		"signature": hashStr,
	})

}

// 素材的删除
func MaterialDelete(ctx *gin.Context) {

}

// 素材的更新
func MaterialPut(ctx *gin.Context) {

}

func Material(ctx *gin.Context) {
	/**
	1、素材 的相关操作，按照rest api方式
	2、视频素材需要首帧处理
	3、素材获取长宽高等
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		CampaignGet(ctx)
		break
	case "POST":
		MaterialPost(ctx)
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
