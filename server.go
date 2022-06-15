package main

import (
	"net/http"

	"github.com/coder-abdo/gin-tutorial/controller"
	"github.com/coder-abdo/gin-tutorial/middlewares"
	"github.com/coder-abdo/gin-tutorial/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	server.Use(gin.Recovery(), middlewares.Logger())
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "video is saved successfully"})

		}
	})
	server.Run(":8080")
}
