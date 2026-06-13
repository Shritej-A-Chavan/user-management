package repository

import (
	"context"
	"go-task/db/sqlc"
	"time"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		queries: q,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	result, err := r.queries.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Name: name,
			Dob:  dob,
		},
	)

	if err != nil {
		return sqlc.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return sqlc.User{}, err
	}

	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int64,
) (sqlc.User, error) {

	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) GetUsers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]sqlc.User, error) {

	return r.queries.GetUsers(
		ctx,
		sqlc.GetUsersParams{
			Limit:  limit,
			Offset: offset,
		},
	)
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (sqlc.User, error) {

	err := r.queries.UpdateUser(
		ctx,
		sqlc.UpdateUserParams{
			Name: name,
			Dob:  dob,
			ID:   id,
		},
	)

	if err != nil {
		return sqlc.User{}, err
	}

	user, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return sqlc.User{}, err
	}

	return user, nil
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int64,
) error {

	_, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	err = r.queries.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
