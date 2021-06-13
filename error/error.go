package err

import (
	"errors"
	"strings"
)

type ParentError struct{}

func (e ParentError) Error() string {
	return "Error parentError"
}

type Child1Error struct{}

func (e Child1Error) Error() string {
	return "Error major child1"
}

type Child2Error struct{}

func (e Child2Error) Error() string {
	return "Error minor child2"
}

func CheckErrorLevel(err error) error {
	if unWrapedErr := errors.Unwrap(err); unWrapedErr != nil {
		if strings.Contains(unWrapedErr.Error(), "minor") {
			return nil
		} else {
			return err
		}
	}
	return nil
}
