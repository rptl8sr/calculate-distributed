package ownErrors

import "fmt"

var (
	ErrExpressionNotFound = fmt.Errorf("expression not found")
	ErrExpressionExists   = fmt.Errorf("expression already exists")
	ErrExpressionInvalid  = fmt.Errorf("expression is invalid")
	ErrIDNotFound         = fmt.Errorf("expression id not found")
	ErrIDInvalid          = fmt.Errorf("expression id is invalid")
	ErrTaskNotFound       = fmt.Errorf("task not found")
	ErrTaskExists         = fmt.Errorf("task already exists")
)
