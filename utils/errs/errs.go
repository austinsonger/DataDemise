package errs

import (
	"fmt"
	"runtime"
	"strings"
)

// StackTraceError includes the original error and its stack trace
type StackTraceError struct {
	Err        error
	StackTrace string
}

// Error implements the error interface for StackTraceError
func (e StackTraceError) Error() string {
	return fmt.Sprintf("%s: %v", e.StackTrace, e.Err)
}

// Wrap enriches the given error with a stack trace. It can take either an error or a string.
func Wrap(any interface{}, a ...interface{}) error {
	if any == nil {
		return nil
	}

	var err error
	switch v := any.(type) {
	case string:
		err = fmt.Errorf(v, a...)
	case error:
		err = v
	default:
		err = fmt.Errorf("%v", v)
	}

	_, fn, line, _ := runtime.Caller(1)
	stackTrace := fmt.Sprintf("%s:%d", fn, line)

	return StackTraceError{Err: err, StackTrace: stackTrace}
}

// Errorf creates a new error with a stack trace.
func Errorf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	_, fn, line, _ := runtime.Caller(1)

	return StackTraceError{Err: err, StackTrace: fmt.Sprintf("%s:%d", fn, line)}
}

// ParseError extracts the error code and message, assuming a specific format.
func ParseError(err error) (code, message string) {
	splits := strings.SplitN(err.Error(), ": ", 2)
	if len(splits) != 2 {
		return "GeneralError", err.Error()
	}

	return strings.TrimSpace(splits[0]), strings.TrimSpace(splits[1])
}

// Attach additional helper functions or custom error types as needed.
