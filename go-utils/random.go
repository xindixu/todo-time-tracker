package goutils

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandEmail() string {
	return fmt.Sprintf("test+%s@example.com", RandString(10))
}

func RandDuration(n int) time.Duration {
	return time.Duration(rand.Intn(n)) * time.Minute
}
