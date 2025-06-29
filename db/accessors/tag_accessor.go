package accessors

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"

	"todo-time-tracker/db"
	"todo-time-tracker/db/models"
)

// TagAccessor handles database operations for tags
type TagAccessor interface {
	CreateTag(ctx context.Context, uuid, name string) (*models.Tag, error)
	GetTagByUUID(ctx context.Context, uuid string) (*models.Tag, error)
}

// Ensure DBAccessor implements TagAccessor
var _ TagAccessor = (*DBAccessor)(nil)

// CreateTag creates a new tag in the database
func (a *DBAccessor) CreateTag(ctx context.Context, uuid, name string) (*models.Tag, error) {
	now := time.Now()
	tagsTable := goqu.T("tags")

	// Create the tag model
	tag := &models.Tag{
		UUID:      uuid,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Convert model to record using the generic function
	record := db.ModelToRecord(tag)

	// Insert the tag
	insert := a.Builder.Insert(tagsTable).Rows(record)

	query, args, err := insert.ToSQL()
	if err != nil {
		return nil, err
	}

	result, err := a.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set the ID on the tag and return it
	tag.ID = id
	return tag, nil
}

// GetTagByUUID retrieves a tag by its UUID
func (a *DBAccessor) GetTagByUUID(ctx context.Context, uuid string) (*models.Tag, error) {
	var tag models.Tag

	tagsTable := goqu.T("tags")

	query, args, err := a.Builder.From(tagsTable).
		Where(tagsTable.Col("uuid").Eq(uuid)).
		ToSQL()
	if err != nil {
		return nil, err
	}

	err = a.DB.QueryRowContext(ctx, query, args...).Scan(
		&tag.ID, &tag.UUID, &tag.Name, &tag.CreatedAt, &tag.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
