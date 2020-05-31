package utilities

import "github.com/bioyeneye/rest-gin-api/core/configs"

type DatabaseConfig struct {
	Database string
	Username    string
	Password    string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Database: GetEnv(configs.DatabaseServer, ""),
		Username:   GetEnv(configs.DatabaseUsernameKey, ""),
		Password:   GetEnv(configs.DatabasePasswordKey, ""),
	}
}

