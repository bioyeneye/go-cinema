package db

import (
	entities2 "github.com/bioyeneye/rest-gin-api/db/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDbModels(dialect string, dbConString string, entities ...interface{}) (*gorm.DB, error) {

	db, err := gorm.Open(dialect, dbConString)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities2.Video{})
	return db, err
}
