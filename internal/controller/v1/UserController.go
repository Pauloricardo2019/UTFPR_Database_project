package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"utfpr_db/dto"
)

type userFacade interface {
	Create(ctx context.Context, createUserRequestDTO *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetByID(ctx context.Context, id uint64) (bool, *dto.CreateUserResponse, error)
	Update(ctx context.Context, updateUserRequestDTO *dto.UpdateUserRequest) error
	Delete(ctx context.Context, id uint64) error
}

type userController struct {
	userFacade userFacade
}

func NewUserController(userFacade userFacade) *userController {
	return &userController{
		userFacade: userFacade,
	}
}

// Create - create a user objects, and returns a user objects
// @Summary - Create user
// @Description - Create a user
// @Tags - User
// @Accept json
// @Produce json
// @Param User body dto.CreateUserRequest true "User to be created"
// @Success 201 {object} dto.CreateUserResponse
// @Error 400 {object} dto.Error
// @Error 500 {object} dto.Error
// @Router /v1/user [post]
// @Security ApiKeyAuth
func (u *userController) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute*15)
	defer cancel()

	createUserRequestDTO := &dto.CreateUserRequest{}
	err := c.BindJSON(createUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Error parsing json: " + err.Error()})
		return
	}

	createUserResponseDTO, err := u.userFacade.Create(ctx, createUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot create a user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createUserResponseDTO)
}

// GetByID - create a user objects, and returns a user objects
// @Summary - Create user
// @Description - Create a user
// @Tags - User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} dto.CreateUserResponse
// @Error 400 {object} dto.Error
// @Error 500 {object} dto.Error
// @Router /v1/user/{id} [get]
// @Security ApiKeyAuth
func (u *userController) GetByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute*15)
	defer cancel()

	userID := c.Param("id")

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	found, createUserResponseDTO, err := u.userFacade.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot create a user: " + err.Error()})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createUserResponseDTO)
}

// Update - create a user objects, and returns a user objects
// @Summary - Create user
// @Description - Create a user
// @Tags - User
// @Accept json
// @Produce json
// @Param User body dto.UpdateUserRequest true "User to be created"
// @Success 204
// @Error 400 {object} dto.Error
// @Error 500 {object} dto.Error
// @Router /v1/user [put]
// @Security ApiKeyAuth
func (u *userController) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute*15)
	defer cancel()

	updateUserRequestDTO := &dto.UpdateUserRequest{}
	err := c.BindJSON(updateUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Error parsing json: " + err.Error()})
		return
	}

	err = u.userFacade.Update(ctx, updateUserRequestDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot create a user: " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Delete - create a user objects, and returns a user objects
// @Summary - Create user
// @Description - Create a user
// @Tags - User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204
// @Error 400 {object} dto.Error
// @Error 500 {object} dto.Error
// @Router /v1/user/{id} [delete]
// @Security ApiKeyAuth
func (u *userController) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*15)
	defer cancel()

	userID := c.Param("id")

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	err = u.userFacade.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Error{Message: "Cannot create a user: " + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
