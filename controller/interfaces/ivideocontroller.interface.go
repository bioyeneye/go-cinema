package interfaces

import (
	"github.com/bioyeneye/rest-gin-api/entity"
	"github.com/gin-gonic/gin"
)

type VideoApiResponse struct {
	Error   error
	Data    entity.Video
	Message string
}

type IVideoController interface {
	GetVideos() []entity.Video
	GetVideo(context *gin.Context) entity.Video
	Post(context *gin.Context) VideoApiResponse
}
