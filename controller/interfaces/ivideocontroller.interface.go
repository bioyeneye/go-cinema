package interfaces

import (
	"github.com/bioyeneye/rest-gin-api/entities"
	"github.com/gin-gonic/gin"
)

type VideoApiResponse struct {
	Error   error
	Data    entities.Video
	Message string
}

type IVideoController interface {
	GetVideos() []entities.Video
	GetVideo(context *gin.Context) entities.Video
	Post(context *gin.Context) VideoApiResponse
}
