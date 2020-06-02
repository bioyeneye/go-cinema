package models

import (
	"fmt"
)

type IdentityResult struct {
	Succeeded		bool
	Errors			[]IdentityError

}

type IdentityError struct {
	/// <summary>
	/// Gets or sets the code for this error.
	/// </summary>
	/// <value>
	/// The code for this error.
	/// </value>
	Code string

	/// <summary>
	/// Gets or sets the description for this error.
	/// </summary>
	/// <value>
	/// The description for this error.
	/// </value>
	Description string
}

func (result *IdentityResult) Failed(errors []IdentityError) IdentityResult  {
	return IdentityResult{
		Succeeded: false,
		Errors:    errors,
	}
}

func (result *IdentityResult) Success() IdentityResult  {
	return IdentityResult{
		Succeeded: true,
		Errors:    nil,
	}
}

func (result *IdentityResult) ToString() string{

	if result.Succeeded {
		return "Succeeded"
	}

	errorcode := ""
	for id, er := range result.Errors {
		if id == 1 {
			errorcode += er.Code
		}else{
			errorcode += fmt.Sprintf(",%s", er.Code)
		}
	}

	return fmt.Sprintf("Failed : %s", errorcode)
}
