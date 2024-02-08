package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hagios2/simple-app/entity"
	"github.com/hagios2/simple-app/service"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) (entity.Video, error)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	//validate = validator.New()
	//err := validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	//if err != nil {
	//	return nil
	//}
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return entity.Video{}, err
	}
	c.service.Save(video)
	return video, nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()

	data := gin.H{
		"videos": videos,
		"title":  "Video Page",
	}

	ctx.HTML(http.StatusOK, "/templates/index.html", data)
}
