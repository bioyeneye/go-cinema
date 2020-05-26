package routes

import (
	"github.com/bioyeneye/rest-gin-api/controller"
	"github.com/bioyeneye/rest-gin-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//routes
func SetVideoRoutes(router *gin.RouterGroup) []gin.IRoutes  {

	var videoService = service.New()
	var videoController = controller.NewVideoController(videoService)

	return []gin.IRoutes {
		router.GET("/videos", func(context *gin.Context) {
			context.JSON(http.StatusOK, videoController.GetVideos())
		}),

		router.GET("/videos/:id", func(context *gin.Context) {
			context.JSON(http.StatusOK, videoController.GetVideo(context))
		}),

		router.POST("/videos", func(context *gin.Context) {

			response := videoController.Post(context)
			if response.Error != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"error": response.Error.Error(),
				})
				return
			}

			context.JSON(http.StatusCreated, gin.H{
				"message": response.Message,
				"error": nil,
				"data": response.Data,
			})
		}),
	}
}
