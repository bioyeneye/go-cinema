package entities

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDbModels(dialect string, dbConString string, entities ...interface{}) (*gorm.DB, error) {

	db, err := gorm.Open(dialect, dbConString)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Video{})
	return db, err
}
