package service

import (
	"context"
	"utfpr_db/internal/model"
)

type tokenRepository interface {
	Create(ctx context.Context, token *model.Token) error
	GetByValue(ctx context.Context, value string) error
}

type TokenService struct {
	tokenRepository tokenRepository
}

func NewTokenService(tokenRepository tokenRepository) *TokenService {
	return &TokenService{
		tokenRepository: tokenRepository,
	}
}

func (u *TokenService) Create(ctx context.Context, token *model.Token) error {
	return u.tokenRepository.Create(ctx, token)
}

func (u *TokenService) GetByValue(ctx context.Context, value string) error {
	return u.tokenRepository.GetByValue(ctx, value)
}
