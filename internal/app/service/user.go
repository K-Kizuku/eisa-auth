package service

import (
	"context"
	"errors"

	"github.com/K-Kizuku/eisa-auth/internal/domain/entity"
	"github.com/K-Kizuku/eisa-auth/internal/domain/repository"
	"github.com/K-Kizuku/eisa-auth/pkg/hash"
	"github.com/K-Kizuku/eisa-auth/pkg/jwt"
	"github.com/K-Kizuku/eisa-auth/pkg/middleware"
)

type IUserService interface {
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user entity.User) error
	UpdatePassword(ctx context.Context, id, password string) error
	UpdateEisaFile(ctx context.Context, id, eisaFile string) error
	CheckID(ctx context.Context, id string) error
	VerifyPassword(ctx context.Context, email, password string) (string, error)
	GenerateJWT(ctx context.Context, id string) (string, error)
	GenerateSignedURL(ctx context.Context, id string) (string, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindUserByID(ctx context.Context, id string) (*entity.User, error) {
	return s.repo.FindUserByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, user entity.User) error {
	u := entity.User{
		ID:       user.ID,
		Username: user.Username,
		Password: hash.EncryptPassword(user.Password),
		Email:    user.Email,
	}
	return s.repo.Create(ctx, u)
}

func (s *UserService) UpdatePassword(ctx context.Context, id, password string) error {
	return s.repo.UpdatePassword(ctx, id, password)
}

func (s *UserService) UpdateEisaFile(ctx context.Context, id, eisaFile string) error {
	return s.repo.UpdateEisaFile(ctx, id, eisaFile)
}

func (s *UserService) CheckID(ctx context.Context, id string) error {
	tokenID := ctx.Value(middleware.UserIDKey).(string)
	if id != tokenID {
		return errors.New("user id is not match")
	}
	return nil
}

func (s *UserService) VerifyPassword(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if err := hash.CompareHashPassword(user.Password, password); err != nil {
		return "", err
	}

	return user.ID, nil
}

func (s *UserService) GenerateJWT(ctx context.Context, id string) (string, error) {
	jwt, err := jwt.GenerateToken(id)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (s *UserService) GenerateSignedURL(ctx context.Context, id string) (string, error) {
	return "", nil
}
