package repository

import (
	"context"
	"gorm.io/gorm"
	"utfpr_db/internal/model"
)

type TokenRepository struct {
	*BaseRepository
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	baseRepo := NewBaseRepository(db)
	return &TokenRepository{
		baseRepo,
	}
}

func (u *TokenRepository) Create(ctx context.Context, token *model.Token) error {
	conn, err := u.getConnection(ctx)
	if err != nil {
		return err
	}

	return conn.Create(token).Error
}

func (u *TokenRepository) GetByValue(ctx context.Context, value string) error {
	conn, err := u.getConnection(ctx)
	if err != nil {
		return err
	}

	return conn.Where(model.Token{
		Value: value,
	}).Error
}
