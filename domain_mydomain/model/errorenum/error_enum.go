package errorenum

import "your/path/project/shared/model/apperror"

const (
	SomethingError         apperror.ErrorType = "ER0000 something error"
	NameMustNotEmpty       apperror.ErrorType = "ER0001 name must not empty"
	AgeMustGreaterThanZero apperror.ErrorType = "ER0002 age must greater than zero"
)
