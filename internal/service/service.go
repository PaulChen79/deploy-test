package service

import (
	"deploy-test/config"
	"deploy-test/domain"
)

type service struct {
	repo   domain.Repository
	config *config.Config
}

func NewService(repo domain.Repository, config *config.Config) domain.Service {
	return &service{
		repo:   repo,
		config: config,
	}
}
