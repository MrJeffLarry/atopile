package errors

import (
	"fmt"
)

// AtoError is the base error type for atopile errors
type AtoError struct {
	Message string
	Cause   error
	Extra   map[string]interface{}
}

func (e *AtoError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *AtoError) Unwrap() error {
	return e.Cause
}

// UserException represents an error caused by user input or action
type UserException struct {
	AtoError
}

// NewUserException creates a new user exception
func NewUserException(message string) *UserException {
	return &UserException{
		AtoError: AtoError{
			Message: message,
			Extra:   make(map[string]interface{}),
		},
	}
}

// UserBadParameterError represents an error in user-provided parameters
type UserBadParameterError struct {
	UserException
	Parameter string
}

// NewUserBadParameterError creates a new bad parameter error
func NewUserBadParameterError(message string, parameter string) *UserBadParameterError {
	return &UserBadParameterError{
		UserException: UserException{
			AtoError: AtoError{
				Message: message,
				Extra:   make(map[string]interface{}),
			},
		},
		Parameter: parameter,
	}
}

// UserNoProjectException represents an error when no project is found
type UserNoProjectException struct {
	UserException
}

// NewUserNoProjectException creates a new no project exception
func NewUserNoProjectException(message string) *UserNoProjectException {
	return &UserNoProjectException{
		UserException: UserException{
			AtoError: AtoError{
				Message: message,
				Extra:   make(map[string]interface{}),
			},
		},
	}
}

// UserResourceException represents an error with a user resource
type UserResourceException struct {
	UserException
	Resource string
}

// NewUserResourceException creates a new resource exception
func NewUserResourceException(message string) *UserResourceException {
	return &UserResourceException{
		UserException: UserException{
			AtoError: AtoError{
				Message: message,
				Extra:   make(map[string]interface{}),
			},
		},
	}
}

// InternalException represents an internal error (bug in atopile)
type InternalException struct {
	AtoError
}

// NewInternalException creates a new internal exception
func NewInternalException(message string) *InternalException {
	return &InternalException{
		AtoError: AtoError{
			Message: message,
			Extra:   make(map[string]interface{}),
		},
	}
}
