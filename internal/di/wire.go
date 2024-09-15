//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/eisa-auth/db"
	"github.com/K-Kizuku/eisa-auth/internal/app/handler"
	"github.com/K-Kizuku/eisa-auth/internal/app/repository"
	"github.com/K-Kizuku/eisa-auth/internal/app/service"
	"github.com/google/wire"
)

func InitHandler() *handler.Root {
	wire.Build(
		db.New,
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
		handler.New,
	)
	return &handler.Root{}
}
