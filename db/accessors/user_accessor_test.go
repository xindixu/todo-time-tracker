package accessors_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"todo-time-tracker/db"
	"todo-time-tracker/db/models"

	th "todo-time-tracker/db/test-helpers"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// UserAccessorTestSuite holds the test suite for user accessor tests
type UserAccessorTestSuite struct {
	suite.Suite
	ctx          context.Context
	dbConnection *db.DBConnection
	accessor     *th.TestDBAccessor
	cleanup      func()
}

// SetupSuite runs once before all tests
func (s *UserAccessorTestSuite) SetupSuite() {
	// Skip if PostgreSQL is not available
	th.SkipIfNoPostgreSQL(s.T())

	// Create test database
	s.dbConnection, s.cleanup = th.CreateTestDB(s.T(), nil)

	// Create accessor
	s.accessor = th.CreateTestDBAccessor(s.T(), s.dbConnection)

	s.ctx = context.Background()
}

// TearDownSuite runs once after all tests
func (s *UserAccessorTestSuite) TearDownSuite() {
	if s.cleanup != nil {
		s.cleanup()
	}
}

// SetupTest runs before each test
func (s *UserAccessorTestSuite) SetupTest() {
	// Clean up any existing data before each test
	th.CleanupTestSQLDB(s.T(), s.dbConnection)
}

// TestMain runs the test suite
func TestUserAccessorSuite(t *testing.T) {
	suite.Run(t, new(UserAccessorTestSuite))
}

// TestCreateUser_Success tests successful user creation
func (s *UserAccessorTestSuite) TestCreateUser_Success() {
	// Test data
	testUUID := uuid.New()
	testName := "John Doe"
	testEmail := "john@example.com"
	testPassword := "password123"

	// Execute
	user, err := s.accessor.CreateUser(s.ctx, testUUID, testName, testEmail, testPassword)

	// Assert
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), user)
	assert.Equal(s.T(), testUUID, user.UUID)
	assert.Equal(s.T(), testName, user.Name)
	assert.Equal(s.T(), testEmail, user.Email)
	assert.Equal(s.T(), testPassword, user.Password)
	assert.Greater(s.T(), user.ID, int64(0))
	assert.Greater(s.T(), user.AccountID, int64(0))
	assert.NotZero(s.T(), user.CreatedAt)
	assert.NotZero(s.T(), user.UpdatedAt)

	// Verify the user was actually created in the database
	userAccount, err := s.accessor.GetUserAccountByUUID(s.ctx, testUUID.String())
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), userAccount)
	assert.Equal(s.T(), testUUID, userAccount.User.UUID)
	assert.Equal(s.T(), testName, userAccount.User.Name)
	assert.Equal(s.T(), testEmail, userAccount.User.Email)
	assert.Equal(s.T(), testPassword, userAccount.User.Password)

	assert.Equal(s.T(), testUUID, userAccount.Account.UUID)
	assert.Equal(s.T(), models.AccountTypeUser, userAccount.Account.Type)
	assert.Greater(s.T(), userAccount.Account.ID, int64(0))
	assert.NotZero(s.T(), userAccount.Account.CreatedAt)
	assert.NotZero(s.T(), userAccount.Account.UpdatedAt)
}

// TestCreateUser_DuplicateUUID tests creating a user with duplicate UUID
func (s *UserAccessorTestSuite) TestCreateUser_DuplicateUUID() {
	// Create first user
	testUUID := uuid.New()
	testName1 := "John Doe"
	testEmail1 := "john@example.com"
	testPassword1 := "password123"

	user1, err := s.accessor.CreateUser(s.ctx, testUUID, testName1, testEmail1, testPassword1)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), user1)

	// Try to create second user with same UUID
	testName2 := "Jane Doe"
	testEmail2 := "jane@example.com"
	testPassword2 := "password456"

	user2, err := s.accessor.CreateUser(s.ctx, testUUID, testName2, testEmail2, testPassword2)

	// Should fail due to unique constraint on UUID
	assert.Error(s.T(), err)
	assert.Nil(s.T(), user2)
}

// TestCreateUser_DuplicateEmail tests creating a user with duplicate email
func (s *UserAccessorTestSuite) TestCreateUser_DuplicateEmail() {
	// Create first user
	testUUID1 := uuid.New()
	testName1 := "John Doe"
	testEmail := "john@example.com" // Same email
	testPassword1 := "password123"

	user1, err := s.accessor.CreateUser(s.ctx, testUUID1, testName1, testEmail, testPassword1)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), user1)

	// Try to create second user with same email
	testUUID2 := uuid.New()
	testName2 := "Jane Doe"
	testPassword2 := "password456"

	user2, err := s.accessor.CreateUser(s.ctx, testUUID2, testName2, testEmail, testPassword2)

	// Should fail due to unique constraint on email
	assert.Error(s.T(), err)
	assert.Nil(s.T(), user2)
}

// TestCreateUser_EmptyFields tests creating a user with empty required fields
func (s *UserAccessorTestSuite) TestCreateUser_EmptyFields() {
	testCases := []struct {
		name     string
		uuid     uuid.UUID
		userName string
		email    string
		password string
	}{
		{"empty name", uuid.New(), "", "test@example.com", "password123"},
		{"empty email", uuid.New(), "John Doe", "", "password123"},
		{"empty password", uuid.New(), "John Doe", "test@example.com", ""},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			user, err := s.accessor.CreateUser(s.ctx, tc.uuid, tc.userName, tc.email, tc.password)

			// Should fail due to NOT NULL constraints
			assert.Error(t, err)
			assert.Nil(t, user)
		})
	}
}

// TestCreateUser_WithContext tests creating a user with context timeout
func (s *UserAccessorTestSuite) TestCreateUser_WithContext() {
	// Test data
	testUUID := uuid.New()
	testName := "Jane Smith"
	testEmail := "jane@example.com"
	testPassword := "securepassword"

	// Execute with context that has timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.accessor.CreateUser(ctx, testUUID, testName, testEmail, testPassword)

	// Assert
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), user)
	assert.Equal(s.T(), testUUID, user.UUID)
	assert.Equal(s.T(), testName, user.Name)
	assert.Equal(s.T(), testEmail, user.Email)
	assert.Equal(s.T(), testPassword, user.Password)
}

// TestCreateUser_MultipleUsers tests creating multiple users successfully
func (s *UserAccessorTestSuite) TestCreateUser_MultipleUsers() {
	users := []struct {
		uuid     uuid.UUID
		name     string
		email    string
		password string
	}{
		{uuid.New(), "Alice Johnson", "alice@example.com", "password1"},
		{uuid.New(), "Bob Smith", "bob@example.com", "password2"},
		{uuid.New(), "Charlie Brown", "charlie@example.com", "password3"},
	}

	for i, userData := range users {
		s.T().Run(fmt.Sprintf("user_%d", i), func(t *testing.T) {
			user, err := s.accessor.CreateUser(s.ctx, userData.uuid, userData.name, userData.email, userData.password)

			require.NoError(t, err)
			assert.NotNil(t, user)
			assert.Equal(t, userData.uuid, user.UUID)
			assert.Equal(t, userData.name, user.Name)
			assert.Equal(t, userData.email, user.Email)
			assert.Equal(t, userData.password, user.Password)
		})
	}

	// Verify all users were created
	var count int
	err := s.dbConnection.SQLDB.Get(&count, "SELECT COUNT(*) FROM users")
	require.NoError(s.T(), err)
	assert.Equal(s.T(), len(users), count)

	// Verify all accounts were created
	err = s.dbConnection.SQLDB.Get(&count, "SELECT COUNT(*) FROM accounts WHERE type = $1", models.AccountTypeUser)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), len(users), count)
}

// TestCreateUser_TransactionRollback tests that transaction rollback works correctly
func (s *UserAccessorTestSuite) TestCreateUser_TransactionRollback() {
	// This test verifies that if there's an error during user creation,
	// the transaction is properly rolled back and no data is left in the database

	// Get initial counts
	initialUserCount, err := s.accessor.CountModels(s.ctx, models.UsersTable)
	require.NoError(s.T(), err)
	initialAccountCount, err := s.accessor.CountModels(s.ctx, models.AccountsTable)
	require.NoError(s.T(), err)

	// Create a user successfully first
	testUUID1 := uuid.New()
	testName1 := "John Doe"
	testEmail1 := "john@example.com"
	testPassword1 := "password123"

	user1, err := s.accessor.CreateUser(s.ctx, testUUID1, testName1, testEmail1, testPassword1)
	require.NoError(s.T(), err)
	assert.NotNil(s.T(), user1)

	// Try to create another user with duplicate email (should fail)
	testUUID2 := uuid.New()
	testName2 := "Jane Doe"
	testEmail2 := testEmail1 // Same email as first user
	testPassword2 := "password456"

	user2, err := s.accessor.CreateUser(s.ctx, testUUID2, testName2, testEmail2, testPassword2)
	assert.Error(s.T(), err)
	assert.Nil(s.T(), user2)

	// Verify that only the first user exists (transaction was rolled back)
	finalUserCount, err := s.accessor.CountModels(s.ctx, models.UsersTable)
	require.NoError(s.T(), err)
	finalAccountCount, err := s.accessor.CountModels(s.ctx, models.AccountsTable)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), initialUserCount+1, finalUserCount)       // Only first user was added
	assert.Equal(s.T(), initialAccountCount+1, finalAccountCount) // Only first account was added

	// Verify the second user doesn't exist
	var count int
	err = s.dbConnection.SQLDB.Get(&count, "SELECT COUNT(*) FROM users WHERE uuid = $1", testUUID2)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), 0, count)
}
