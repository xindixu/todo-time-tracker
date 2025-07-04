package goutils

// ContextKey is a type for context keys
type ContextKey string

// ContextKeys
const (
	ContextKeyAccountID   ContextKey = "account_id"
	ContextKeyUserID      ContextKey = "user_id"
	ContextKeyUserName    ContextKey = "user_name"
	ContextKeyToken       ContextKey = "token"
	ContextKeyRequestTime ContextKey = "request_time"
)
