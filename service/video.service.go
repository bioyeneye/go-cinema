package service

import (
	"github.com/bioyeneye/rest-gin-api/entities"
	"github.com/bioyeneye/rest-gin-api/repositories"
	repositoryinterfaces "github.com/bioyeneye/rest-gin-api/repositories/interfaces"
	"github.com/jinzhu/gorm"
)

type IVideoService interface {
	Find() entities.Video
	Save(video entities.Video) entities.Video
	Update(video entities.Video) entities.Video
	Delete(video entities.Video) bool
	FindAll() []entities.Video
}

func New(db *gorm.DB) IVideoService {
	return &VideoService{
		videoRepository: repositories.NewVideoRepository(db),
	}
}

type VideoService struct {
	videoRepository repositoryinterfaces.IVideoRepository
}

func (service *VideoService) Find() entities.Video {
	panic("implement me")
}

func (service *VideoService) Update(video entities.Video) entities.Video {
	panic("implement me")
}

func (service *VideoService) Delete(video entities.Video) bool {
	panic("implement me")
}

func (service *VideoService) Save(video entities.Video) entities.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *VideoService) FindAll() []entities.Video {
	return service.videoRepository.All()
}