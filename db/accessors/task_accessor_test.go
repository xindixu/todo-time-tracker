package accessors_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"todo-time-tracker/db"
	"todo-time-tracker/db/models"
	goutils "todo-time-tracker/go-utils"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	th "todo-time-tracker/db/test-helpers"
)

// TaskAccessorTestSuite holds the test suite for task accessor tests
type TaskAccessorTestSuite struct {
	suite.Suite
	ctx          context.Context
	dbConnection *db.DBConnection
	accessor     *th.TestDBAccessor
	cleanup      func()
}

// SetupSuite runs once before all tests
func (s *TaskAccessorTestSuite) SetupSuite() {
	s.ctx = context.Background()

	// Skip if PostgreSQL is not available
	th.SkipIfNoPostgreSQL(s.T())

	// Create test database
	s.dbConnection, s.cleanup = th.CreateTestDB(s.T(), nil)

	// Create accessor
	s.accessor = th.CreateTestDBAccessor(s.T(), s.dbConnection)
}

// TearDownSuite runs once after all tests
func (s *TaskAccessorTestSuite) TearDownSuite() {
	if s.cleanup != nil {
		s.cleanup()
	}
}

// SetupTest runs before each test
func (s *TaskAccessorTestSuite) SetupTest() {
	// Clean up any existing data before each test
	th.CleanupTestSQLDB(s.T(), s.dbConnection)
}

// TestMain runs the test suite
func TestTaskAccessorSuite(t *testing.T) {
	suite.Run(t, new(TaskAccessorTestSuite))
}

// TestCreateTask_Success tests successful task creation
func (s *TaskAccessorTestSuite) TestCreateTask_Success() {
	// Create a user first to get a valid account ID
	user := s.accessor.CreateTestUser(s.T())

	// Test data
	testUUID := uuid.New()
	testName := "Test Task"
	testDescription := "This is a test task"
	testEstimatedDuration := time.Duration(30 * time.Minute)
	testStatus := models.TaskStatusTodo

	// Create context with account ID from the created user
	ctxWithAcct := context.WithValue(s.ctx, goutils.ContextKeyAccountID, user.AccountID)

	// Execute
	task, err := s.accessor.CreateTask(ctxWithAcct, testUUID, testName, testDescription, &testEstimatedDuration, testStatus)

	// Assert
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), task)
	assert.Equal(s.T(), testUUID, task.UUID)
	assert.Equal(s.T(), testName, task.Name)
	assert.Equal(s.T(), testDescription, task.Description)
	assert.Equal(s.T(), &testEstimatedDuration, task.EstimatedDuration)
	assert.Equal(s.T(), testStatus, task.Status)
	assert.Equal(s.T(), user.AccountID, task.AccountID)
	assert.Greater(s.T(), task.ID, int64(0))
	assert.NotZero(s.T(), task.CreatedAt)
	assert.NotZero(s.T(), task.UpdatedAt)

	// Verify the task was actually created in the database
	retrievedTask, err := s.accessor.GetTaskByUUID(s.ctx, testUUID)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), retrievedTask)
	assert.Equal(s.T(), testUUID, retrievedTask.UUID)
	assert.Equal(s.T(), testName, retrievedTask.Name)
	assert.Equal(s.T(), testDescription, retrievedTask.Description)
	assert.Equal(s.T(), &testEstimatedDuration, retrievedTask.EstimatedDuration)
	assert.Equal(s.T(), testStatus, retrievedTask.Status)
	assert.Equal(s.T(), user.AccountID, retrievedTask.AccountID)
}

// TestCreateTask_WithNilEstimatedDuration tests task creation with nil estimated duration
func (s *TaskAccessorTestSuite) TestCreateTask_WithNilEstimatedDuration() {
	user := s.accessor.CreateTestUser(s.T())

	// Test data
	testUUID := uuid.New()
	testName := "Test Task"
	testDescription := "This is a test task"
	testStatus := models.TaskStatusTodo

	// Create context with account ID from the created user
	ctxWithAcct := context.WithValue(s.ctx, goutils.ContextKeyAccountID, user.AccountID)

	// Execute
	task, err := s.accessor.CreateTask(ctxWithAcct, testUUID, testName, testDescription, nil, testStatus)

	// Assert
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), task)
	assert.Equal(s.T(), testUUID, task.UUID)
	assert.Equal(s.T(), testName, task.Name)
	assert.Equal(s.T(), testDescription, task.Description)
	assert.Nil(s.T(), task.EstimatedDuration)
	assert.Equal(s.T(), testStatus, task.Status)
	assert.Equal(s.T(), user.AccountID, task.AccountID)
}

// TestCreateTask_DifferentStatuses tests task creation with different statuses
func (s *TaskAccessorTestSuite) TestCreateTask_DifferentStatuses() {
	user := s.accessor.CreateTestUser(s.T())
	ctxWithAcct := context.WithValue(s.ctx, goutils.ContextKeyAccountID, user.AccountID)

	testCases := []struct {
		name   string
		status models.TaskStatus
	}{
		{"TODO status", models.TaskStatusTodo},
		{"IN_PROGRESS status", models.TaskStatusInProgress},
		{"BLOCKED status", models.TaskStatusBlocked},
		{"DONE status", models.TaskStatusDone},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			testUUID := uuid.New()
			testName := fmt.Sprintf("Test Task %s", tc.status)
			testDescription := "This is a test task"

			task, err := s.accessor.CreateTask(ctxWithAcct, testUUID, testName, testDescription, nil, tc.status)

			require.NoError(t, err)
			assert.NotNil(t, task)
			assert.Equal(t, tc.status, task.Status)

			// Verify retrieval
			retrievedTask, err := s.accessor.GetTaskByUUID(s.ctx, testUUID)
			require.NoError(t, err)
			assert.Equal(t, tc.status, retrievedTask.Status)
		})
	}
}

// TestCreateTask_DuplicateUUID tests creating a task with duplicate UUID
func (s *TaskAccessorTestSuite) TestCreateTask_DuplicateUUID() {
	user := s.accessor.CreateTestUser(s.T())
	ctxWithAcct := context.WithValue(s.ctx, goutils.ContextKeyAccountID, user.AccountID)

	// Create first task
	testUUID := uuid.New()
	testName1 := "First Task"
	testDescription1 := "First task description"
	testStatus1 := models.TaskStatusTodo

	task1, err := s.accessor.CreateTask(ctxWithAcct, testUUID, testName1, testDescription1, nil, testStatus1)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), task1)

	// Try to create second task with same UUID
	testName2 := "Second Task"
	testDescription2 := "Second task description"
	testStatus2 := models.TaskStatusInProgress

	task2, err := s.accessor.CreateTask(ctxWithAcct, testUUID, testName2, testDescription2, nil, testStatus2)

	// Should fail due to unique constraint on UUID
	assert.Error(s.T(), err)
	assert.Nil(s.T(), task2)
}

// TestCreateTask_EmptyFields tests creating a task with empty fields (which are allowed)
func (s *TaskAccessorTestSuite) TestCreateTask_EmptyFields() {
	user := s.accessor.CreateTestUser(s.T())
	ctxWithAcct := context.WithValue(s.ctx, goutils.ContextKeyAccountID, user.AccountID)

	testCases := []struct {
		name        string
		taskName    string
		description string
	}{
		{"empty name", "", "Test description"},
		{"empty description", "Test Task", ""},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			testUUID := uuid.New()
			testStatus := models.TaskStatusTodo

			task, err := s.accessor.CreateTask(ctxWithAcct, testUUID, tc.taskName, tc.description, nil, testStatus)

			// Empty fields are allowed by the database schema
			require.NoError(t, err)
			assert.NotNil(t, task)
			assert.Equal(t, tc.taskName, task.Name)
			assert.Equal(t, tc.description, task.Description)
		})
	}
}

// TestCreateTask_WithoutAccountID tests creating a task without account ID in context
func (s *TaskAccessorTestSuite) TestCreateTask_WithoutAccountID() {
	// Test data
	task := s.accessor.GetRandomTask()

	// Execute - this should panic due to missing account ID in context
	assert.Panics(s.T(), func() {
		_, _ = s.accessor.CreateTask(s.ctx, task.UUID, task.Name, task.Description, nil, task.Status)
	})
}

// TestCreateTask_MultipleTasks tests creating multiple tasks
func (s *TaskAccessorTestSuite) TestCreateTask_MultipleTasks() {
	user := s.accessor.CreateTestUser(s.T())
	ctxWithAcct := context.WithValue(s.ctx, goutils.ContextKeyAccountID, user.AccountID)

	// Create multiple tasks
	tasks := []struct {
		uuid        uuid.UUID
		name        string
		description string
		status      models.TaskStatus
	}{
		{uuid.New(), "Task 1", "First task", models.TaskStatusTodo},
		{uuid.New(), "Task 2", "Second task", models.TaskStatusInProgress},
		{uuid.New(), "Task 3", "Third task", models.TaskStatusDone},
	}

	for i, tc := range tasks {
		s.T().Run(fmt.Sprintf("task_%d", i+1), func(t *testing.T) {
			task, err := s.accessor.CreateTask(ctxWithAcct, tc.uuid, tc.name, tc.description, nil, tc.status)

			require.NoError(t, err)
			assert.NotNil(t, task)
			assert.Equal(t, tc.uuid, task.UUID)
			assert.Equal(t, tc.name, task.Name)
			assert.Equal(t, tc.description, task.Description)
			assert.Equal(t, tc.status, task.Status)

			// Verify retrieval
			retrievedTask, err := s.accessor.GetTaskByUUID(s.ctx, tc.uuid)
			require.NoError(t, err)
			assert.Equal(t, tc.uuid, retrievedTask.UUID)
		})
	}
}

// TestCreateTask_WithContext tests creating a task with context timeout
func (s *TaskAccessorTestSuite) TestCreateTask_WithContext() {
	user := s.accessor.CreateTestUser(s.T())
	// Create context with timeout and account ID
	ctxWithAcct := context.WithValue(s.ctx, goutils.ContextKeyAccountID, user.AccountID)
	ctxWithTimeout, cancel := context.WithTimeout(ctxWithAcct, 5*time.Second)
	defer cancel()

	// Test data
	task := s.accessor.GetRandomTask()

	// Execute
	task, err := s.accessor.CreateTask(ctxWithTimeout, task.UUID, task.Name, task.Description, nil, task.Status)

	// Assert
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), task)
	assert.Equal(s.T(), task.UUID, task.UUID)
}

// TestGetTaskByUUID_Success tests successful task retrieval by UUID
func (s *TaskAccessorTestSuite) TestGetTaskByUUID_Success() {
	user := s.accessor.CreateTestUser(s.T())

	task := s.accessor.CreateTestTask(s.T(), user)

	// Retrieve the task
	retrievedTask, err := s.accessor.GetTaskByUUID(s.ctx, task.UUID)

	// Assert
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), retrievedTask)
	assert.Equal(s.T(), task.UUID, retrievedTask.UUID)
	assert.Equal(s.T(), task.Name, retrievedTask.Name)
	assert.Equal(s.T(), task.Description, retrievedTask.Description)
	assert.Equal(s.T(), task.EstimatedDuration, retrievedTask.EstimatedDuration)
	assert.Equal(s.T(), task.Status, retrievedTask.Status)
	assert.Equal(s.T(), user.AccountID, retrievedTask.AccountID)
	assert.Equal(s.T(), task.ID, retrievedTask.ID)
	assert.Equal(s.T(), task.CreatedAt, retrievedTask.CreatedAt)
	assert.Equal(s.T(), task.UpdatedAt, retrievedTask.UpdatedAt)
}

// TestGetTaskByUUID_NotFound tests retrieving a non-existent task
func (s *TaskAccessorTestSuite) TestGetTaskByUUID_NotFound() {
	nonExistentUUID := uuid.New()

	// Try to retrieve non-existent task
	task, err := s.accessor.GetTaskByUUID(s.ctx, nonExistentUUID)

	// Should return error for non-existent task
	assert.Error(s.T(), err)
	assert.Nil(s.T(), task)
}

// TestCreateTaskLinks_Success tests successful task linking
func (s *TaskAccessorTestSuite) TestCreateTaskLinks_Success() {
	// Create test user and tasks
	user := s.accessor.CreateTestUser(s.T())
	task1 := s.accessor.CreateTestTask(s.T(), user)
	task2 := s.accessor.CreateTestTask(s.T(), user)

	// Test different link types
	testCases := []struct {
		name     string
		linkType models.TaskLink
	}{
		{"Parent of link", models.TaskLinkParentOf},
		{"Blocks link", models.TaskLinkBlocks},
		{"Relates to link", models.TaskLinkRelatesTo},
		{"Duplicate of link", models.TaskLinkDuplicateOf},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			// Execute
			err := s.accessor.CreateTaskLinks(s.ctx, task1.UUID, task2.UUID, tc.linkType)

			// Assert
			require.NoError(t, err)

			// Verify the link was created in Neo4j
			session, err := s.accessor.DBConnection.NewGraphDBSession(s.ctx)
			require.NoError(t, err)
			defer session.Close(s.ctx)

			result, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
				query := fmt.Sprintf(`
					MATCH (from:Task {uuid: $fromTaskUUID})-[r:%s]->(to:Task {uuid: $toTaskUUID})
					RETURN r
				`, tc.linkType)
				return tx.Run(s.ctx, query, map[string]any{
					"fromTaskUUID": task1.UUID.String(),
					"toTaskUUID":   task2.UUID.String(),
				})
			})
			require.NoError(t, err)

			records := result.(neo4j.ResultWithContext)
			record, err := records.Single(s.ctx)
			require.NoError(t, err)
			assert.NotNil(t, record)
		})
	}
}

// TestCreateTaskLinks_SameTask tests linking a task to itself
func (s *TaskAccessorTestSuite) TestCreateTaskLinks_SameTask() {

	user := s.accessor.CreateTestUser(s.T())
	task := s.accessor.CreateTestTask(s.T(), user)

	// Execute
	err := s.accessor.CreateTaskLinks(s.ctx, task.UUID, task.UUID, models.TaskLinkParentOf)

	// Assert - should succeed (Neo4j allows self-relationships)
	require.NoError(s.T(), err)

	// Verify the self-link was created
	session, err := s.accessor.DBConnection.NewGraphDBSession(s.ctx)
	require.NoError(s.T(), err)
	defer session.Close(s.ctx)

	result, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (from:Task {uuid: $taskUUID})-[r:PARENT_OF]->(to:Task {uuid: $taskUUID})
			RETURN r
		`
		return tx.Run(s.ctx, query, map[string]any{
			"taskUUID": task.UUID.String(),
		})
	})
	require.NoError(s.T(), err)

	records := result.(neo4j.ResultWithContext)
	record, err := records.Single(s.ctx)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), record)
}

// TestCreateTaskLinks_NonExistentTasks tests linking non-existent tasks
func (s *TaskAccessorTestSuite) TestCreateTaskLinks_NonExistentTasks() {
	nonExistentUUID1 := uuid.New()
	nonExistentUUID2 := uuid.New()

	// Execute
	err := s.accessor.CreateTaskLinks(s.ctx, nonExistentUUID1, nonExistentUUID2, models.TaskLinkParentOf)

	// Assert - should succeed because MERGE creates nodes if they don't exist
	require.NoError(s.T(), err)

	// Verify the nodes and relationship were created
	session, err := s.accessor.DBConnection.NewGraphDBSession(s.ctx)
	require.NoError(s.T(), err)
	defer session.Close(s.ctx)

	result, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (from:Task {uuid: $fromTaskUUID})-[r:PARENT_OF]->(to:Task {uuid: $toTaskUUID})
			RETURN from, r, to
		`
		return tx.Run(s.ctx, query, map[string]any{
			"fromTaskUUID": nonExistentUUID1.String(),
			"toTaskUUID":   nonExistentUUID2.String(),
		})
	})
	require.NoError(s.T(), err)

	records := result.(neo4j.ResultWithContext)
	record, err := records.Single(s.ctx)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), record)
}

// TestCreateTaskLinks_DuplicateLink tests creating the same link multiple times
func (s *TaskAccessorTestSuite) TestCreateTaskLinks_DuplicateLink() {

	if s.accessor.GraphDB == nil {
		s.T().Skip("Neo4j not available, skipping test")
	}
	user := s.accessor.CreateTestUser(s.T())
	task1 := s.accessor.CreateTestTask(s.T(), user)
	task2 := s.accessor.CreateTestTask(s.T(), user)

	// Create the same link multiple times
	for i := 0; i < 3; i++ {
		err := s.accessor.CreateTaskLinks(s.ctx, task1.UUID, task2.UUID, models.TaskLinkBlocks)
		require.NoError(s.T(), err)
	}

	// Verify only one relationship exists (MERGE prevents duplicates)
	session, err := s.accessor.DBConnection.NewGraphDBSession(s.ctx)
	require.NoError(s.T(), err)
	defer session.Close(s.ctx)

	result, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (from:Task {uuid: $fromTaskUUID})-[r:BLOCKS]->(to:Task {uuid: $toTaskUUID})
			RETURN count(r) as relationshipCount
		`
		return tx.Run(s.ctx, query, map[string]any{
			"fromTaskUUID": task1.UUID.String(),
			"toTaskUUID":   task2.UUID.String(),
		})
	})
	require.NoError(s.T(), err)

	records := result.(neo4j.ResultWithContext)
	record, err := records.Single(s.ctx)
	require.NoError(s.T(), err)

	var count int64
	count = record.Values[0].(int64)
	assert.Equal(s.T(), int64(1), count)
}

// TestCreateTaskLinks_InvalidLinkType tests linking with invalid link type
func (s *TaskAccessorTestSuite) TestCreateTaskLinks_InvalidLinkType() {
	user := s.accessor.CreateTestUser(s.T())
	task1 := s.accessor.CreateTestTask(s.T(), user)
	task2 := s.accessor.CreateTestTask(s.T(), user)

	// Execute with invalid link type
	err := s.accessor.CreateTaskLinks(s.ctx, task1.UUID, task2.UUID, models.TaskLinkInvalid)

	// Assert - should succeed because Neo4j will create the relationship with the invalid type
	// The validation should happen at a higher level if needed
	require.NoError(s.T(), err)

	// Verify the relationship was created with the invalid type
	session, err := s.accessor.DBConnection.NewGraphDBSession(s.ctx)
	require.NoError(s.T(), err)
	defer session.Close(s.ctx)

	result, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (from:Task {uuid: $fromTaskUUID})-[r:INVALID]->(to:Task {uuid: $toTaskUUID})
			RETURN r
		`
		return tx.Run(s.ctx, query, map[string]any{
			"fromTaskUUID": task1.UUID.String(),
			"toTaskUUID":   task2.UUID.String(),
		})
	})
	require.NoError(s.T(), err)

	records := result.(neo4j.ResultWithContext)
	record, err := records.Single(s.ctx)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), record)
}

// TestCreateTaskLinks_MultipleRelationships tests creating multiple different relationships between the same tasks
func (s *TaskAccessorTestSuite) TestCreateTaskLinks_MultipleRelationships() {
	user := s.accessor.CreateTestUser(s.T())
	task1 := s.accessor.CreateTestTask(s.T(), user)
	task2 := s.accessor.CreateTestTask(s.T(), user)

	// Create multiple different relationships
	relationships := []models.TaskLink{
		models.TaskLinkParentOf,
		models.TaskLinkBlocks,
		models.TaskLinkRelatesTo,
		models.TaskLinkDuplicateOf,
	}

	for _, rel := range relationships {
		err := s.accessor.CreateTaskLinks(s.ctx, task1.UUID, task2.UUID, rel)
		require.NoError(s.T(), err)
	}

	// Verify all relationships exist
	session, err := s.accessor.DBConnection.NewGraphDBSession(s.ctx)
	require.NoError(s.T(), err)
	defer session.Close(s.ctx)

	result, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (from:Task {uuid: $fromTaskUUID})-[r]->(to:Task {uuid: $toTaskUUID})
			RETURN type(r) as relationshipType
			ORDER BY type(r)
		`
		return tx.Run(s.ctx, query, map[string]any{
			"fromTaskUUID": task1.UUID.String(),
			"toTaskUUID":   task2.UUID.String(),
		})
	})
	require.NoError(s.T(), err)

	records := result.(neo4j.ResultWithContext)
	var relationshipTypes []string
	for records.Next(s.ctx) {
		record := records.Record()
		var relType string
		relType = record.Values[0].(string)
		relationshipTypes = append(relationshipTypes, relType)
	}

	require.NoError(s.T(), records.Err())
	assert.Len(s.T(), relationshipTypes, 4)
	assert.Contains(s.T(), relationshipTypes, "PARENT_OF")
	assert.Contains(s.T(), relationshipTypes, "BLOCKS")
	assert.Contains(s.T(), relationshipTypes, "RELATES_TO")
	assert.Contains(s.T(), relationshipTypes, "DUPLICATE_OF")
}

// TestCreateTaskLinks_BidirectionalRelationships tests creating relationships in both directions
func (s *TaskAccessorTestSuite) TestCreateTaskLinks_BidirectionalRelationships() {

	user := s.accessor.CreateTestUser(s.T())
	task1 := s.accessor.CreateTestTask(s.T(), user)
	task2 := s.accessor.CreateTestTask(s.T(), user)

	// Create relationships in both directions
	err := s.accessor.CreateTaskLinks(s.ctx, task1.UUID, task2.UUID, models.TaskLinkBlocks)
	require.NoError(s.T(), err)

	err = s.accessor.CreateTaskLinks(s.ctx, task2.UUID, task1.UUID, models.TaskLinkRelatesTo)
	require.NoError(s.T(), err)

	// Verify both relationships exist
	session, err := s.accessor.DBConnection.NewGraphDBSession(s.ctx)
	require.NoError(s.T(), err)
	defer session.Close(s.ctx)

	result, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (from:Task {uuid: $fromTaskUUID})-[r]->(to:Task {uuid: $toTaskUUID})
			RETURN type(r) as relationshipType
			ORDER BY type(r)
		`
		return tx.Run(s.ctx, query, map[string]any{
			"fromTaskUUID": task1.UUID.String(),
			"toTaskUUID":   task2.UUID.String(),
		})
	})
	require.NoError(s.T(), err)

	records := result.(neo4j.ResultWithContext)
	record, err := records.Single(s.ctx)
	require.NoError(s.T(), err)

	var relType string
	relType = record.Values[0].(string)
	assert.Equal(s.T(), "BLOCKS", relType)

	// Check reverse relationship
	result2, err := session.ExecuteRead(s.ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (from:Task {uuid: $fromTaskUUID})-[r]->(to:Task {uuid: $toTaskUUID})
			RETURN type(r) as relationshipType
		`
		return tx.Run(s.ctx, query, map[string]any{
			"fromTaskUUID": task2.UUID.String(),
			"toTaskUUID":   task1.UUID.String(),
		})
	})
	require.NoError(s.T(), err)

	records2 := result2.(neo4j.ResultWithContext)
	record2, err := records2.Single(s.ctx)
	require.NoError(s.T(), err)

	var relType2 string
	relType2 = record2.Values[0].(string)
	assert.Equal(s.T(), "RELATES_TO", relType2)
}
