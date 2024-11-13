package controllers

import (
	"dsp_program_api/models/Table"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 广告账户的获取
func AccountGet(ctx *gin.Context) {

}

// 广告账户的创建
func AccountPost(ctx *gin.Context) {
	// 广告账户
	//插入数据
	advertiser_name, _ := ctx.GetPostForm("name")
	var defaultCompanyId int64 = 1
	company_id, err := strconv.ParseInt(ctx.PostForm("company_id"), 10, 64)
	if err != nil {
		company_id = defaultCompanyId
	}
	//company_id := uint64(company_id)
	defaultStatus := 1
	status, err1 := strconv.Atoi(ctx.PostForm("status"))
	if err1 != nil {
		status = defaultStatus
	}

	defer func() {

		e := recover()
		if e != nil {
			Table.AdvertiserAccountDB.Create(&Table.AdvertiserAccount{
				Name:      advertiser_name,
				CompanyId: uint64(company_id),
				Status:    uint(status),
				CreatedId: 1,
			})
		} else {
			fmt.Println(12111)
		}

	}()

}

// 广告账户的删除
func AccountDelete(ctx *gin.Context) {

}

// 广告账户的更新
func AccountPut(ctx *gin.Context) {

}

func Account(ctx *gin.Context) {
	/**
	1、campaign 的相关操作，按照rest api方式
	*/
	method := ctx.Request.Method
	switch method {
	case "GET":
		AccountGet(ctx)
		break
	case "POST":
		AccountPost(ctx)
		break
	case "PUT":
		AccountPut(ctx)
		break
	case "DELETE":
		AccountDelete(ctx)
		break
	default:
		break
	}
}
