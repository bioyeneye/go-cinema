package models

/// Options for user validation.
type UserOptions struct {
	/// Gets or sets the list of allowed characters in the username used to validate user names. Defaults to abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._@+
	/// The list of allowed characters in the username used to validate user names.
	AllowedUserNameCharacters string

	/// <summary>
	/// Gets or sets a flag indicating whether the application requires unique emails for its users. Defaults to false.
	/// </summary>
	/// <value>
	/// True if the application requires each user to have their own, unique email, otherwise false.
	/// </value>
	RequireUniqueEmail bool
}

// create default user options
func 	NewDefaultUserOptions() UserOptions  {
	return UserOptions{
		AllowedUserNameCharacters: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._@+",
		RequireUniqueEmail:        true,
	}
}
