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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TaskAccessor is the interface for the task accessor
type TaskAccessor interface {
	CreateTask(ctx context.Context, uuid uuid.UUID, name string, description string, estimatedDuration *time.Duration, status models.TaskStatus) (*models.Task, error)
	GetTaskByUUID(ctx context.Context, uuid uuid.UUID) (*models.Task, error)
	CreateTaskLinks(ctx context.Context, fromTaskUUID uuid.UUID, toTaskUUID uuid.UUID, link models.TaskLink) error
	GetTaskLinks(ctx context.Context, fromTaskUUID uuid.UUID, toTaskUUID uuid.UUID) ([]models.TaskLink, error)
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

// CreateTaskLinks links two tasks together
func (a *DBAccessor) CreateTaskLinks(ctx context.Context, fromTaskUUID uuid.UUID, toTaskUUID uuid.UUID, link models.TaskLink) error {
	if !link.IsValid() {
		return status.Errorf(codes.InvalidArgument, "invalid link type: %s", link)
	}
	if fromTaskUUID == toTaskUUID {
		return status.Errorf(codes.InvalidArgument, "cannot link task to itself")
	}

	session, err := a.Connection.NewGraphDBSession(ctx)
	if err != nil {
		return err
	}
	defer session.Close(ctx)

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

// GetTaskLinks retrieves the links between two tasks
func (a *DBAccessor) GetTaskLinks(ctx context.Context, fromTaskUUID uuid.UUID, toTaskUUID uuid.UUID) ([]models.TaskLink, error) {
	session, err := a.Connection.NewGraphDBSession(ctx)
	if err != nil {
		return nil, err
	}
	defer session.Close(ctx)

	q := `
		MATCH (from:Task {uuid: $fromTaskUUID})-[r]->(to:Task {uuid: $toTaskUUID})
		ORDER BY type(r)
		RETURN type(r) as relationshipType
	`

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, q,
			map[string]any{
				"fromTaskUUID": fromTaskUUID.String(),
				"toTaskUUID":   toTaskUUID.String(),
			},
		)
		if err != nil {
			return nil, err
		}
		records, err := result.Collect(ctx)
		if err != nil {
			return nil, err
		}
		return records, nil
	})
	if err != nil {
		return nil, err
	}

	neo4jRecords := records.([]*neo4j.Record)
	taskLinks := make([]models.TaskLink, 0, len(neo4jRecords))
	for _, record := range neo4jRecords {
		relationshipType, ok := record.Get("relationshipType")
		if !ok {
			continue
		}
		taskLinks = append(taskLinks, models.TaskLink(relationshipType.(string)))
	}
	return taskLinks, nil
}
