package service

import (
	"github.com/p-12s/arch-lab/1-sync-monolit/api/common"
	"github.com/p-12s/arch-lab/1-sync-monolit/api/pkg/repository"
)

type LoyaltyService struct {
	repo repository.Loyalty
}

func NewLoyaltyService(repo repository.Loyalty) *LoyaltyService {
	return &LoyaltyService{repo: repo}
}

func (s *LoyaltyService) CreateUserCard(loyaltyCard common.LoyaltyCard) error {
	return s.repo.CreateUserCard(loyaltyCard)
}
