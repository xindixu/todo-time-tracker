package accessors

import (
	"context"
	"time"

	"todo-time-tracker/db/models"
	goutils "todo-time-tracker/go-utils"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type TaskAccessor interface {
	CreateTask(ctx context.Context, uuid uuid.UUID, name string, description string, estimatedDuration *time.Duration, status models.TaskStatus) (*models.Task, error)
	GetTaskByUUID(ctx context.Context, uuid string) (*models.Task, error)
}

// Ensure DBAccessor implements TaskAccessor
var _ TaskAccessor = (*DBAccessor)(nil)

// CreateTask creates a new task in the database
func (a *DBAccessor) CreateTask(ctx context.Context, uuid uuid.UUID, name string, description string, estimatedDuration *time.Duration, status models.TaskStatus) (*models.Task, error) {
	accountID := ctx.Value(goutils.ContextKeyAccountID).(int64)

	task := &models.Task{
		UUID:              uuid,
		Name:              name,
		Description:       description,
		EstimatedDuration: estimatedDuration,
		Status:            status,
		AccountID:         accountID,
	}

	query, args, err := a.Builder.Insert(models.TasksTable).Rows(task).
		Returning("id", "created_at", "updated_at").ToSQL()
	if err != nil {
		return nil, err
	}

	var id int64
	var createdAt time.Time
	var updatedAt time.Time
	err = a.DB.QueryRowxContext(ctx, query, args...).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	task.ID = id
	task.CreatedAt = createdAt
	task.UpdatedAt = updatedAt
	return task, nil
}

// GetTaskByUUID retrieves a task by its UUID
func (a *DBAccessor) GetTaskByUUID(ctx context.Context, uuid string) (*models.Task, error) {
	tasksTable := goqu.T(models.TasksTable)
	query, args, err := a.Builder.From(tasksTable).Where(tasksTable.Col("uuid").Eq(uuid)).ToSQL()
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = a.DB.QueryRowxContext(ctx, query, args...).StructScan(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
