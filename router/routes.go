package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Show Opening
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Opening",
			})
		})
		// POST Opening
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Post Opening",
			})
		})
		// PUT Opening
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})
		// DELETE Opening
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "DELETE Opening",
			})
		})
		// Show Openings
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Openings",
			})
		})
	}
}
