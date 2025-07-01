package accessors

import (
	"context"
	"errors"
	"time"

	"todo-time-tracker/db/models"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type UserAccessor interface {
	CreateUser(ctx context.Context, uuid uuid.UUID, name string, email string, password string) (*models.User, error)
	// GetUserByUUID(ctx context.Context, uuid string) (*models.User, error)
}

// Ensure DBAccessor implements UserAccessor
var _ UserAccessor = (*DBAccessor)(nil)

// CreateUser creates a new user in the database
func (a *DBAccessor) CreateUser(ctx context.Context, uuid uuid.UUID, name string, email string, password string) (*models.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("invalid user data")
	}

	tx, err := a.Builder.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	usersTable := goqu.T("users")
	accountsTable := goqu.T("accounts")

	account := &models.Account{
		UUID: uuid,
		Type: models.AccountTypeUser,
	}

	insertAccount := tx.Insert(accountsTable).Rows(account).Returning("id")

	query, args, err := insertAccount.ToSQL()
	if err != nil {
		return nil, err
	}

	var insertedAccountID int64
	err = tx.QueryRowContext(ctx, query, args...).Scan(&insertedAccountID)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		UUID:      uuid,
		Name:      name,
		Email:     email,
		Password:  password,
		AccountID: insertedAccountID,
	}

	insertUser := tx.Insert(usersTable).
		Rows(user).
		Returning("id", "created_at", "updated_at")

	query, args, err = insertUser.ToSQL()
	if err != nil {
		return nil, err
	}

	var insertedUserID int64
	var createdAt time.Time
	var updatedAt time.Time
	err = tx.QueryRowContext(ctx, query, args...).Scan(&insertedUserID, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// Set the ID on the user and return it
	user.ID = insertedUserID
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return user, nil
}
