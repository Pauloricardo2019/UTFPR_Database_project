package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"utfpr_db/config"
)

func TestBeginTransaction(t *testing.T) {
	cfg := config.GetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		panic(err.Error())
	}

	baseRepo := NewBaseRepository(db)

	ctx := context.Background()
	_, err = baseRepo.BeginTransaction(ctx)
	assert.NoError(t, err)
}

func TestCommitTransaction_OK(t *testing.T) {
	cfg := config.GetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		panic(err.Error())
	}

	baseRepo := NewBaseRepository(db)

	ctx := context.Background()
	newCtx, err := baseRepo.BeginTransaction(ctx)
	assert.NoError(t, err)

	err = baseRepo.CommitTransaction(newCtx)
	assert.NoError(t, err)
}

func TestCommitTransaction_NoBeginTransaction(t *testing.T) {
	cfg := config.GetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		panic(err.Error())
	}

	baseRepo := NewBaseRepository(db)

	ctx := context.Background()
	err = baseRepo.CommitTransaction(ctx)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "no transaction found on current context")
}
