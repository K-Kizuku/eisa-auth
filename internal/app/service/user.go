package service

import (
	"context"
	nomal_errors "errors"
	"net/http"

	"github.com/K-Kizuku/eisa-auth/internal/domain/entity"
	"github.com/K-Kizuku/eisa-auth/internal/domain/repository"
	"github.com/K-Kizuku/eisa-auth/pkg/errors"
	"github.com/K-Kizuku/eisa-auth/pkg/hash"
	"github.com/K-Kizuku/eisa-auth/pkg/jwt"
	"github.com/K-Kizuku/eisa-auth/pkg/middleware"
	"github.com/K-Kizuku/eisa-auth/pkg/uuid"
)

type IUserService interface {
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user entity.User) (*entity.User, error)
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

func (s *UserService) Create(ctx context.Context, user entity.User) (*entity.User, error) {
	u := entity.User{
		ID:       uuid.New(),
		Username: user.Username,
		Password: hash.EncryptPassword(user.Password),
		Email:    user.Email,
	}
	createdUser, err := s.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
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
		return errors.New(http.StatusForbidden, nomal_errors.New("token id and request id are different"))
	}
	return nil
}

func (s *UserService) VerifyPassword(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if err := hash.CompareHashPassword(user.Password, password); err != nil {
		return "", errors.New(http.StatusUnauthorized, nomal_errors.New("password is incorrect"))
	}
	return user.ID, nil
}

func (s *UserService) GenerateJWT(ctx context.Context, id string) (string, error) {
	jwt, err := jwt.GenerateToken(id)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err)
	}
	return jwt, nil
}

func (s *UserService) GenerateSignedURL(ctx context.Context, id string) (string, error) {
	return "", nil
}
