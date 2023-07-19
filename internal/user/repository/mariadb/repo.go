package mariadb

import (
	"context"

	"gorm.io/gorm"

	"github.com/hongry18/gorm-preload-example/internal/domain"
	"github.com/hongry18/gorm-preload-example/internal/model"
)

type userMariadbRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserRepository {
	return &userMariadbRepository{db}
}

func (r userMariadbRepository) Create(ctx context.Context, user model.User) error {
	return r.db.Create(&user).Error
}

func (r userMariadbRepository) List(ctx context.Context) ([]model.User, error) {
	var list []model.User
	err := r.db.Preload("Orders").Find(&list).Error
	return list, err
}
