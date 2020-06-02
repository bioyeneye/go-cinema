package identity

import (
	"github.com/bioyeneye/rest-gin-api/core/identity/managers/base"
	"github.com/bioyeneye/rest-gin-api/core/identity/models"
)

type IUserValidator interface {
	Validate(base base.UserStoreBase, user models.IdentityUser) models.IdentityResult
	ValidateUserName(base base.UserStoreBase, user models.IdentityUser, errors []models.IdentityError)
	ValidateUserEmail(base base.UserStoreBase, user models.IdentityUser, errors []models.IdentityError)
}

type UserValidator struct {
	Describer IIdentityErrorDescriber
}

func (u UserValidator) Validate(base base.UserStoreBase, user models.IdentityUser) models.IdentityResult {
	panic("implement me")
}

func (u UserValidator) ValidateUserName(base base.UserStoreBase, user models.IdentityUser, errors []models.IdentityError) {
	panic("implement me")
}

func (u UserValidator) ValidateUserEmail(base base.UserStoreBase, user models.IdentityUser, errors []models.IdentityError) {
	panic("implement me")
}

//Creates a new user validator instant
//	if you don't have an implementation of the IIdentityErrorDescriber, pass nil then NewIdentityErrorDescriber() is used automatically
func NewUserValidator(descriptor IIdentityErrorDescriber) IUserValidator {

	if descriptor == nil{
		descriptor = NewIdentityErrorDescriber()
	}

	return &UserValidator{
		Describer: descriptor,
	}
}
