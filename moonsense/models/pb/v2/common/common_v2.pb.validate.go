// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common_v2.proto

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on ErrorResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ErrorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ErrorResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ErrorResponseMultiError, or
// nil if none found.
func (m *ErrorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ErrorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Param

	// no validation rules for Message

	if len(errors) > 0 {
		return ErrorResponseMultiError(errors)
	}

	return nil
}

// ErrorResponseMultiError is an error wrapping multiple validation errors
// returned by ErrorResponse.ValidateAll() if the designated constraints
// aren't met.
type ErrorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ErrorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ErrorResponseMultiError) AllErrors() []error { return m }

// ErrorResponseValidationError is the validation error returned by
// ErrorResponse.Validate if the designated constraints aren't met.
type ErrorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorResponseValidationError) ErrorName() string { return "ErrorResponseValidationError" }

// Error satisfies the builtin error interface
func (e ErrorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorResponseValidationError{}

// Validate checks the field values on TokenSelfResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TokenSelfResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TokenSelfResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TokenSelfResponseMultiError, or nil if none found.
func (m *TokenSelfResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TokenSelfResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ProjectId

	// no validation rules for Scopes

	// no validation rules for UserId

	// no validation rules for CredentialId

	// no validation rules for Type

	if len(errors) > 0 {
		return TokenSelfResponseMultiError(errors)
	}

	return nil
}

// TokenSelfResponseMultiError is an error wrapping multiple validation errors
// returned by TokenSelfResponse.ValidateAll() if the designated constraints
// aren't met.
type TokenSelfResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokenSelfResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokenSelfResponseMultiError) AllErrors() []error { return m }

// TokenSelfResponseValidationError is the validation error returned by
// TokenSelfResponse.Validate if the designated constraints aren't met.
type TokenSelfResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenSelfResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenSelfResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenSelfResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenSelfResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenSelfResponseValidationError) ErrorName() string {
	return "TokenSelfResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TokenSelfResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenSelfResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenSelfResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenSelfResponseValidationError{}
