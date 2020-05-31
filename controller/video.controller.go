package controller

import (
	controllerinterface "github.com/bioyeneye/rest-gin-api/controller/interfaces"
	"github.com/bioyeneye/rest-gin-api/entities"
	"github.com/bioyeneye/rest-gin-api/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type VideoController struct {
	service service.IVideoService
}

func (api *APIRoutes) InitVideoRoutes() {
	var videoService = service.New()
	var videoController = NewVideoController(videoService)

	api.BaseRoutes.Video.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.GetVideos())
	},)

	api.BaseRoutes.Video.GET("/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.GetVideo(context))
	})

	api.BaseRoutes.Video.POST("", func(context *gin.Context) {

		response := videoController.Post(context)
		if response.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": response.Error.Error(),
			})
			return
		}

		context.JSON(http.StatusCreated, gin.H{
			"message": response.Message,
			"error":   nil,
			"data":    response.Data,
		})
	})
}

func NewVideoController(service service.IVideoService) controllerinterface.IVideoController {
	return &VideoController{
		service: service,
	}
}

func (videoController *VideoController) GetVideos() []entities.Video {
	return videoController.service.FindAll()
}

func (videoController *VideoController) GetVideo(context *gin.Context) entities.Video {

	idparam := context.Param("id")

	//firstname := context.DefaultQuery("name", "name")

	id, err := strconv.Atoi(idparam)

	if err != nil {
		return entities.Video{}
	}

	videos := videoController.service.FindAll()
	log.Println(id, idparam, videos)

	if videos == nil {
		return entities.Video{}
	}

	videoIndex := sort.Search(len(videos), func(i int) bool {
		return videos[i].Id == id
	})

	log.Println(id, idparam, videos, videoIndex)

	if videoIndex >= len(videos) {
		return entities.Video{}
	}

	return videos[videoIndex]
}

func (videoController *VideoController) Post(context *gin.Context) controllerinterface.VideoApiResponse {
	var video entities.Video
	err := context.ShouldBindJSON(&video)

	if err != nil {
		//context.Writer.WriteHeader(http.StatusInternalServerError)
		//context.Writer.Write([]byte(`{"error": "Error unmarshalling the post"}`))
		return controllerinterface.VideoApiResponse{
			Error:   err,
			Data:    entities.Video{},
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
