# Database Accessor Tests

This directory contains comprehensive unit tests for all database accessor functionality.

## Test Structure

The tests are organized using the `testify/suite` package and include:

- **Test Suites**:
  - `UserAccessorTestSuite` - Tests for user management functionality
  - `TaskAccessorTestSuite` - Tests for task management functionality
- **Test Helpers**: `test_helpers.go` - Utilities for database testing
- **Individual Tests**: Simple tests that don't require database connection

## Running the Tests

### Prerequisites

1. **PostgreSQL**: You need PostgreSQL running locally or accessible via environment variables
2. **Dependencies**: All Go dependencies should be installed (`go mod tidy`)

### Environment Variables

You can configure the test database connection using these environment variables:

```bash
export TEST_DB_HOST=localhost
export TEST_DB_PORT=5432
export TEST_DB_USER=postgres
export TEST_DB_PASSWORD=postgres
```

If not set, the tests will use these defaults:
- Host: `localhost`
- Port: `5432`
- User: `postgres`
- Password: `postgres`

### Running Tests

```bash
# Run all tests in the accessors package
go test ./db/accessors -v

# Run with coverage
go test ./db/accessors -v -cover

# Run the user accessor test suite
go test ./db/accessors -v -run "TestUserAccessorSuite"

# Run a specific test method
go test ./db/accessors -v -run "TestUserAccessorSuite/TestCreateUser_Success"
```

## Test Database Management

The tests automatically:

1. **Create** a unique test database for each test run
2. **Migrate** the database schema using the migration files
3. **Clean up** test data between individual tests
4. **Drop** the test database when all tests complete

### Test Database Naming

Test databases are named with the pattern: `test_ttt_{timestamp}`

Example: `test_ttt_1703123456789012345`

## Test Data Cleanup

Before each test, the following tables are cleaned up in dependency order:

1. `task_tags`
2. `task_users`
3. `organization_users`
4. `sessions`
5. `tasks`
6. `tags`
7. `users`
8. `organizations`
9. `accounts`

## Test Helpers

The test helpers provide utilities for creating test data:

### TestDBAccessor Methods
- `CreateTestUser(t *testing.T)` - Creates a random user for testing
- `CreateTestUserWithUUID(t *testing.T, uuid uuid.UUID)` - Creates a user with specific UUID
- `CreateTestTask(t *testing.T, user *models.User)` - Creates a random task for a user
- `GetRandomUser()` - Returns a random user model (not saved to DB)
- `GetRandomTask()` - Returns a random task model (not saved to DB)

### Database Utilities
- `IsPostgreSQLAvailable()` - Check if PostgreSQL is accessible
- `CreateTestSQLDB()` - Create and configure a test database
- `CleanupTestSQLDB()` - Remove all test data
- `SkipIfNoPostgreSQL()` - Skip tests if PostgreSQL is unavailable
- `DefaultTestDBConfig()` - Get default test database configuration

## Troubleshooting

### PostgreSQL Not Available

If PostgreSQL is not running or accessible, the tests will be skipped:

```
=== RUN   TestUserAccessorSuite
    user_accessor_test.go:XX: PostgreSQL not available, skipping test
--- SKIP: TestUserAccessorSuite (0.00s)
```

### Database Connection Issues

If you encounter connection issues:

1. Ensure PostgreSQL is running
2. Check your environment variables
3. Verify the postgres user has permissions to create databases
4. Check firewall settings if using a remote database

### Permission Issues

The postgres user needs these permissions:
- `CREATEDB` - to create test databases
- `CONNECT` - to connect to databases
- `USAGE` - to use the postgres database

## Adding New Tests

When adding new tests to the suite:

1. Add test methods to the appropriate test suite (`UserAccessorTestSuite` or `TaskAccessorTestSuite`)
2. Use the existing `suite.accessor` and `suite.dbConnection`
3. Follow the naming convention: `Test{MethodName}_{Scenario}`
4. Use `require` for setup assertions and `assert` for test assertions

Example:
```go
func (suite *UserAccessorTestSuite) TestCreateUser_NewScenario() {
    // Test implementation
}
```
