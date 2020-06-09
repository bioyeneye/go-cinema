package utilities

import "github.com/bioyeneye/rest-gin-api/core/configs"

type DBConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

func NewDBConfig(host string, username string, password string, name string, port string) *DBConfig {
	return &DBConfig{Host: host, Username: username, Password: password, Name: name, Port: port}
}

func NewDBConfigFromEnv() *DBConfig {
	return &DBConfig{
		Host:     GetEnv(configs.DatabaseServer, ""),
		Username: GetEnv(configs.DatabaseUsernameKey, ""),
		Password: GetEnv(configs.DatabasePasswordKey, ""),
		Name:     GetEnv(configs.DatabaseNameKey, ""),
		Port:     GetEnv(configs.DatabasePortKey, ""),
	}
}
