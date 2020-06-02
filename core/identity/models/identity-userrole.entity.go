package models

import "github.com/jinzhu/gorm"

type IdentityUserRole struct {
	gorm.Model
	UserId string
	RoleId string
}
