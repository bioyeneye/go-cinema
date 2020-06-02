package models

type IdentityOptions struct {
	User UserOptions
	Password PasswordOptions
	Lockout LockoutOptions
	SignIn SignInOptions
}

func (options *IdentityOptions) NewDefault() IdentityOptions {

	return IdentityOptions{
		User:     NewDefaultUserOptions(),
		Password: NewDefaultPasswordOptions(),
		Lockout:  NewDefaultLockoutOptions(),
		SignIn:   NewDefaultSignInOptions(),
	}
}

