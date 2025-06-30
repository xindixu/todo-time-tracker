package models

import (
	"time"

	"github.com/google/uuid"
)

// This file contains the models for the database

// AccountType represents a account type in the database
type AccountType string

const (
	AccountTypeUser         AccountType = "USER"
	AccountTypeOrganization AccountType = "ORGANIZATION"
)

func (a AccountType) IsValid() bool {
	switch a {
	case AccountTypeUser, AccountTypeOrganization:
		return true
	default:
		return false
	}
}

// Account represents a account in the database
type Account struct {
	ID        int64       `db:"id"`
	UUID      uuid.UUID   `db:"uuid"`
	Type      AccountType `db:"type"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
}

// TableName returns the table name for the Account model
func (Account) TableName() string {
	return "accounts"
}

// User represents a user in the database
type User struct {
	ID        int64     `db:"id" json:"id" goqu:"skipinsert"`
	UUID      uuid.UUID `db:"uuid" json:"uuid"`
	AccountID int64     `db:"account_id" json:"account_id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
}

// TableName returns the table name for the User model
func (User) TableName() string {
	return "users"
}

// Organization represents a organization in the Database
type Organization struct {
	ID        int64     `db:"id" json:"id" goqu:"skipinsert"`
	UUID      uuid.UUID `db:"uuid" json:"uuid"`
	AccountID int64     `db:"account_id" json:"account_id"`
	Name      string    `db:"name" json:"name"`
}

// TableName returns the table name for the Organization model
func (Organization) TableName() string {
	return "organizations"
}

// Tag represents a tag in the database
type Tag struct {
	ID          int64     `db:"id" json:"id" goqu:"skipinsert"`
	UUID        uuid.UUID `db:"uuid" json:"uuid"`
	AccountID   int64     `db:"account_id" json:"account_id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Color       string    `db:"color" json:"color"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// TableName returns the table name for the Tag model
func (Tag) TableName() string {
	return "tags"
}

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "TODO"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusBlocked    TaskStatus = "BLOCKED"
	TaskStatusDone       TaskStatus = "DONE"
)

func (s TaskStatus) IsValid() bool {
	switch s {
	case TaskStatusTodo, TaskStatusInProgress, TaskStatusDone, TaskStatusBlocked:
		return true
	default:
		return false
	}
}

// Task represents a task in the database
type Task struct {
	ID                int64          `db:"id" json:"id" goqu:"skipinsert"`
	UUID              uuid.UUID      `db:"uuid" json:"uuid"`
	AccountID         int64          `db:"account_id" json:"account_id"`
	Name              string         `db:"name" json:"name"`
	Description       string         `db:"description" json:"description"`
	EstimatedDuration *time.Duration `db:"estimated_duration" json:"estimated_duration"` // nullable
	Status            TaskStatus     `db:"status" json:"status"`
	CreatedAt         time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updated_at"`
}

// TableName returns the table name for the Task model
func (Task) TableName() string {
	return "tasks"
}

// Session represents a session in the database
type Session struct {
	ID       int64      `db:"id" json:"id" goqu:"skipinsert"`
	UUID     uuid.UUID  `db:"uuid" json:"uuid"`
	TaskID   int64      `db:"task_id" json:"task_id"`
	UserID   int64      `db:"user_id" json:"user_id"`
	StartsAt time.Time  `db:"starts_at" json:"starts_at"`
	EndsAt   *time.Time `db:"ends_at" json:"ends_at"` // nullable
}

// TableName returns the table name for the Session model
func (Session) TableName() string {
	return "sessions"
}

// TaskTag represents a task tag in the database
type TaskTag struct {
	ID     int64 `db:"id" json:"id" goqu:"skipinsert"`
	TaskID int64 `db:"task_id" json:"task_id"`
	TagID  int64 `db:"tag_id" json:"tag_id"`
}

func (TaskTag) TableName() string {
	return "task_tags"
}

// TaskUser represents a task user in the database
type TaskUser struct {
	ID     int64 `db:"id" json:"id" goqu:"skipinsert"`
	TaskID int64 `db:"task_id" json:"task_id"`
	UserID int64 `db:"user_id" json:"user_id"`
}

func (TaskUser) TableName() string {
	return "task_users"
}
