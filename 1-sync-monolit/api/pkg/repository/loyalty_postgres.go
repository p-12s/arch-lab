package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

type LoyaltyPostgres struct {
	db *sqlx.DB
}

func NewLoyaltyPostgres(db *sqlx.DB) *LoyaltyPostgres {
	return &LoyaltyPostgres{db: db}
}

func (r *LoyaltyPostgres) CreateUserCard(loyaltyCard common.LoyaltyCard) error {
	// здесь можно и без транзакции, но пусть будет для примера,
	// вдруг создание будет в несколько таблиц
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, code, status) values ($1, $2, $3)", loyaltyCardTable)
	_, err = r.db.Exec(query, loyaltyCard.UserId, loyaltyCard.Code, loyaltyCard.Status)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return tx.Commit()
}
