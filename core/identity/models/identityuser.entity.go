package models

import (
	"database/sql"
	"time"
)

type IdentityUser struct {
	ID        				string `gorm:"primary_key;unique"`
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
	DeletedAt 				*time.Time `sql:"index"`

	Email                	string `gorm:"type:varchar(100);unique_index"`
	NormalizedEmail         string `gorm:"type:varchar(100);unique_index"`
	EmailConfirmed       	bool
	PasswordHash         	string
	ConcurrencyStamp        	string //A random value that must change whenever a user is persisted to the store
	PhoneNumber          	string
	PhoneNumberConfirmed 	bool
	TwoFactorEnabled     	bool

	//DateTime in UTC when lockout ends, any time in the past is considered not locked out.
	LockoutEndDateUtc 		sql.NullTime

	// Is lockout enabled for this user
	LockoutEnabled 			bool

	AccessFailedCount 		int
	UserName          		string
	NormalizedUserName		string
}