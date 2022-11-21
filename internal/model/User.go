package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID             uint64    `gorm:"primaryKey"`
	FirstName      string    `valid:"required" gorm:"varchar(100)"`
	LastName       string    `valid:"required" gorm:"varchar(100)"`
	Email          string    `valid:"required" gorm:"varchar(255)"`
	DDD            string    `valid:"-" gorm:"varchar(2)"`
	Phone          string    `valid:"-" gorm:"varchar(9)"`
	Login          string    `valid:"required" gorm:"varchar(100)"`
	Password       string    `valid:"required" gorm:"-"`
	HashedPassword string    `valid:"required" gorm:"varchar(64)"`
	CreatedAt      time.Time `valid:"-" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `valid:"-" gorm:"autoUpdateTime:milli"`
	LastLogin      time.Time `valid:"-"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}
