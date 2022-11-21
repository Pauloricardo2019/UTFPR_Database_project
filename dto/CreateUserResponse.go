package dto

import (
	"utfpr_db/internal/model"
)

type CreateUserResponse struct {
	ID uint64 `json:"id"`
}

func (r *CreateUserResponse) ParseFromUserObject(user *model.User) {
	r.ID = user.ID
}
