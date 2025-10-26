package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sqlc "github.com/Anagh3/go-backend/db/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{queries: sqlc.New(db)}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, dob time.Time) error {
	params := sqlc.CreateUserParams{Name: name, Dob: dob}
	_, err := r.queries.CreateUser(ctx, params)
	return err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uint64) (sqlc.User, error) {
	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sqlc.User{}, errors.New("user not found")
		}
		return sqlc.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]sqlc.User, error) {
	return r.queries.GetAllUsers(ctx)
}

func (r *UserRepository) UpdateUser(ctx context.Context, id uint64, name string, dob time.Time) error {
	params := sqlc.UpdateUserParams{ID: id, Name: name, Dob: dob}
	_, err := r.queries.UpdateUser(ctx, params)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uint64) error {
	return r.queries.DeleteUser(ctx, id)
}
