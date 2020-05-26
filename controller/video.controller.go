package controller

import (
	controllerinterface "github.com/bioyeneye/rest-gin-api/controller/interfaces"
	"github.com/bioyeneye/rest-gin-api/entity"
	"github.com/bioyeneye/rest-gin-api/service"
	"github.com/gin-gonic/gin"
	"log"
	"sort"
	"strconv"
)

type VideoController struct {
	service service.IVideoService
}

func NewVideoController(service service.IVideoService) controllerinterface.IVideoController {
	return &VideoController{
		service: service,
	}
}

func (videoController *VideoController) GetVideos() []entity.Video {
	return videoController.service.FindAll()
}

func (videoController *VideoController) GetVideo(context *gin.Context) entity.Video {
	idparam := context.Param("id")

	//firstname := context.DefaultQuery("name", "name")

	id, err := strconv.Atoi(idparam)


	if err != nil {
		return entity.Video{}
	}

	videos := videoController.service.FindAll()
	log.Println(id, idparam, videos)

	if videos == nil{
		return entity.Video{}
	}

	videoIndex := sort.Search(len(videos), func(i int) bool {
		return videos[i].Id == id
	})

	log.Println(id, idparam, videos, videoIndex)

	if videoIndex >= len(videos) {
		return entity.Video{}
	}

	return videos[videoIndex]
}

func (videoController *VideoController) Post(context *gin.Context) controllerinterface.VideoApiResponse {
	var video entity.Video
	err := context.ShouldBindJSON(&video)



	if err != nil {
		//context.Writer.WriteHeader(http.StatusInternalServerError)
		//context.Writer.Write([]byte(`{"error": "Error unmarshalling the post"}`))
		return controllerinterface.VideoApiResponse{
			Error:   err,
			Data:    entity.Video{},
			Message: "Error with validation",
		}
	}

	videoController.service.Save(video)
	return controllerinterface.VideoApiResponse{
		Error:   nil,
		Data:    video,
		Message: "Videos created successfully",
	}
}


