package identity

import (
	"github.com/bioyeneye/rest-gin-api/core/identity/models"
	"github.com/bioyeneye/rest-gin-api/core/utilities"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Provides an abstraction for hashing passwords.
type IPasswordHasher interface {

	// Returns a hashed representation of the supplied <paramref name="password"/> for the specified <paramref name="user"/>.
	//	 <param name="user">The user whose password is to be hashed.</param>
	// 	 <param name="password">The password to hash.</param>
	// 	 <returns>A hashed representation of the supplied <paramref name="password"/> for the specified <paramref name="user"/>.</returns>
	HashPassword(user models.IdentityUser, password string) string

	// Returns a <see cref="PasswordVerificationResult"/> indicating the result of a password hash comparison.
	// 	<param name="user">The user whose password should be verified.</param>
	//	<param name="hashedPassword">The hash value for a user's stored password.</param>
	//	<param name="providedPassword">The password supplied for comparison.</param>
	// 	<returns>A <see cref="PasswordVerificationResult"/> indicating the result of a password hash comparison.</returns>
	// 	<remarks>Implementations of this method should be time consistent.</remarks>
	VerifyHashedPassword(user models.IdentityUser, hashedPassword string, providedPassword string) bool
}

type PasswordHasher struct {

}

func (p PasswordHasher) HashPassword(user models.IdentityUser, password string) string {
	hash, err := utilities.Hash(password, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func (p PasswordHasher) VerifyHashedPassword(user models.IdentityUser, hashedPassword string, providedPassword string) bool {
	return utilities.VerifyHash(hashedPassword, providedPassword)
}

func NewPasswordHasher() IPasswordHasher {
	return &PasswordHasher{}
}
