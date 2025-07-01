package goutils

type ContextKey string

const (
	ContextKeyAccountID   ContextKey = "account_id"
	ContextKeyUserID      ContextKey = "user_id"
	ContextKeyToken       ContextKey = "token"
	ContextKeyRequestTime ContextKey = "request_time"
)
