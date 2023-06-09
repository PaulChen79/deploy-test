// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package provider

import (
	"deploy-test/config"
	"deploy-test/domain"
	"deploy-test/internal/database"
	"deploy-test/internal/service"
	"deploy-test/repo/gorm"
	gorm2 "gorm.io/gorm"
	"log"
	"sync"
)

// Injectors from wire.go:

func NewRepo() (domain.Repository, error) {
	gormDB, err := NewDB()
	if err != nil {
		return nil, err
	}
	config := NewConfig()
	repository := gorm.NewRepository(gormDB, config)
	return repository, nil
}

func NewService() (domain.Service, error) {
	gormDB, err := NewDB()
	if err != nil {
		return nil, err
	}
	config := NewConfig()
	repository := gorm.NewRepository(gormDB, config)
	domainService := service.NewService(repository, config)
	return domainService, nil
}

// wire.go:

var db *gorm2.DB

var dbOnce sync.Once

func NewDB() (*gorm2.DB, error) {
	var err error
	if db == nil {
		dbOnce.Do(func() {
			log.Println("connect db")
			db, err = database.DatabaseConnection(NewConfig().DB)
			if err != nil {
				return
			}
		})
	}
	return db, err
}

var cg *config.Config

var configOnce sync.Once

func NewConfig() *config.Config {
	configOnce.Do(func() {
		cg = config.NewConfig()
	})
	return cg
}
