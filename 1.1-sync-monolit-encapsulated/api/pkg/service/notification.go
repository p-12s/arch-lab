package service

import (
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/pkg/repository"
)

type NotificationService struct {
	repo repository.Notification
}

func NewNotificationService(repo repository.Notification) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) SaveNotification(userId int, notification common.NotificationSendReq) error {
	return s.repo.SaveNotification(userId, notification)
}
