package models

import (
	"time"
)

type IdentityRole struct {
	ID        				string `gorm:"primary_key;unique"`
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
	DeletedAt 				*time.Time `sql:"index"`

	ConcurrencyStamp        	string //A random value that must change whenever a user is persisted to the store
	Name          		string
	NormalizedName		string
}
