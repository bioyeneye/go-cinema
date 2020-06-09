package entities

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	VideoId         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `json:"title" binding:"required" gorm:"type:varchar(100)"`
	Description string `json:"description" binding:"required,min=20" gorm:"type:varchar(200)"`
	Url         string `json:"url" binding:"required" gorm:"type:varchar(256);UNIQUE"`
}
