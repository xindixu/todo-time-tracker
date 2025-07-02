package accessors

import (
	"context"

	"todo-time-tracker/db"

	"github.com/doug-martin/goqu/v9"
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

func (a *DBAccessor) CountModels(ctx context.Context, tableName string) (int64, error) {
	query, args, err := a.Builder.Select(goqu.COUNT("*")).From(tableName).ToSQL()
	if err != nil {
		return 0, err
	}

	var count int64
	err = a.DB.GetContext(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}

	return count, nil
}
