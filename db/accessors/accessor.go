package accessors

import (
	"todo-time-tracker/db"
)

// DBAccessor is the main database accessor that implements all model accessors
type DBAccessor struct {
	*db.Database
}

func NewDBAccessor(database *db.Database) *DBAccessor {
	return &DBAccessor{
		Database: database,
	}
}
