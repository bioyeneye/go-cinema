package base

import "github.com/bioyeneye/rest-gin-api/core/identity/models"

type UserRoleStoreBase interface {
	CreateUserRole(user models.IdentityUser, role models.IdentityRole) models.IdentityUserRole
	GetUsersInRole(normalizedRoleName string) []models.IdentityUser
	AddToRole(user models.IdentityUser, normalizedRoleName string)
	RemoveFromRole(user models.IdentityUser, normalizedRoleName string)
	GetRoles(user models.IdentityUser) []string
	IsInRole(user models.IdentityUser, normalizedRoleName string) bool
	FindRole(normalizedRoleName string) models.IdentityRole
	FindUserRoleAsync(userId string, roleId string) models.IdentityRole
}
