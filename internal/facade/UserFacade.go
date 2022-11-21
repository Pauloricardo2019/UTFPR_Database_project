package facade

import (
	"context"
	"crypto/sha256"
	"fmt"
	"utfpr_db/dto"
	"utfpr_db/internal/model"
)

type userService interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id uint64) (bool, *model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
}

type UserFacade struct {
	userService userService
}

func NewUserFacade(userService userService) *UserFacade {
	return &UserFacade{
		userService: userService,
	}
}

func (u *UserFacade) Create(ctx context.Context, createUserRequestDTO *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	user := createUserRequestDTO.ParseToUserObject()

	sum := sha256.Sum256([]byte(user.Password))
	user.HashedPassword = fmt.Sprintf("%x", sum)
	user.Password = ""

	createdUser, err := u.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	createUserResponse := &dto.CreateUserResponse{}
	createUserResponse.ParseFromUserObject(createdUser)

	return createUserResponse, nil
}

func (u *UserFacade) GetByID(ctx context.Context, id uint64) (bool, *dto.CreateUserResponse, error) {

	found, user, err := u.userService.GetByID(ctx, id)
	if err != nil {
		return false, nil, err
	}

	createUserResponse := &dto.CreateUserResponse{}
	createUserResponse.ParseFromUserObject(user)

	if !found {
		return false, nil, err
	}

	return true, createUserResponse, nil
}

func (u *UserFacade) Update(ctx context.Context, updateUserRequestDTO *dto.UpdateUserRequest) error {

	user := updateUserRequestDTO.ParseToUserObject()

	if err := u.userService.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserFacade) Delete(ctx context.Context, id uint64) error {

	if err := u.userService.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
