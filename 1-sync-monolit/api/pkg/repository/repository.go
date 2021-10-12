package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
)

// Authorization - интерфейс для регистрации/авторизации
type Authorization interface {
	CreateUser(user common.User) (int, error)
	GetUser(username, passwordHash string) (common.User, error)
}

type User interface {
	GetById(userId int) (common.User, error)
}

type Loyalty interface {
	CreateUserCard(loyaltyCard common.LoyaltyCard) error
}

type Notification interface {
	SaveNotification(userId int, notification common.NotificationSendReq) error
}

// Repository - репозиторий
type Repository struct {
	Authorization
	User
	Loyalty
	Notification
}

// NewRepository - конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Loyalty:       NewLoyaltyPostgres(db),
		Notification:  NewNotificationPostgres(db),
	}
}
