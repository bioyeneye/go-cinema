package service

import "github.com/bioyeneye/rest-gin-api/entities"

type IVideoService interface {
	Save(video entities.Video) entities.Video
	FindAll() []entities.Video
}

type VideoService struct {
	videos []entities.Video
}

func New() IVideoService {
	return &VideoService{}
}

func (service *VideoService) Save(video entities.Video) entities.Video {
	service.videos = append(service.videos, video)
	return video
}


func (service *VideoService) FindAll() []entities.Video {
	return service.videos
}