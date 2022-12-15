package repository

import (
	"context"
	"database/sql"
)

type AuthRepository struct {
	Db  sql.DB
	Ctx context.Context
}

func NewAuthRepository(db sql.DB, ctx context.Context) *AuthRepository {
	return &AuthRepository{
		Db:  db,
		Ctx: ctx,
	}
}

func (a *AuthRepository) FindAccount(email string) bool {

	return false
}
