package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hongry18/gorm-preload-example/internal/domain"
	"github.com/hongry18/gorm-preload-example/internal/model"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func New(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo}
}

func (u userUsecase) Create(ctx context.Context, idx uint) error {
	var user = model.User{
		Name: fmt.Sprintf("[%d] - name", idx),
		Orders: model.Orders{
			{
				Price: (float64(idx) * 100) + 1,
			},
			{
				Price: (float64(idx) * 100) + 2,
			},
			{
				Price: (float64(idx) * 100) + 3,
			},
			{
				Price: (float64(idx) * 100) + 4,
			},
			{
				Price: (float64(idx) * 100) + 5,
			},
		},
	}
	return u.userRepo.Create(ctx, user)
}

func (u userUsecase) List(ctx context.Context) error {
	list, err := u.userRepo.List(ctx)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(list, "", "  ")
	fmt.Println(string(b))
	return nil
}
