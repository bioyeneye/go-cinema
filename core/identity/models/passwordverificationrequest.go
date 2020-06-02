package models

type Enum = int

type list struct {
	Failed Enum
	Success Enum
	SuccessRehashNeeded Enum
}

// Enum for public use
var PasswordVerificationResult = &list{
	Failed:              0,
	Success:             1,
	SuccessRehashNeeded: 3,
}
