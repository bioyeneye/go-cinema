package repositoryinterfaces

import (
	"github.com/bioyeneye/rest-gin-api/db/entities"
)

type IVideoRepository interface {
	Find() entities.Video
	Save(video entities.Video) entities.Video
	Update(video entities.Video) entities.Video
	Delete(video entities.Video) bool
	All() []entities.Video
}
