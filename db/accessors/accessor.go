package accessors

import (
	"context"

	"todo-time-tracker/db"

	"github.com/doug-martin/goqu/v9"
)

// DBAccessor is the main database accessor that implements all model accessors
type DBAccessor struct {
	db.Connection
}

// NewDBAccessor creates a new DBAccessor
func NewDBAccessor(dbConnection *db.Connection) *DBAccessor {
	return &DBAccessor{
		*dbConnection,
	}
}

// CountModels counts the number of models in the database
func (a *DBAccessor) CountModels(ctx context.Context, tableName string) (int64, error) {
	query, args, err := a.SQLBuilder.Select(goqu.COUNT("*")).From(tableName).ToSQL()
	if err != nil {
		return 0, err
	}

	var count int64
	err = a.SQLDB.GetContext(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}
