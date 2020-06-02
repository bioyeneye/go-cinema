package models

import "time"

// LockoutOptions Options for configuring user lockout.
type LockoutOptions struct {

	/// <summary>
	/// Gets or sets a flag indicating whether a new user can be locked out. Defaults to true.
	/// </summary>
	/// <value>
	/// True if a newly created user can be locked out, otherwise false.
	/// </value>
	AllowedForNewUsers bool

	/// <summary>
	/// Gets or sets the number of failed access attempts allowed before a user is locked out,
	/// assuming lock out is enabled. Defaults to 5.
	/// </summary>
	/// <value>
	/// The number of failed access attempts allowed before a user is locked out, if lockout is enabled.
	/// </value>
	MaxFailedAccessAttempts int

	/// <summary>
	/// Gets or sets the <see cref="time.Time"/> a user is locked out for when a lockout occurs. Defaults to 5 minutes.
	/// </summary>
	/// <value>The <see cref="time.Time"/> a user is locked out for when a lockout occurs.</value>
	DefaultLockoutTimeSpan time.Time
}

func NewDefaultLockoutOptions() LockoutOptions  {
	return LockoutOptions{
		AllowedForNewUsers:      true,
		MaxFailedAccessAttempts: 5,
		DefaultLockoutTimeSpan:  time.Now().Local().Add(time.Minute * time.Duration(5)),
	}
}