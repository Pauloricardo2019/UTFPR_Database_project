package dto

import "utfpr_db/internal/model"

type UpdateUserRequest struct {
	ID        uint64 `gorm:"primaryKey"`
	FirstName string `valid:"required" gorm:"varchar(100)"`
	LastName  string `valid:"required" gorm:"varchar(100)"`
}

func (u *UpdateUserRequest) ParseToUserObject() *model.User {
	return &model.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
