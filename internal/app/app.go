package app

import (
	"context"

	"github.com/hongry18/gorm-preload-example/internal/domain"
	"go.uber.org/fx"
)

func Run(shutdowner fx.Shutdowner, userUsecase domain.UserUsecase) {
	ctx := context.Background()
	userUsecase.Create(ctx, 1)
	userUsecase.Create(ctx, 2)
	userUsecase.Create(ctx, 3)
	userUsecase.Create(ctx, 4)
	userUsecase.List(ctx)

	shutdowner.Shutdown()
}
