// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: error_reporting.proto

package error_reporting

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

	v2 "github.com/moonsense/go-sdk/moonsense/models/pb/v2/common"
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

	_ = v2.DevicePlatform(0)
)

// Validate checks the field values on Frame with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Frame) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Frame with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FrameMultiError, or nil if none found.
func (m *Frame) ValidateAll() error {
	return m.validate(true)
}

func (m *Frame) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Function

	// no validation rules for Module

	// no validation rules for Filename

	// no validation rules for LineNo

	// no validation rules for ColNo

	if len(errors) > 0 {
		return FrameMultiError(errors)
	}

	return nil
}

// FrameMultiError is an error wrapping multiple validation errors returned by
// Frame.ValidateAll() if the designated constraints aren't met.
type FrameMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FrameMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FrameMultiError) AllErrors() []error { return m }

// FrameValidationError is the validation error returned by Frame.Validate if
// the designated constraints aren't met.
type FrameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FrameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FrameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FrameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FrameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FrameValidationError) ErrorName() string { return "FrameValidationError" }

// Error satisfies the builtin error interface
func (e FrameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFrame.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FrameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FrameValidationError{}

// Validate checks the field values on Exception with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Exception) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Exception with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExceptionMultiError, or nil
// if none found.
func (m *Exception) ValidateAll() error {
	return m.validate(true)
}

func (m *Exception) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetType()) < 1 {
		err := ExceptionValidationError{
			field:  "Type",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := ExceptionValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetStacktrace() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExceptionValidationError{
						field:  fmt.Sprintf("Stacktrace[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExceptionValidationError{
						field:  fmt.Sprintf("Stacktrace[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExceptionValidationError{
					field:  fmt.Sprintf("Stacktrace[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ExceptionMultiError(errors)
	}

	return nil
}

// ExceptionMultiError is an error wrapping multiple validation errors returned
// by Exception.ValidateAll() if the designated constraints aren't met.
type ExceptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExceptionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExceptionMultiError) AllErrors() []error { return m }

// ExceptionValidationError is the validation error returned by
// Exception.Validate if the designated constraints aren't met.
type ExceptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExceptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExceptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExceptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExceptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExceptionValidationError) ErrorName() string { return "ExceptionValidationError" }

// Error satisfies the builtin error interface
func (e ExceptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sException.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExceptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExceptionValidationError{}

// Validate checks the field values on ErrorMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ErrorMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ErrorMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ErrorMessageMultiError, or
// nil if none found.
func (m *ErrorMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ErrorMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := ErrorMessageValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimestamp() <= 0 {
		err := ErrorMessageValidationError{
			field:  "Timestamp",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _ErrorMessage_Platform_NotInLookup[m.GetPlatform()]; ok {
		err := ErrorMessageValidationError{
			field:  "Platform",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := v2.DevicePlatform_name[int32(m.GetPlatform())]; !ok {
		err := ErrorMessageValidationError{
			field:  "Platform",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Manufacturer

	// no validation rules for Model

	// no validation rules for OsVersion

	if _, ok := _ErrorMessage_Level_NotInLookup[m.GetLevel()]; ok {
		err := ErrorMessageValidationError{
			field:  "Level",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := ErrorMessage_Level_name[int32(m.GetLevel())]; !ok {
		err := ErrorMessageValidationError{
			field:  "Level",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ReleaseVersion

	// no validation rules for ProguardUuid

	if all {
		switch v := interface{}(m.GetException()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ErrorMessageValidationError{
					field:  "Exception",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ErrorMessageValidationError{
					field:  "Exception",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetException()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ErrorMessageValidationError{
				field:  "Exception",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AdditionalTags

	if len(errors) > 0 {
		return ErrorMessageMultiError(errors)
	}

	return nil
}

// ErrorMessageMultiError is an error wrapping multiple validation errors
// returned by ErrorMessage.ValidateAll() if the designated constraints aren't met.
type ErrorMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ErrorMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ErrorMessageMultiError) AllErrors() []error { return m }

// ErrorMessageValidationError is the validation error returned by
// ErrorMessage.Validate if the designated constraints aren't met.
type ErrorMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorMessageValidationError) ErrorName() string { return "ErrorMessageValidationError" }

// Error satisfies the builtin error interface
func (e ErrorMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorMessageValidationError{}

var _ErrorMessage_Platform_NotInLookup = map[v2.DevicePlatform]struct{}{
	0: {},
}

var _ErrorMessage_Level_NotInLookup = map[ErrorMessage_Level]struct{}{
	0: {},
}