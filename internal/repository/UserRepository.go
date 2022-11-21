package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"utfpr_db/internal/model"
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	baseRepo := NewBaseRepository(db)
	return &UserRepository{
		baseRepo,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	conn, err := u.getConnection(ctx)
	if err != nil {
		return nil, err
	}

	err = conn.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetByID(ctx context.Context, id uint64) (bool, *model.User, error) {
	coon, err := u.getConnection(ctx)
	if err != nil {
		return false, nil, err
	}

	user := &model.User{}

	if err = coon.Where(&model.User{
		ID: id,
	}).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, user, err
		}
		return false, nil, err
	}

	return true, user, nil
}

func (u *UserRepository) Update(ctx context.Context, user *model.User) error {
	coon, err := u.getConnection(ctx)
	if err != nil {
		return err
	}

	if err = coon.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Delete(ctx context.Context, id uint64) error {
	coon, err := u.getConnection(ctx)
	if err != nil {
		return err
	}

	user := &model.User{}

	if err = coon.Where(&model.User{
		ID: id,
	}).Delete(user).Error; err != nil {
		return err
	}

	return nil
}
