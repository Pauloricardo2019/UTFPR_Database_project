package service

import (
	"context"
	"utfpr_db/internal/model"
)

type userRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id uint64) (bool, *model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
}

type UserService struct {
	userRepository userRepository
}

func NewUserService(userRepository userRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return u.userRepository.Create(ctx, user)
}

func (u *UserService) GetByID(ctx context.Context, id uint64) (bool, *model.User, error) {
	return u.userRepository.GetByID(ctx, id)
}

func (u *UserService) Update(ctx context.Context, user *model.User) error {
	return u.userRepository.Update(ctx, user)
}

func (u *UserService) Delete(ctx context.Context, id uint64) error {
	return u.userRepository.Delete(ctx, id)
}
