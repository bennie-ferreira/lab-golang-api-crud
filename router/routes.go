package router

import (
	"lab-golang-api-crud/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {

	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		// Show Opening
		v1.GET("/opening", handler.ShowOpeningHandler)
		// POST Opening
		v1.POST("/opening", handler.CreateOpeningHandler)
		// PUT Opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		// DELETE Opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		// List Openings
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
