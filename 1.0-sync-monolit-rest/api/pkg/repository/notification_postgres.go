package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

type NotificationPostgres struct {
	db *sqlx.DB
}

func NewNotificationPostgres(db *sqlx.DB) *NotificationPostgres {
	return &NotificationPostgres{db: db}
}

func (r *NotificationPostgres) SaveNotification(userId int, notification common.NotificationSendReq) error {
	// здесь можно и без транзакции, но пусть будет для примера,
	// вдруг создание будет в несколько таблиц
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, notification_program_type, domain_event_type, status) values ($1, $2, $3, $4)", notificationTable)
	_, err = r.db.Exec(query, userId, common.NOTIFICATION_PROGRAM_EMAIL, common.DOMAIN_EVENT_LOYALTY_CARD_CREATE, common.NOTIFICATION_SEND_STATUS_OK)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return tx.Commit()
}
