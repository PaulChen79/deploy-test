package gorm

import (
	"deploy-test/config"
	"deploy-test/domain"

	"gorm.io/gorm"
)

type repository struct {
	db     *gorm.DB
	config *config.Config
}

func NewRepository(db *gorm.DB, config *config.Config) domain.Repository {
	return &repository{
		db:     db,
		config: config,
	}
}
