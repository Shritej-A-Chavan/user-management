package service

import (
	"context"
	"errors"
	"go-task/db/sqlc"
	"go-task/internal/repository"
	"time"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob string,
) (sqlc.User, error) {

	parsedDOB, err := time.Parse(
		"2006-01-02",
		dob,
	)

	if err != nil {
		return sqlc.User{}, errors.New(
			"invalid date format, use YYYY-MM-DD",
		)
	}

	if parsedDOB.After(time.Now()) {
		return sqlc.User{}, errors.New(
			"date of birth cannot be in the future",
		)
	}

	return s.repo.CreateUser(
		ctx,
		name,
		parsedDOB,
	)
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int64,
) (sqlc.User, error) {

	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) GetUsers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]sqlc.User, error) {

	return s.repo.GetUsers(
		ctx,
		limit,
		offset,
	)
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob string,
) (sqlc.User, error) {

	parsedDOB, err := time.Parse(
		"2006-01-02",
		dob,
	)

	if err != nil {
		return sqlc.User{}, errors.New(
			"invalid date format, use YYYY-MM-DD",
		)
	}

	if parsedDOB.After(time.Now()) {
		return sqlc.User{}, errors.New(
			"date of birth cannot be in the future",
		)
	}

	return s.repo.UpdateUser(
		ctx,
		id,
		name,
		parsedDOB,
	)
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int64,
) error {
	return s.repo.DeleteUser(ctx, id)
}
