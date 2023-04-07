//go:build wireinject
// +build wireinject

package provider

import (
	"log"
	"sync"

	"deploy-test/config"
	"deploy-test/domain"
	"deploy-test/internal/database"
	svc "deploy-test/internal/service"
	repo "deploy-test/repo/gorm"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbOnce sync.Once

func NewDB() (*gorm.DB, error) {
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

func NewRepo() (domain.Repository, error) {
	panic(wire.Build(repo.NewRepository, NewDB, NewConfig))
}

func NewService() (domain.Service, error) {
	panic(wire.Build(svc.NewService, repo.NewRepository, NewDB, NewConfig))
}
