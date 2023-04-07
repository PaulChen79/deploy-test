package main

import (
	"deploy-test/internal/provider"
	model "deploy-test/repo/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	config := provider.NewConfig()

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/?%s", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Flag)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connection,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	// 方便重置的arg
	if len(os.Args) >= 2 {
		if arg := os.Args[1]; arg == "--reset" {
			db.Exec("DROP DATABASE " + config.DB.Database)
		}
	}

	// 建立DB
	db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DB.Database + " CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")

	// 建立table
	connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database, config.DB.Flag)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: connection,
	}), &gorm.Config{})

	db.AutoMigrate(&model.Todo{})
}
