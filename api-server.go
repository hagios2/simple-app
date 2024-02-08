package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/hagios2/simple-app/controller"
	"github.com/hagios2/simple-app/middlewares"
	"github.com/hagios2/simple-app/repository"
	"github.com/hagios2/simple-app/service"
	"github.com/hagios2/simple-app/validators"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

var (
	videoRepository                            = repository.NewVideoRepository()
	videoService                               = service.New(videoRepository)
	jwtService                                 = service.NewJWTService()
	loginService                               = service.NewLoginService()
	videoController                            = controller.New(videoService)
	loginController controller.LoginController = controller.NewLogin(loginService, jwtService)
)

// SetupLogoutput function allows to specify a log file for all request
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	defer videoRepository.CloseDB()

	server := gin.New()

	//Loading static files and html contents
	server.Static("/css", "templates/css")
	server.LoadHTMLGlob("templates/*.html")

	//middleware
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	var ctx gin.Context

	if v, ok := binding.Validator.Engine().(validator.Validate); ok {
		err := v.RegisterValidation("is-cool", validators.ValidateCoolTitle)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	//setting up routes
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": " pong",
		})
	})

	//setting up routes
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}
	})

	//API ROUTES
	videoGroupRouter := server.Group("/api/videos", middlewares.AuthorizeJWT())
	{
		videoGroupRouter.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})
		videoGroupRouter.POST("/", func(ctx *gin.Context) {
			video, err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, video)
		})
		videoGroupRouter.PUT("/:id", func(ctx *gin.Context) {
			video, err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, video)
		})
		videoGroupRouter.DELETE("/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, nil)
		})
	}

	//VIEW ROUTES
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
