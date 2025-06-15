package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirwan77/goLang/entity"
	"github.com/nirwan77/goLang/service"
)

type VideoController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	ctx.JSON(http.StatusOK, videos)
}

func (c *controller) Save(ctx *gin.Context) {
	var video entity.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.service.Save(video)
	ctx.JSON(http.StatusOK, video)
}
