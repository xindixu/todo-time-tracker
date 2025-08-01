package accessors

import (
	"context"
	"time"

	"todo-time-tracker/db/models"
	goutils "todo-time-tracker/go-utils"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

// TagAccessor handles database operations for tags
type TagAccessor interface {
	CreateTag(ctx context.Context, uuid uuid.UUID, name string) (*models.Tag, error)
	GetTagByUUID(ctx context.Context, uuid uuid.UUID) (*models.Tag, error)
}

// Ensure DBAccessor implements TagAccessor
var _ TagAccessor = (*DBAccessor)(nil)

// CreateTag creates a new tag in the database
func (a *DBAccessor) CreateTag(ctx context.Context, uuid uuid.UUID, name string) (*models.Tag, error) {
	tagsTable := goqu.T(models.TagsTable)

	accountID := ctx.Value(goutils.ContextKeyAccountID).(int64)

	tag := &models.Tag{
		UUID:      uuid,
		Name:      name,
		AccountID: accountID,
	}

	// Build insert query with RETURNING clause to get the inserted ID
	insert := a.SQLBuilder.Insert(tagsTable).
		Rows(tag).
		Returning("id", "created_at", "updated_at")

	query, args, err := insert.ToSQL()
	if err != nil {
		return nil, err
	}

	var insertedID int64
	var createdAt time.Time
	var updatedAt time.Time
	err = a.SQLDB.QueryRowContext(ctx, query, args...).Scan(&insertedID, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	// Set the ID on the tag and return it
	tag.ID = insertedID
	tag.CreatedAt = createdAt
	tag.UpdatedAt = updatedAt
	return tag, nil
}

// GetTagByUUID retrieves a tag by its UUID
func (a *DBAccessor) GetTagByUUID(ctx context.Context, uuid uuid.UUID) (*models.Tag, error) {
	var tag models.Tag

	tagsTable := goqu.T("tags")

	query, args, err := a.SQLBuilder.From(tagsTable).
		Where(tagsTable.Col("uuid").Eq(uuid)).
		ToSQL()
	if err != nil {
		return nil, err
	}

	err = a.SQLDB.QueryRowxContext(ctx, query, args...).StructScan(&tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
