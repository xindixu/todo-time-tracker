package testhelpers_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"todo-time-tracker/db"
	"todo-time-tracker/db/accessors"
	"todo-time-tracker/db/models"
	goutils "todo-time-tracker/go-utils"
)

// TestDBAccessor is a test database accessor
type TestDBAccessor struct {
	accessors.DBAccessor
}

// NewTestDBAccessor creates a new test database accessor
func NewTestDBAccessor(dbConnection *db.Connection) *TestDBAccessor {
	return &TestDBAccessor{
		*accessors.NewDBAccessor(dbConnection),
	}
}

// GetRandomUser gets a random user
func (a *TestDBAccessor) GetRandomUser() *models.User {
	return &models.User{
		UUID:     uuid.New(),
		Name:     goutils.RandString(10),
		Email:    goutils.RandEmail(),
		Password: goutils.RandString(10),
	}
}

// GetRandomTask gets a random task
func (a *TestDBAccessor) GetRandomTask() *models.Task {
	duration := goutils.RandDuration(10)
	return &models.Task{
		UUID:              uuid.New(),
		Name:              goutils.RandString(10),
		Description:       goutils.RandString(10),
		EstimatedDuration: &duration,
		Status:            models.TaskStatusTodo,
	}
}

// CreateTestUser creates a test user
func (a *TestDBAccessor) CreateTestUser(t *testing.T) *models.User {
	user := a.GetRandomUser()
	user, err := a.CreateUser(t.Context(), user.UUID, user.Name, user.Email, user.Password)
	require.NoErrorf(t, err, "Failed to create test user")
	return user
}

// CreateTestUserWithUUID creates a test user with a specific UUID
func (a *TestDBAccessor) CreateTestUserWithUUID(t *testing.T, uuid uuid.UUID) *models.User {
	user := a.GetRandomUser()
	_, err := a.CreateUser(t.Context(), uuid, user.Name, user.Email, user.Password)
	require.NoErrorf(t, err, "Failed to create test user")
	return user
}

// CreateTestTask creates a test task
func (a *TestDBAccessor) CreateTestTask(t *testing.T, user *models.User) *models.Task {
	task := a.GetRandomTask()
	ctx := context.WithValue(t.Context(), goutils.ContextKeyAccountID, user.AccountID)
	task, err := a.CreateTask(ctx, task.UUID, task.Name, task.Description, task.EstimatedDuration, task.Status)
	require.NoErrorf(t, err, "Failed to create test task")
	return task
}
