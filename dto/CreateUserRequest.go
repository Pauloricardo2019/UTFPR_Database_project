package dto

import "utfpr_db/internal/model"

type CreateUserRequest struct {
	FirstName string `valid:"required" gorm:"varchar(100)"`
	LastName  string `valid:"required" gorm:"varchar(100)"`
	Email     string `valid:"required" gorm:"varchar(255)"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

func (r *CreateUserRequest) ParseToUserObject() *model.User {
	return &model.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Login:     r.Login,
		Password:  r.Password,
	}
}
