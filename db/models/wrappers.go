package models

import (
	"time"

	"github.com/google/uuid"
)

type UserAccountWrapperFlat struct {
	UserID           int64       `db:"user_id" json:"user_id"`
	UserUUID         uuid.UUID   `db:"user_uuid" json:"user_uuid"`
	UserCreatedAt    time.Time   `db:"user_created_at" json:"user_created_at"`
	UserUpdatedAt    time.Time   `db:"user_updated_at" json:"user_updated_at"`
	UserAccountID    int64       `db:"user_account_id" json:"user_account_id"`
	UserPassword     string      `db:"user_password" json:"user_password"`
	UserName         string      `db:"user_name" json:"user_name"`
	UserEmail        string      `db:"user_email" json:"user_email"`
	AccountID        int64       `db:"account_id" json:"account_id"`
	AccountUUID      uuid.UUID   `db:"account_uuid" json:"account_uuid"`
	AccountCreatedAt time.Time   `db:"account_created_at" json:"account_created_at"`
	AccountUpdatedAt time.Time   `db:"account_updated_at" json:"account_updated_at"`
	AccountType      AccountType `db:"account_type" json:"account_type"`
}

func (u *UserAccountWrapperFlat) ToUserAccountWrapper() *UserAccountWrapper {
	return &UserAccountWrapper{
		User: &User{
			ID:        u.UserID,
			UUID:      u.UserUUID,
			CreatedAt: u.UserCreatedAt,
			UpdatedAt: u.UserUpdatedAt,
			AccountID: u.UserAccountID,
			Name:      u.UserName,
			Email:     u.UserEmail,
			Password:  u.UserPassword,
		},
		Account: &Account{
			ID:        u.AccountID,
			UUID:      u.AccountUUID,
			CreatedAt: u.AccountCreatedAt,
			UpdatedAt: u.AccountUpdatedAt,
			Type:      u.AccountType,
		},
	}
}

type UserAccountWrapper struct {
	User    *User    `db:"user" json:"user"`
	Account *Account `db:"account" json:"account"`
}

type TaskWrapper struct {
	Task *Task
	Tags []*Tag
}
