package middlewares

import (
	"dsp_program_api/models/Table"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

func AddRequestLog(args ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录日志
		ip := ctx.ClientIP()
		nowTimestamp := time.Now().Local()
		var logName string
		var logDesc string
		for i := range args {
			if i == 0 {
				logName = args[0]
			}
			if i == 1 {
				logDesc = args[1]
			}
		}

		method := ctx.Request.Method

		// 处理数据
		params := make(map[string]interface{})

		if method == "GET" {
			for key, values := range ctx.Request.URL.Query() {
				if len(values) > 0 {
					params[key] = values[0]
				}
			}
		} else {
			if err := ctx.Request.Form; err == nil {
				for key, values := range ctx.Request.Form {
					params[key] = values[0]
				}
			} else {
				var jsonData map[string]interface{}
				if err := ctx.BindJSON(&jsonData); err == nil {
					for key, value := range jsonData {
						params[key] = value
					}
				}
			}
		}

		jsonDatas, _ := json.Marshal(params)
		requestBody := string(jsonDatas)

		// 放入日志表
		defer func() {
			e := recover()
			if e == nil {
				//插入数据
				Table.RequestLogDB.Create(&Table.RequestLog{
					Ip:          ip,
					Method:      method,
					Name:        logName,
					Desc:        logDesc,
					RequestBody: requestBody,
					CreatedAt:   nowTimestamp,
				})
			} else {
				//插入数据
				Table.RequestLogDB.Create(&Table.RequestLog{
					Ip:          ip,
					Method:      method,
					Name:        logName,
					Desc:        logDesc,
					RequestBody: requestBody,
					CreatedAt:   nowTimestamp,
				})
			}
		}()

		ctx.Next()
	}
}
