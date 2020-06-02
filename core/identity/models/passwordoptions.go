package models

/// Specifies options for password requirements.
type PasswordOptions struct {

	/// Gets or sets the minimum length a password must be. Defaults to 6.
	RequiredLength int

	/// Gets or sets the minimum number of unique chars a password must comprised of. Defaults to 1.
	RequiredUniqueChars int

	/// Gets or sets a flag indicating if passwords must contain a non-alphanumeric character. Defaults to true.
	/// <value>True if passwords must contain a non-alphanumeric character, otherwise false.</value>
	RequireNonAlphanumeric bool

	/// Gets or sets a flag indicating if passwords must contain a lower case ASCII character. Defaults to true.
	/// <value>True if passwords must contain a lower case ASCII character.</value>
	RequireLowercase bool

	/// <summary>
	/// Gets or sets a flag indicating if passwords must contain a upper case ASCII character. Defaults to true.
	/// </summary>
	/// <value>True if passwords must contain a upper case ASCII character.</value>
	RequireUppercase bool

	/// <summary>
	/// Gets or sets a flag indicating if passwords must contain a digit. Defaults to true.
	/// </summary>
	/// <value>True if passwords must contain a digit.</value>
	RequireDigit bool
}

func NewDefaultPasswordOptions() PasswordOptions  {
	return PasswordOptions{
		RequiredLength:         6,
		RequiredUniqueChars:    1,
		RequireNonAlphanumeric: true,
		RequireLowercase:       true,
		RequireUppercase:       true,
		RequireDigit:           true,
	}
}

