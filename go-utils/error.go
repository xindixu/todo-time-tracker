package goutils

import (
	"fmt"
	"log"
)

// Wrap wraps an error with a message
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	wrappedErr := fmt.Errorf("%s: %w", message, err)
	log.Println(wrappedErr)
	return wrappedErr
}

// WrapAsStr wraps an error with a message and returns a string
func WrapAsStr(err error, message string) string {
	return Wrap(err, message).Error()
}
