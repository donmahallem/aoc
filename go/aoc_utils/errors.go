package aoc_utils

import (
	"errors"
	"fmt"
)

// ErrParsing can be used with errors.Is to check for parse errors
var ErrParsing = errors.New("parsing error")

// ParseError represents an error that occurred while parsing input for a part
type ParseError struct {
	Msg string
	Err error
}

// NewParseError creates a new ParseError
func NewParseError(msg string, err error) error {
	return &ParseError{Msg: msg, Err: err}
}

func NewUnexpectedInputError(ch byte) error {
	return NewParseError(fmt.Sprintf("unexpected input: %q", ch), nil)
}

func (e ParseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("parsing error: %s: %v", e.Msg, e.Err)
	}
	return fmt.Sprintf("parsing error: %s", e.Msg)
}

func (e ParseError) Unwrap() error { return e.Err }

// Is allows errors.Is to match ErrParsing
func (e ParseError) Is(target error) bool {
	return target == ErrParsing || errors.Is(e.Err, target)
}

type NotImplementedError struct {
	Msg string
}

func (e NotImplementedError) Error() string {
	return fmt.Sprintf("not implemented: %s", e.Msg)
}

func (e NotImplementedError) Is(target error) bool {
	_, ok := target.(*NotImplementedError)
	return ok
}

func (e NotImplementedError) Unwrap() error {
	return nil
}

func NewNotImplementedError(msg string) error {
	return &NotImplementedError{Msg: msg}
}

// TimeoutError is returned by the solver when a part takes too long to
// complete.
type TimeoutError struct {
	Msg string
}

func (e TimeoutError) Error() string {
	if e.Msg == "" {
		return "timeout"
	}
	return fmt.Sprintf("timeout: %s", e.Msg)
}

func (e TimeoutError) Is(target error) bool {
	_, ok := target.(*TimeoutError)
	return ok
}

func (e TimeoutError) Unwrap() error { return nil }

func NewTimeoutError(msg string) error {
	return &TimeoutError{Msg: msg}
}

type SolverError struct {
	Message string
	Err     error
}

func (e SolverError) Error() string {
	return fmt.Sprintf("solver error: %s: %v", e.Message, e.Err)
}

func (e SolverError) Is(target error) bool {
	_, ok := target.(*SolverError)
	return ok
}

func (e SolverError) Unwrap() error { return e.Err }

func NewSolverError(message string, err error) error {
	return &SolverError{Message: message, Err: err}
}

func NewSolverErrorf(format string, args ...interface{}) error {
	return &SolverError{Message: fmt.Sprintf(format, args...), Err: nil}
}

func WrapSolverError(err error) error {
	return &SolverError{Message: err.Error(), Err: err}
}
