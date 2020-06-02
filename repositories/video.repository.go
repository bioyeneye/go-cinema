package repositories

import (
	"github.com/bioyeneye/rest-gin-api/entities"
	repositoryinterfaces "github.com/bioyeneye/rest-gin-api/repositories/interfaces"
	"github.com/jinzhu/gorm"
	"log"
)

type VideoRepository struct {
	db *gorm.DB
}

func (repository VideoRepository) Find() entities.Video {
	panic("implement me")
}

func (repository VideoRepository) Save(video entities.Video) entities.Video {
	savedVideo := entities.Video{
		Title:       video.Title,
		Description: video.Description,
		Url:         video.Url,
	}
	repository.db.Create(&savedVideo)

	log.Println(savedVideo.ID)
	return savedVideo
}

func (repository VideoRepository) Update(video entities.Video) entities.Video {
	repository.db.Update(video)
	return video
}

func (repository VideoRepository) Delete(video entities.Video) bool {
	repository.db.Delete(video)
	return true
}

func (repository VideoRepository) All() []entities.Video {
	var videos []entities.Video
	repository.db.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

func NewVideoRepository(db *gorm.DB) repositoryinterfaces.IVideoRepository {
	return &VideoRepository{
		db: db,
	}
}


