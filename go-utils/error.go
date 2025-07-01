package goutils

import (
	"fmt"
	"log"
)

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	wrappedErr := fmt.Errorf("%s: %w", message, err)
	log.Println(wrappedErr)
	return wrappedErr
}

func WrapAsStr(err error, message string) string {
	return Wrap(err, message).Error()
}
