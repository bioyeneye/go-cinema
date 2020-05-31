package utilities

import "github.com/bioyeneye/rest-gin-api/core/configs"

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     GetEnv(configs.DatabaseServer, ""),
		Username: GetEnv(configs.DatabaseUsernameKey, ""),
		Password: GetEnv(configs.DatabasePasswordKey, ""),
		Name:     GetEnv(configs.DatabaseNameKey, ""),
		Port:     GetEnv(configs.DatabasePortKey, ""),
	}
}
