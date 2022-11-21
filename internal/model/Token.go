package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Token struct {
	Value     string `valid:"notnull" gorm:"varchar(255)"`
	UserID    uint64
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" valid:"notnull"`
	CreatedAt time.Time `valid:"-" gorm:"autoCreateTime"`
	ExpiresAt time.Time `valid:"-" gorm:"autoUpdateTime:milli"`
}

func (Token) TableName() string {
	return "tokens"
}

func (t *Token) Validate() error {
	return validator.New().Struct(t)
}
