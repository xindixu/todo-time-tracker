package models

import (
	"time"
)

// Tag represents a tag in the database
type Tag struct {
	ID        int64     `db:"id" json:"id"`
	UUID      string    `db:"uuid" json:"uuid"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// TableName returns the table name for the Tag model
func (Tag) TableName() string {
	return "tags"
}
