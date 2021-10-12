package service

import (
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(userId int) (common.User, error) {
	return s.repo.GetById(userId)
}
