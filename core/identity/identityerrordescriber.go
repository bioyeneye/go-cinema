package identity

import (
	"fmt"
	"github.com/bioyeneye/rest-gin-api/core/identity/helpers"
	"github.com/bioyeneye/rest-gin-api/core/identity/models"
)

type IIdentityErrorDescriber interface {
	DefaultError() models.IdentityError
	ConcurrencyFailure() models.IdentityError
	PasswordMismatch() models.IdentityError
	InvalidUserName(username string) models.IdentityError
	InvalidEmail(email string) models.IdentityError
	DuplicateUserName(username string) models.IdentityError
	DuplicateEmail(email string) models.IdentityError
	InvalidRoleName(rolename string) models.IdentityError
	DuplicateRoleName(rolename string) models.IdentityError
	UserAlreadyHasPassword() models.IdentityError
	UserLockoutNotEnabled() models.IdentityError
	UserAlreadyInRole(rolename string) models.IdentityError
	UserNotInRole(rolename string) models.IdentityError
	PasswordTooShort(length int) models.IdentityError
	PasswordRequiresUniqueChars(length int) models.IdentityError
	PasswordRequiresNonAlphanumeric() models.IdentityError
	PasswordRequiresDigit() models.IdentityError
	PasswordRequiresLower() models.IdentityError
	PasswordRequiresUpper() models.IdentityError
}

type IdentityErrorDescriber struct {
}

func NewIdentityErrorDescriber() IIdentityErrorDescriber {
	return &IdentityErrorDescriber{}
}

func (describer *IdentityErrorDescriber) DefaultError() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.DefaultError),
		Description: "An unknown failure has occurred.",
	}
}

func (describer *IdentityErrorDescriber) ConcurrencyFailure() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.ConcurrencyFailure),
		Description: "Optimistic concurrency failure, object has been modified.",
	}
}

func (describer *IdentityErrorDescriber) PasswordMismatch() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.PasswordMismatch),
		Description: "Incorrect password.",
	}
}

func (describer *IdentityErrorDescriber) InvalidUserName(username string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.InvalidUserName),
		Description: fmt.Sprintf("Username %s is invalid, can only contain letters or digits.", username),
	}
}

func (describer *IdentityErrorDescriber) InvalidEmail(email string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.InvalidEmail),
		Description: fmt.Sprintf("Email %s is invalid.", email),
	}
}

func (describer *IdentityErrorDescriber) DuplicateUserName(username string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.DuplicateUserName),
		Description: fmt.Sprintf("Username %s is already taken.", username),
	}
}

func (describer *IdentityErrorDescriber) DuplicateEmail(email string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.DuplicateEmail),
		Description: fmt.Sprintf("Email %s is already taken.", email),
	}
}

func (describer *IdentityErrorDescriber) InvalidRoleName(rolename string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.InvalidRoleName),
		Description: fmt.Sprintf("Role name %s is invalid.", rolename),
	}
}

func (describer *IdentityErrorDescriber) DuplicateRoleName(rolename string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.DuplicateRoleName),
		Description: fmt.Sprintf("Role name %s is already taken.", rolename),
	}
}

func (describer *IdentityErrorDescriber) UserAlreadyHasPassword() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.UserAlreadyHasPassword),
		Description: "User already has a password set.",
	}
}

func (describer *IdentityErrorDescriber) UserLockoutNotEnabled() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.UserLockoutNotEnabled),
		Description: "Lockout is not enabled for this user.",
	}
}

func (describer *IdentityErrorDescriber) UserAlreadyInRole(role string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.UserAlreadyInRole),
		Description: fmt.Sprintf("User already in role %s.", role),
	}
}

func (describer *IdentityErrorDescriber) UserNotInRole(role string) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.UserNotInRole),
		Description: fmt.Sprintf("User is in not role %s.", role),
	}
}

func (describer *IdentityErrorDescriber) PasswordTooShort(length int) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.PasswordTooShort),
		Description: fmt.Sprintf("Passwords must be at least %d characters.", length),
	}
}

func (describer *IdentityErrorDescriber) PasswordRequiresUniqueChars(length int) models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.PasswordRequiresUniqueChars),
		Description: fmt.Sprintf("Passwords must use at least %d different characters.", length),
	}
}

func (describer *IdentityErrorDescriber) PasswordRequiresNonAlphanumeric() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.PasswordRequiresNonAlphanumeric),
		Description: "Passwords must have at least one non alphanumeric character.",
	}
}

func (describer *IdentityErrorDescriber) PasswordRequiresDigit() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.PasswordRequiresDigit),
		Description: "Passwords must have at least one digit ('0'-'9').",
	}
}

func (describer *IdentityErrorDescriber) PasswordRequiresLower() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.PasswordRequiresLower),
		Description: "Passwords must have at least one lowercase ('a'-'z').",
	}
}

func (describer *IdentityErrorDescriber) PasswordRequiresUpper() models.IdentityError {
	return models.IdentityError{
		Code:        helpers.GetFunctionName(IIdentityErrorDescriber.PasswordRequiresUpper),
		Description: "Passwords must have at least one uppercase ('A'-'Z').",
	}
}
