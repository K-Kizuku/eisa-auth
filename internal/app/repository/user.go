package repository

import (
	"context"

	"github.com/K-Kizuku/eisa-auth/db/sql/query"
	"github.com/K-Kizuku/eisa-auth/internal/domain/entity"
	"github.com/K-Kizuku/eisa-auth/internal/domain/repository"
)

type UserRepository struct {
	queries *query.Queries
}

func NewUserRepository(queries *query.Queries) repository.IUserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := r.queries.GetUserByIDWithEisaFiles(ctx, id)
	if err != nil {
		return nil, err
	}
	e := &entity.User{
		ID:       user.UserID,
		Username: user.Name,
		Password: user.HashedPassword,
		Email:    user.Mail,
		EisaFile: user.FilePath,
	}
	return e, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	e := &entity.User{
		ID:       user.UserID,
		Username: user.Name,
		Password: user.HashedPassword,
		Email:    user.Mail,
	}
	return e, nil
}

func (r *UserRepository) Create(ctx context.Context, user entity.User) error {
	_, err := r.queries.CreateUser(ctx, query.CreateUserParams{
		UserID:         user.ID,
		Mail:           user.Email,
		Name:           user.Username,
		HashedPassword: user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id, password string) error {
	err := r.queries.UpdatePassword(ctx, query.UpdatePasswordParams{
		UserID:         id,
		HashedPassword: password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateEisaFile(ctx context.Context, id, eisaFile string) error {
	err := r.queries.UpsertEisaFile(ctx, query.UpsertEisaFileParams{
		UserID:   id,
		FilePath: eisaFile,
	})
	if err != nil {
		return err
	}
	return nil
}
