package accessors

import (
	"context"
	"fmt"
	"time"

	"todo-time-tracker/db/models"
	goutils "todo-time-tracker/go-utils"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type TaskAccessor interface {
	CreateTask(ctx context.Context, uuid uuid.UUID, name string, description string, estimatedDuration *time.Duration, status models.TaskStatus) (*models.Task, error)
	GetTaskByUUID(ctx context.Context, uuid uuid.UUID) (*models.Task, error)
	LinkTasks(ctx context.Context, fromTaskUUID uuid.UUID, toTaskUUID uuid.UUID, link models.TaskLink) error
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

	q := a.SQLBuilder.Insert(models.TasksTable).Rows(task).
		Returning("id", "created_at", "updated_at")

	query, args, err := q.ToSQL()
	if err != nil {
		return nil, err
	}

	var id int64
	var createdAt time.Time
	var updatedAt time.Time
	err = a.SQLDB.QueryRowxContext(ctx, query, args...).Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	task.ID = id
	task.CreatedAt = createdAt
	task.UpdatedAt = updatedAt
	return task, nil
}

// GetTaskByUUID retrieves a task by its UUID
func (a *DBAccessor) GetTaskByUUID(ctx context.Context, uuid uuid.UUID) (*models.Task, error) {
	tasksTable := goqu.T(models.TasksTable)
	q := a.SQLBuilder.From(tasksTable).Where(tasksTable.Col("uuid").Eq(uuid.String()))

	query, args, err := q.ToSQL()
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = a.SQLDB.QueryRowxContext(ctx, query, args...).StructScan(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// LinkTasks links two tasks together
func (a *DBAccessor) LinkTasks(ctx context.Context, fromTaskUUID uuid.UUID, toTaskUUID uuid.UUID, link models.TaskLink) error {
	session, err := a.DBConnection.NewGraphDBSession(ctx)
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`
		MERGE (from:Task {uuid: $fromTaskUUID})
		MERGE (to:Task {uuid: $toTaskUUID})
		MERGE (from)-[:%s]->(to)
	`, link)

	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, q,
			map[string]any{
				"fromTaskUUID": fromTaskUUID.String(),
				"toTaskUUID":   toTaskUUID.String(),
			},
		)
	})

	return err
}
