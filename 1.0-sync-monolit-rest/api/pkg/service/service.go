package service

import (
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/pkg/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user common.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	User
	Loyalty
	Notification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Loyalty:       NewLoyaltyService(repos.Loyalty),
		Notification:  NewNotificationService(repos.Notification),
	}
}
