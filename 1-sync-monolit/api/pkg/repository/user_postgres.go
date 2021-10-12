package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

// UserPostgres - репозиторий
type UserPostgres struct {
	db *sqlx.DB
}

// NewUserPostgres - конструктор объекта репозитория
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// GetUser - получение пользователя из БД
func (r *UserPostgres) GetById(userId int) (common.User, error) {
	var user common.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)

	err := r.db.Get(&user, query, userId)

	return user, err
}
