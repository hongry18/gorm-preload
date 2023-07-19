package domain

import (
	"context"

	"github.com/hongry18/gorm-preload-example/internal/model"
)

type UserUsecase interface {
	Create(context.Context, uint) error
	List(context.Context) error
}

type UserRepository interface {
	Create(context.Context, model.User) error
	List(context.Context) ([]model.User, error)
}
