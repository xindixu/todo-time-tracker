package accessors

import (
	"database/sql"

	"todo-time-tracker/db"

	"github.com/doug-martin/goqu/v9"
)

// DBAccessor is the main database accessor that implements all model accessors
type DBAccessor struct {
	DB      *sql.DB
	Builder *goqu.Database
}

func NewDBAccessor(dbConnection *db.DBConnection) *DBAccessor {
	return &DBAccessor{
		DB:      dbConnection.DB,
		Builder: dbConnection.Builder,
	}
}
