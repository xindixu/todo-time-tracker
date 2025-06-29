package accessors

import (
	"todo-time-tracker/db"
)

// DBAccessor is the main database accessor that implements all model accessors
type DBAccessor struct {
	*db.DBConnection
}

func NewDBAccessor(dbConnection *db.DBConnection) *DBAccessor {
	return &DBAccessor{
		DBConnection: dbConnection,
	}
}
