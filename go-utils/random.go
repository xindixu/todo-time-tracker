package goutils

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandString generates a random string
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandEmail generates a random email
func RandEmail() string {
	return fmt.Sprintf("test+%s@example.com", RandString(10))
}

// RandDuration generates a random duration
func RandDuration(n int) time.Duration {
	return time.Duration(rand.Intn(n)) * time.Minute
}
