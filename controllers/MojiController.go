package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const AdIdBrand = "brand"       // 广告位标识-品牌
const AdIdShopping = "shopping" // 广告位标识-电商

func GetCall(ctx *gin.Context) {
	/**
	1、获取请求数据
	2、处理业务竞价逻辑
	3、无广告填充返回{ "code",400}，正常返回code 200 及竞价信息
	*/
	var isPost bool = true
	fmt.Println("进入墨迹控制器")
	adid, _ := ctx.GetQuery("adid") //  广告位标识
	if adid != AdIdBrand {          //非品牌广告直接不参与
		isPost = false
	}

	sessionid, _ := ctx.GetQuery("sessionid") // 墨迹每次请求的唯一标识
	adtype := ctx.GetInt("adtype")            //广告类型
	adstyle := ctx.GetInt("adstyle")          //广告样式
	basic_price := ctx.GetInt("basic_price")  //最低出价
	//pkgname, _ := ctx.GetQuery("pkgname")                       //应用的包名
	//appname := ctx.DefaultQuery("appname", "mojiWeather")       //APP 名称 固定值 mojiWeather
	//net, _ := ctx.GetQuery("net")                               //连网方式
	//carrier, _ := ctx.GetQuery("carrier")                       //运营商
	//os, _ := ctx.GetQuery("os")                                 //操作系统类型
	//osv, _ := ctx.GetQuery("osv")                               //操作系统版本

	//device, _ := ctx.GetQuery("device")                         //设备品牌和型号
	//ua, _ := ctx.GetQuery("ua")                                 //用户浏览器标识
	//ip, _ := ctx.GetQuery("ip")                                 //用户来源 IP
	//imei, _ := ctx.GetQuery("imei")                             //Android 系统的设备号
	//wma, _ := ctx.GetQuery("wma")                               //终端网卡的 MAC 地址
	//andid, _ := ctx.GetQuery("andid")                           //用户终端的 Android ID
	//idfa, _ := ctx.GetQuery("idfa")                             //仅 iOS 6.0 以上系统的 IDFA
	//oaid, _ := ctx.GetQuery("oaid")                             //Android 系统的匿名设备标识符oaid
	//openudid, _ := ctx.GetQuery("openudid")                     //iOS 终端设备的 OpenUDID
	//unqid, _ := ctx.GetQuery("unqid")                           //Andorid、iOS 操作系统设备的唯一标示码
	//scrro, _ := ctx.GetQuery("scrro")                           //屏幕方向，广告请求时用户屏幕的方向
	//scrwidth, _ := ctx.GetQuery("scrwidth")                     //屏幕宽度
	//scrheight, _ := ctx.GetQuery("scrheight")                   //屏幕高度
	//debug, _ := ctx.GetQuery("debug")                           //debug 模式
	//lon, _ := ctx.GetQuery("lon")                               //地理位置精度
	//lat, _ := ctx.GetQuery("lat")                               //地理位置纬度
	//comment, _ := ctx.GetQuery("comment")                       //当为 pdb/pd 或者需要使用dealid场景时使用，广告主可定义；
	//feed_support_types, _ := ctx.GetQuery("feed_support_types") //该广告位支持的 Feed 类型, 仅 feed 广告位使用
	//show_date, _ := ctx.GetQuery("show_date")                   //广告的展示日期
	//user_tags, _ := ctx.GetQuery("user_tags")                   //用户标签（墨迹）
	//adstyle_request, _ := ctx.GetQuery("adstyle_request")       //除 2048,4096,8192 广告位支持的样式多个类型分号隔开
	//weather_code, _ := ctx.GetQuery("weather_code")             //天气参数
	//boot_mark, _ := ctx.GetQuery("boot_mark")                   //系统启动标识，原值传输
	//update_mark, _ := ctx.GetQuery("update_mark")               //系统更新标识，原值传输
	//brand, _ := ctx.GetQuery("brand")                           //手机品牌
	//request_time, _ := ctx.GetQuery("request_time")             //接口请求时间戳
	//andaid, _ := ctx.GetQuery("andaid")                         //Android Advertiser ID
	//update_time, _ := ctx.GetQuery("update_time")               //系统更新时间；仅IOS端传值
	//substyle, _ := ctx.GetQuery("substyle")                     //双feed信息流必传
	//feed_style, _ := ctx.GetQuery("feed_style")                 //支持多种样式
	//birth_time, _ := ctx.GetQuery("birth_time")                 //设备初始化时间

	dataObj := make(map[string]interface{})
	// 返回数据 视频有3种方式：H5、native、短视频，图片有动态和静态
	dataObj["adid"] = adid
	dataObj["sessionid"] = sessionid
	dataObj["type"] = 2 // 1 图片 2视频
	dataObj["price"] = basic_price * 100
	dataObj["chargingtype"] = "CPM"  // 品牌出价CPM
	dataObj["videoUrl"] = "sss"      // 视频地址
	dataObj["videoImageUrl"] = "sss" // 视频封面图
	dataObj["fileSize"] = 15         // 最大15M
	dataObj["videoLogo"] = "xxxx"    // 素材格式
	dataObj["iconDesc"] = "xxxx"     // logo 描述
	dataObj["shortVideoType"] = 1    // 0 立即下载 1 查看详情
	dataObj["adwidth"] = 900         // 素材宽度
	dataObj["adheight"] = 1200       // 素材高度
	dataObj["adtitle"] = "sss"
	dataObj["adtext"] = "sss"
	dataObj["adtype"] = adtype   // 广告类型 2banner 3开屏
	dataObj["adstyle"] = adstyle // 广告样式

	// 判断是否有广告填充
	if isPost {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": dataObj,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
		})
	}
}
