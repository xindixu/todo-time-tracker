package accessors

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"

	"todo-time-tracker/db/models"
)

// TagAccessor handles database operations for tags
type TagAccessor interface {
	CreateTag(ctx context.Context, uuid, name string) (*models.Tag, error)
	GetTagByUUID(ctx context.Context, uuid string) (*models.Tag, error)
}

// Ensure DBAccessor implements TagAccessor
var _ TagAccessor = (*DBAccessor)(nil)

// Create creates a new tag in the database
func (a *DBAccessor) CreateTag(ctx context.Context, uuid, name string) (*models.Tag, error) {
	now := time.Now()

	// Insert the tag
	insert := a.Builder.Insert("tags").Rows(goqu.Record{
		"uuid":       uuid,
		"name":       name,
		"created_at": now,
		"updated_at": now,
	})

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

	// Return the created tag
	return &models.Tag{
		ID:        id,
		UUID:      uuid,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// GetTagByUUID retrieves a tag by its UUID
func (a *DBAccessor) GetTagByUUID(ctx context.Context, uuid string) (*models.Tag, error) {
	var tag models.Tag

	query, args, err := a.Builder.From("tags").
		Select("id", "uuid", "name", "created_at", "updated_at").
		Where(goqu.C("uuid").Eq(uuid)).
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
