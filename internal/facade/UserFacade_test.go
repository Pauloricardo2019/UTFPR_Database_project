package facade

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"utfpr_db/dto"
	"utfpr_db/internal/mocks"
	"utfpr_db/internal/model"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	createUserRequestDTO := &dto.CreateUserRequest{
		Login:    "mock_login",
		Password: "mock_password",
	}

	userServiceMock := &mocks.UserServiceMock{}
	userServiceMock.On("Create", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			receivedUser := args.Get(1).(*model.User)
			assert.NotNil(t, receivedUser)
			assert.Equal(t, "mock_login", receivedUser.Login)
			assert.Equal(t, "mock_password", receivedUser.Password)
		}).
		Return(&model.User{ID: 33555}, nil)
	userFacade := NewUserFacade(userServiceMock)

	createUserResponse, err := userFacade.Create(ctx, createUserRequestDTO)
	assert.NoError(t, err)
	assert.NotNil(t, createUserResponse)
	assert.Equal(t, uint64(33555), createUserResponse.ID)
}

func TestGetByID(t *testing.T) {
	ctx := context.Background()

	var id uint64 = 1

	userServiceMock := &mocks.UserServiceMock{}
	userServiceMock.On("GetByID", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			userIDReceived := args.Get(1).(uint64)
			assert.Equal(t, id, userIDReceived)
		}).
		Return(true, &model.User{
			ID: id,
		}, nil)
	userFacade := NewUserFacade(userServiceMock)

	found, createUserResponse, err := userFacade.GetByID(ctx, id)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.NotNil(t, createUserResponse)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()

	var id uint64 = 6

	updateDTO := &dto.UpdateUserRequest{
		ID:        id,
		FirstName: "updateUser1",
		LastName:  "updateUser2",
	}

	userServiceMock := &mocks.UserServiceMock{}
	userServiceMock.On("Update", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			receivedUser := args.Get(1).(*model.User)
			assert.NotNil(t, receivedUser)
			assert.Equal(t, id, receivedUser.ID)
			assert.Equal(t, "updateUser1", receivedUser.FirstName)
			assert.Equal(t, "updateUser2", receivedUser.LastName)
		}).
		Return(nil)
	userFacade := NewUserFacade(userServiceMock)

	err := userFacade.Update(ctx, updateDTO)
	assert.NoError(t, err)

}

func TestDelete(t *testing.T) {
	ctx := context.Background()

	var id uint64 = 6

	userServiceMock := &mocks.UserServiceMock{}
	userServiceMock.On("Delete", ctx, mock.Anything).
		Run(func(args mock.Arguments) {
			receivedUserID := args.Get(1)
			assert.Equal(t, id, receivedUserID)
		}).
		Return(nil)
	userFacade := NewUserFacade(userServiceMock)

	err := userFacade.Delete(ctx, id)
	assert.NoError(t, err)
}
