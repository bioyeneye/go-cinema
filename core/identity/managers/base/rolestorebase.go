package base

import (
	"github.com/bioyeneye/rest-gin-api/core/identity/models"
)

type IdentityResult models.IdentityResult

type RoleStoreBase interface {
	Create(role models.IdentityRole) IdentityResult
	Update(role models.IdentityRole) IdentityResult
	Delete(role models.IdentityRole) IdentityResult
	GetRoleId(role models.IdentityRole) string
	SetRoleName(role models.IdentityRole)
	FindById(id string) models.IdentityRole
	FindByNameAsync(normalizedName string) models.IdentityRole
	SetNormalizedRoleNameAsync(role models.IdentityRole, normalizedName string) models.IdentityRole
}
