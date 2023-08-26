// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mgmt/v1alpha1/transformer.proto

package mgmtv1alpha1

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

// Validate checks the field values on Transformer with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Transformer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Transformer with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TransformerMultiError, or
// nil if none found.
func (m *Transformer) ValidateAll() error {
	return m.validate(true)
}

func (m *Transformer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Value

	if len(errors) > 0 {
		return TransformerMultiError(errors)
	}

	return nil
}

// TransformerMultiError is an error wrapping multiple validation errors
// returned by Transformer.ValidateAll() if the designated constraints aren't met.
type TransformerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TransformerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TransformerMultiError) AllErrors() []error { return m }

// TransformerValidationError is the validation error returned by
// Transformer.Validate if the designated constraints aren't met.
type TransformerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransformerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransformerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransformerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransformerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransformerValidationError) ErrorName() string { return "TransformerValidationError" }

// Error satisfies the builtin error interface
func (e TransformerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransformer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransformerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransformerValidationError{}

// Validate checks the field values on GetTransformersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTransformersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTransformersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTransformersRequestMultiError, or nil if none found.
func (m *GetTransformersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTransformersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetTransformersRequestMultiError(errors)
	}

	return nil
}

// GetTransformersRequestMultiError is an error wrapping multiple validation
// errors returned by GetTransformersRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTransformersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTransformersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTransformersRequestMultiError) AllErrors() []error { return m }

// GetTransformersRequestValidationError is the validation error returned by
// GetTransformersRequest.Validate if the designated constraints aren't met.
type GetTransformersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransformersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransformersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransformersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransformersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransformersRequestValidationError) ErrorName() string {
	return "GetTransformersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransformersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransformersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransformersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransformersRequestValidationError{}

// Validate checks the field values on GetTransformersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTransformersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTransformersResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTransformersResponseMultiError, or nil if none found.
func (m *GetTransformersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTransformersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTransformers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTransformersResponseValidationError{
						field:  fmt.Sprintf("Transformers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTransformersResponseValidationError{
						field:  fmt.Sprintf("Transformers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTransformersResponseValidationError{
					field:  fmt.Sprintf("Transformers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTransformersResponseMultiError(errors)
	}

	return nil
}

// GetTransformersResponseMultiError is an error wrapping multiple validation
// errors returned by GetTransformersResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTransformersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTransformersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTransformersResponseMultiError) AllErrors() []error { return m }

// GetTransformersResponseValidationError is the validation error returned by
// GetTransformersResponse.Validate if the designated constraints aren't met.
type GetTransformersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransformersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransformersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransformersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransformersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransformersResponseValidationError) ErrorName() string {
	return "GetTransformersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransformersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransformersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransformersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransformersResponseValidationError{}
