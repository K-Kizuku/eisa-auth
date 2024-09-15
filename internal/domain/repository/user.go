package repository

import (
	"context"

	"github.com/K-Kizuku/eisa-auth/internal/domain/entity"
)

type IUserRepository interface {
	FindUserByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user entity.User) error
	UpdatePassword(ctx context.Context, id, password string) error
	UpdateEisaFile(ctx context.Context, id, eisaFile string) error
}
