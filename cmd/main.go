package main

import (
	"go.uber.org/fx"

	"github.com/hongry18/gorm-preload-example/internal/app"
	"github.com/hongry18/gorm-preload-example/internal/model"
	gormPkg "github.com/hongry18/gorm-preload-example/internal/pkg/gorm"
	userMysqlRepository "github.com/hongry18/gorm-preload-example/internal/user/repository/mariadb"
	userUsecase "github.com/hongry18/gorm-preload-example/internal/user/usecase"
)

func main() {
	fx.New(
		fx.Provide(
			gormPkg.New,
			userMysqlRepository.New,
			userUsecase.New,
		),
		fx.Invoke(
			model.Migration,
			app.Run,
		),
	).Run()
}
