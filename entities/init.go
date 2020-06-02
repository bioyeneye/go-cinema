package entities

import (
	"fmt"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	dbConfiguration := utilities.NewDatabaseConfig()

	viperUser := dbConfiguration.Username
	viperPassword := dbConfiguration.Password
	viperDb := dbConfiguration.Name
	viperHost := dbConfiguration.Host
	viperPort := dbConfiguration.Port

	progressSqlConnectionString := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viperHost, viperPort, viperUser, viperDb, viperPassword)
	fmt.Println("Connection String is\t\t", progressSqlConnectionString)

	db, err := gorm.Open("postgres", progressSqlConnectionString)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Video{})
	return db
}
