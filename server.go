package main

import (
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
		ctx.JSON(200, VideoController.Save(ctx))
	})
	server.Run(":8080")
}
