package middlewares

import (
	"dsp_program_api/models/Redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func CheckAuth(args ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token
		Cookie := ctx.GetHeader("token")
		if Cookie == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    904,
				"message": "登录已失效！",
			})
			ctx.Abort()
		}
		// 判断是否存在
		HashKey := fmt.Sprintf("laravel_database_OpenId:%s", Cookie)
		_, err := Redis.RedisCon.HGet(ctx, HashKey, "id").Result()
		if err == redis.Nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    905,
				"message": "登录已失效！",
			})
			ctx.Abort()
		} else if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    906,
				"message": "登录已失效！",
			})
			ctx.Abort()
		}
		ctx.Next()

	}
}
