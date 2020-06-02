package base

import (
	"github.com/bioyeneye/rest-gin-api/core/identity/models"
	"time"
)

type UserStoreBase interface {
	GetUserId(user models.IdentityUser) string
	GetUserName(user models.IdentityUser) string
	SetUserName(user models.IdentityUser) string
	GetNormalizedUserName(user models.IdentityUser) string
	SetNormalizedUserName(user models.IdentityUser, normalizedName string)

	Create(user models.IdentityUser) IdentityResult
	Update(user models.IdentityUser) IdentityResult
	Delete(user models.IdentityUser) IdentityResult

	FindById(id string) models.IdentityUser
	FindByName(normalizedName string) models.IdentityUser
	SetPasswordHash(user models.IdentityUser, passwordHash string) IdentityResult
	GetPasswordHash(user models.IdentityUser) string
	HasPassword(user models.IdentityUser) string

	FindUserLogin(userId string) models.IdentityUser
	FindByEmail(normalizedEmail string) models.IdentityUser
	GetLockoutEndDate(user models.IdentityUser) time.Time
	SetLockoutEndDate(user models.IdentityUser, lockoutEnd time.Time)
	IncrementAccessFailedCount(user models.IdentityUser) int
	ResetAccessFailedCount(user models.IdentityUser)
	GetAccessFailedCount(user models.IdentityUser) int
	GetLockoutEnabled(user models.IdentityUser) bool
	SetLockoutEnabled(user models.IdentityUser, enabled bool)
	SetPhoneNumber(user models.IdentityUser, phoneNumber string)
	GetPhoneNumber(user models.IdentityUser) string
	GetPhoneNumberConfirmed(user models.IdentityUser) bool
	SetPhoneNumberConfirmed(user models.IdentityUser, confirmed bool)

	SetSecurityStamp(user models.IdentityUser, stamp string)
	GetSecurityStamp(user models.IdentityUser) string

	SetTwoFactorEnabledAsync(user models.IdentityUser, enabled bool)
	GetTwoFactorEnabledAsync(user models.IdentityUser) string


	GetEmailConfirmed(user models.IdentityUser) bool
	SetEmailConfirmed(user models.IdentityUser, confirmed bool)
	SetEmail(user IdentityResult, email string)
	GetEmailAsync(user IdentityResult) string
	SetNormalizedEmail(user IdentityResult, email string)
	GetNormalizedEmailAsync(user IdentityResult) string
}
