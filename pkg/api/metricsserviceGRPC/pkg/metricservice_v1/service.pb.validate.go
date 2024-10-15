// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/metricservice_v1/service.proto

package metricservice_v1

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

// Validate checks the field values on Metric with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Metric) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Metric with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MetricMultiError, or nil if none found.
func (m *Metric) ValidateAll() error {
	return m.validate(true)
}

func (m *Metric) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetMType()) < 1 {
		err := MetricValidationError{
			field:  "MType",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Delta

	// no validation rules for Value

	if len(errors) > 0 {
		return MetricMultiError(errors)
	}

	return nil
}

// MetricMultiError is an error wrapping multiple validation errors returned by
// Metric.ValidateAll() if the designated constraints aren't met.
type MetricMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricMultiError) AllErrors() []error { return m }

// MetricValidationError is the validation error returned by Metric.Validate if
// the designated constraints aren't met.
type MetricValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricValidationError) ErrorName() string { return "MetricValidationError" }

// Error satisfies the builtin error interface
func (e MetricValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetric.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricValidationError{}

// Validate checks the field values on MetricRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MetricRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetricRequestMultiError, or
// nil if none found.
func (m *MetricRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetricRequestValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetricRequestValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetricRequestValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MetricRequestMultiError(errors)
	}

	return nil
}

// MetricRequestMultiError is an error wrapping multiple validation errors
// returned by MetricRequest.ValidateAll() if the designated constraints
// aren't met.
type MetricRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricRequestMultiError) AllErrors() []error { return m }

// MetricRequestValidationError is the validation error returned by
// MetricRequest.Validate if the designated constraints aren't met.
type MetricRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricRequestValidationError) ErrorName() string { return "MetricRequestValidationError" }

// Error satisfies the builtin error interface
func (e MetricRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricRequestValidationError{}

// Validate checks the field values on MetricResponce with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MetricResponce) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetricResponce with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetricResponceMultiError,
// or nil if none found.
func (m *MetricResponce) ValidateAll() error {
	return m.validate(true)
}

func (m *MetricResponce) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return MetricResponceMultiError(errors)
	}

	return nil
}

// MetricResponceMultiError is an error wrapping multiple validation errors
// returned by MetricResponce.ValidateAll() if the designated constraints
// aren't met.
type MetricResponceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricResponceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricResponceMultiError) AllErrors() []error { return m }

// MetricResponceValidationError is the validation error returned by
// MetricResponce.Validate if the designated constraints aren't met.
type MetricResponceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricResponceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricResponceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricResponceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricResponceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricResponceValidationError) ErrorName() string { return "MetricResponceValidationError" }

// Error satisfies the builtin error interface
func (e MetricResponceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricResponce.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricResponceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricResponceValidationError{}
