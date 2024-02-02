package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hagios2/simple-app/controller"
	"github.com/hagios2/simple-app/middlewares"
	"github.com/hagios2/simple-app/service"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// SetupLogoutput function allows to specify a log file for all request
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	//middleware
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	//setting up routes
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": " pong",
		})
	})

	videoGroupRouter := server.Group("/videos")
	videoGroupRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})
	videoGroupRouter.POST("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
