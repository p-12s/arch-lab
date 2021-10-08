package service

import (
	"github.com/google/uuid"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/pkg/repository"
)

type User interface {
	CreateUser() (uuid.UUID, error)
	GenerateToken(userCode uuid.UUID) (string, error)
	ParseToken(accessToken string) (uuid.UUID, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
