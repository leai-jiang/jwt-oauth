package jwt

import (
	"errors"
)

// Error constants
var (
	ErrHashUnavailable = errors.New("the requested hash function is unavailable")
	ErrSigningMethodUnavailable = errors.New("the signing method is unavailable")
)

// The errors that might occur when parsing and validating a token
const (
	ValidationErrorMalformed        uint32 = 1 << iota // 签名不规范
	ValidationErrorUnverifiable                        // Token could not be verified because of signing problems
	ValidationErrorSignatureInvalid                    // Signature validation failed

	// Standard Claim validation errors
	ValidationErrorExpired       // 签名过期
	ValidationErrorNotValidYet   // 签名未生效
	ValidationErrorClaimsInvalid // Generic claims validation error
)

// Helper for constructing a ValidationError with a string error message
func NewValidationError(errorText string, errorFlags uint32) *ValidationError {
	return &ValidationError{
		text:   errorText,
		Errors: errorFlags,
	}
}

// The error from Parse if token is not valid
type ValidationError struct {
	Inner  error  // 存放解析过程之中的错误信息
	Errors uint32 // bitfield.  see ValidationError... constants
	text   string // errors that do not have a valid error just have text
}

// Validation error is an error type
func (e ValidationError) Error() string {
	if e.Inner != nil {
		return e.Inner.Error()
	} else if e.text != "" {
		return e.text
	} else {
		return "token is invalid"
	}
}

// No errors
func (e *ValidationError) valid() bool {
	return e.Errors == 0
}