package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"utfpr_db/internal/model"
)

type (
	UserServiceMock struct {
		mock.Mock
	}
)

func (m *UserServiceMock) Create(ctx context.Context, user *model.User) (*model.User, error) {
	args := m.Called(ctx, user)
	resultUser := args.Get(0).(*model.User)

	if args.Get(1) != nil {
		err := args.Get(1).(error)
		return resultUser, err
	}
	return resultUser, nil
}

func (m *UserServiceMock) GetByID(ctx context.Context, id uint64) (bool, *model.User, error) {
	args := m.Called(ctx, id)

	found := args.Get(0).(bool)

	var resultUser *model.User
	if args.Get(1) != nil {
		resultUser = args.Get(1).(*model.User)
	}

	var err error
	if args.Get(2) != nil {
		err = args.Get(2).(error)
	}

	return found, resultUser, err
}

func (m *UserServiceMock) Update(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)

	var err error
	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}

func (m *UserServiceMock) Delete(ctx context.Context, id uint64) error {
	args := m.Called(ctx, id)

	var err error
	if args.Get(0) != nil {
		err = args.Get(1).(error)
	}

	return err
}
