package liberrors

import (
	"fmt"
	"strings"
)

// ////////////////
// InvalidArgsError
// ////////////////

type InvalidArgsError struct {
	reason string
}

func NewInvalidArgsError(reason string) error {
	return &InvalidArgsError{reason: reason}
}

func (err *InvalidArgsError) Error() string {
	return fmt.Sprintf("invalid argument provided: %s", err.reason)
}

// ////////////////
// CollectedError
// ////////////////

type CollectedError struct {
	errors  []error
	context string
}

func NewCollectedError(context string) *CollectedError {
	return &CollectedError{context: context}
}

func (err *CollectedError) AppendError(newError error) {
	err.errors = append(err.errors, newError)
}

func (err *CollectedError) Error() string {
	var builder strings.Builder
	builder.WriteString("errors collected in '")
	builder.WriteString(err.context)
	builder.WriteString("': ")
	errorMessages := make([]string, 0, len(err.errors))
	for _, elem := range err.errors {
		errorMessages = append(errorMessages, elem.Error())
	}
	builder.WriteString(strings.Join(errorMessages, ", "))
	return builder.String()
}
