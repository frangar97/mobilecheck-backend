package service

import "github.com/frangar97/mobilecheck-backend/internal/repository"

type Service struct {
	repositories *repository.Repository
}

func NewServices(repositories *repository.Repository) *Service {
	return &Service{
		repositories: repositories,
	}
}
