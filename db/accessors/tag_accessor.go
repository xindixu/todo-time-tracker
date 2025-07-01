package accessors

import (
	"context"

	"todo-time-tracker/db/models"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

// TagAccessor handles database operations for tags
type TagAccessor interface {
	CreateTag(ctx context.Context, uuid uuid.UUID, name string) (*models.Tag, error)
	GetTagByUUID(ctx context.Context, uuid string) (*models.Tag, error)
}

// Ensure DBAccessor implements TagAccessor
var _ TagAccessor = (*DBAccessor)(nil)

// CreateTag creates a new tag in the database
func (a *DBAccessor) CreateTag(ctx context.Context, uuid uuid.UUID, name string) (*models.Tag, error) {
	tagsTable := goqu.T("tags")

	tag := &models.Tag{
		UUID: uuid,
		Name: name,
	}

	// Build insert query with RETURNING clause to get the inserted ID
	insert := a.Builder.Insert(tagsTable).
		Rows(tag).
		Returning("id")

	query, args, err := insert.ToSQL()
	if err != nil {
		return nil, err
	}

	var insertedID int64
	err = a.DB.QueryRowContext(ctx, query, args...).Scan(&insertedID)
	if err != nil {
		return nil, err
	}

	// Set the ID on the tag and return it
	tag.ID = insertedID
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

	err = a.DB.QueryRowxContext(ctx, query, args...).StructScan(&tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
