// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: journey_feedback.proto

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

// Validate checks the field values on FraudFeedback with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FraudFeedback) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FraudFeedback with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FraudFeedbackMultiError, or
// nil if none found.
func (m *FraudFeedback) ValidateAll() error {
	return m.validate(true)
}

func (m *FraudFeedback) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FeedbackId

	// no validation rules for IsFraud

	if all {
		switch v := interface{}(m.GetReportedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FraudFeedbackValidationError{
					field:  "ReportedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FraudFeedbackValidationError{
					field:  "ReportedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReportedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FraudFeedbackValidationError{
				field:  "ReportedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FraudReason

	if len(errors) > 0 {
		return FraudFeedbackMultiError(errors)
	}

	return nil
}

// FraudFeedbackMultiError is an error wrapping multiple validation errors
// returned by FraudFeedback.ValidateAll() if the designated constraints
// aren't met.
type FraudFeedbackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FraudFeedbackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FraudFeedbackMultiError) AllErrors() []error { return m }

// FraudFeedbackValidationError is the validation error returned by
// FraudFeedback.Validate if the designated constraints aren't met.
type FraudFeedbackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FraudFeedbackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FraudFeedbackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FraudFeedbackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FraudFeedbackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FraudFeedbackValidationError) ErrorName() string { return "FraudFeedbackValidationError" }

// Error satisfies the builtin error interface
func (e FraudFeedbackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFraudFeedback.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FraudFeedbackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FraudFeedbackValidationError{}

// Validate checks the field values on ChargebackFeedback with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChargebackFeedback) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChargebackFeedback with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChargebackFeedbackMultiError, or nil if none found.
func (m *ChargebackFeedback) ValidateAll() error {
	return m.validate(true)
}

func (m *ChargebackFeedback) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FeedbackId

	// no validation rules for IsChargeback

	if all {
		switch v := interface{}(m.GetReportedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChargebackFeedbackValidationError{
					field:  "ReportedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChargebackFeedbackValidationError{
					field:  "ReportedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReportedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChargebackFeedbackValidationError{
				field:  "ReportedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ChargebackReason

	if len(errors) > 0 {
		return ChargebackFeedbackMultiError(errors)
	}

	return nil
}

// ChargebackFeedbackMultiError is an error wrapping multiple validation errors
// returned by ChargebackFeedback.ValidateAll() if the designated constraints
// aren't met.
type ChargebackFeedbackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChargebackFeedbackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChargebackFeedbackMultiError) AllErrors() []error { return m }

// ChargebackFeedbackValidationError is the validation error returned by
// ChargebackFeedback.Validate if the designated constraints aren't met.
type ChargebackFeedbackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChargebackFeedbackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChargebackFeedbackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChargebackFeedbackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChargebackFeedbackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChargebackFeedbackValidationError) ErrorName() string {
	return "ChargebackFeedbackValidationError"
}

// Error satisfies the builtin error interface
func (e ChargebackFeedbackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChargebackFeedback.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChargebackFeedbackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChargebackFeedbackValidationError{}

// Validate checks the field values on JourneyFeedback with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *JourneyFeedback) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JourneyFeedback with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// JourneyFeedbackMultiError, or nil if none found.
func (m *JourneyFeedback) ValidateAll() error {
	return m.validate(true)
}

func (m *JourneyFeedback) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for JourneyId

	if all {
		switch v := interface{}(m.GetFraudFeedback()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, JourneyFeedbackValidationError{
					field:  "FraudFeedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, JourneyFeedbackValidationError{
					field:  "FraudFeedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFraudFeedback()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JourneyFeedbackValidationError{
				field:  "FraudFeedback",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetChargebackFeedback()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, JourneyFeedbackValidationError{
					field:  "ChargebackFeedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, JourneyFeedbackValidationError{
					field:  "ChargebackFeedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetChargebackFeedback()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return JourneyFeedbackValidationError{
				field:  "ChargebackFeedback",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return JourneyFeedbackMultiError(errors)
	}

	return nil
}

// JourneyFeedbackMultiError is an error wrapping multiple validation errors
// returned by JourneyFeedback.ValidateAll() if the designated constraints
// aren't met.
type JourneyFeedbackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JourneyFeedbackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JourneyFeedbackMultiError) AllErrors() []error { return m }

// JourneyFeedbackValidationError is the validation error returned by
// JourneyFeedback.Validate if the designated constraints aren't met.
type JourneyFeedbackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JourneyFeedbackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JourneyFeedbackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JourneyFeedbackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JourneyFeedbackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JourneyFeedbackValidationError) ErrorName() string { return "JourneyFeedbackValidationError" }

// Error satisfies the builtin error interface
func (e JourneyFeedbackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJourneyFeedback.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JourneyFeedbackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JourneyFeedbackValidationError{}
