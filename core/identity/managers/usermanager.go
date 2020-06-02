package managers

import (
	"github.com/bioyeneye/rest-gin-api/core/identity"
	"github.com/bioyeneye/rest-gin-api/core/identity/managers/base"
	"github.com/bioyeneye/rest-gin-api/core/identity/models"
	"time"
)

type UserManager struct {
	hasher 			identity.IPasswordHasher
	validator			identity.UserValidator
}

func NewUserManager(hasher identity.PasswordHasher, validator identity.UserValidator) base.UserStoreBase {
	return &UserManager{
		hasher: hasher,
		validator:  validator,
	}
}

func (u *UserManager) GetUserId(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) GetUserName(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) SetUserName(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) GetNormalizedUserName(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) SetNormalizedUserName(user models.IdentityUser, normalizedName string) {
	panic("implement me")
}

func (u *UserManager) Create(user models.IdentityUser) base.IdentityResult {
	panic("implement me")
}

func (u *UserManager) Update(user models.IdentityUser) base.IdentityResult {
	panic("implement me")
}

func (u *UserManager) Delete(user models.IdentityUser) base.IdentityResult {
	panic("implement me")
}

func (u *UserManager) FindById(id string) models.IdentityUser {
	panic("implement me")
}

func (u *UserManager) FindByName(normalizedName string) models.IdentityUser {
	panic("implement me")
}

func (u *UserManager) SetPasswordHash(user models.IdentityUser, passwordHash string) base.IdentityResult {
	panic("implement me")
}

func (u *UserManager) GetPasswordHash(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) HasPassword(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) FindUserLogin(userId string) models.IdentityUser {
	panic("implement me")
}

func (u *UserManager) FindByEmail(normalizedEmail string) models.IdentityUser {
	panic("implement me")
}

func (u *UserManager) GetLockoutEndDate(user models.IdentityUser) time.Time {
	panic("implement me")
}

func (u *UserManager) SetLockoutEndDate(user models.IdentityUser, lockoutEnd time.Time) {
	panic("implement me")
}

func (u *UserManager) IncrementAccessFailedCount(user models.IdentityUser) int {
	panic("implement me")
}

func (u *UserManager) ResetAccessFailedCount(user models.IdentityUser) {
	panic("implement me")
}

func (u *UserManager) GetAccessFailedCount(user models.IdentityUser) int {
	panic("implement me")
}

func (u *UserManager) GetLockoutEnabled(user models.IdentityUser) bool {
	panic("implement me")
}

func (u *UserManager) SetLockoutEnabled(user models.IdentityUser, enabled bool) {
	panic("implement me")
}

func (u *UserManager) SetPhoneNumber(user models.IdentityUser, phoneNumber string) {
	panic("implement me")
}

func (u *UserManager) GetPhoneNumber(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) GetPhoneNumberConfirmed(user models.IdentityUser) bool {
	panic("implement me")
}

func (u *UserManager) SetPhoneNumberConfirmed(user models.IdentityUser, confirmed bool) {
	panic("implement me")
}

func (u *UserManager) SetSecurityStamp(user models.IdentityUser, stamp string) {
	panic("implement me")
}

func (u *UserManager) GetSecurityStamp(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) SetTwoFactorEnabledAsync(user models.IdentityUser, enabled bool) {
	panic("implement me")
}

func (u *UserManager) GetTwoFactorEnabledAsync(user models.IdentityUser) string {
	panic("implement me")
}

func (u *UserManager) GetEmailConfirmed(user models.IdentityUser) bool {
	panic("implement me")
}

func (u *UserManager) SetEmailConfirmed(user models.IdentityUser, confirmed bool) {
	panic("implement me")
}

func (u *UserManager) SetEmail(user base.IdentityResult, email string) {
	panic("implement me")
}

func (u *UserManager) GetEmailAsync(user base.IdentityResult) string {
	panic("implement me")
}

func (u *UserManager) SetNormalizedEmail(user base.IdentityResult, email string) {
	panic("implement me")
}

func (u *UserManager) GetNormalizedEmailAsync(user base.IdentityResult) string {
	panic("implement me")
}