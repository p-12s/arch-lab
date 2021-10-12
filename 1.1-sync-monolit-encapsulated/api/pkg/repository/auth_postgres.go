package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

// AuthPostgres - репозиторий
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres - конструктор объекта репозитория
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser - создание пользователя
func (r *AuthPostgres) CreateUser(user common.User) (int, error) {
	// здесь можно и без транзакции, но пусть будет для примера,
	// вдруг создание будет в несколько таблиц
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, first_name, last_name, email, password, phone, address) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.Address)
	if err := row.Scan(&id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	return id, tx.Commit()
}

// GetUser - получение пользователя из БД
func (r *AuthPostgres) GetUser(username, passwordHash string) (common.User, error) {

	var user common.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
