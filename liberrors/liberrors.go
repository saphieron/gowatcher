package liberrors

import "fmt"

type InvalidArgsErrors struct {
	reason string
}

func NewInvalidArgsErrors(reason string) error {
	return &InvalidArgsErrors{reason: reason}
}

func (err *InvalidArgsErrors) Error() string {
	return fmt.Sprintf("invalid argument provided: %s", err.reason)
}
