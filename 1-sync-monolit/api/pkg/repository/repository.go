package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

// User - интерфейс для работы с пользователем
type User interface {
	CreateUser(userCode uuid.UUID) error
	GetUser(code uuid.UUID) (common.User, error)
}

// Repository - репозиторий
type Repository struct {
	User
}

// NewRepository - конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
