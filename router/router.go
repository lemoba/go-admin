package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lemoba/go-admin/api/v1/admin"
	"github.com/lemoba/go-admin/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	api := r.Group("/bv8/api/")
	{
		// 健康检查
		api.GET("health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code": 200,
				"data": "web and db is running!",
			})
		})

		// 用户管理
		api = api.Group("admin")
		{
			api.GET("/users", admin.GetUsers)
		}
	}
	return r
}
