package main

import (
	"deploy-test/config"
	model "deploy-test/repo/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func main() {
	config := config.NewConfig()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database, config.DB.Flag)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: connection,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}

	// 建立data

	err = seedTodo(db)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("success")
}

func seedTodo(db *gorm.DB) error {
	var err error

	falseStatus := false
	trueStatus := true

	modelSeeds := []model.Todo{
		{
			ID:      1,
			Title:   "帶狗散步",
			Content: "14:00 帶狗散步",
			IsDone:  &falseStatus,
		},
		{
			ID:      2,
			Title:   "寫作業",
			Content: "寫作業",
			IsDone:  &trueStatus,
		},
	}

	err = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&modelSeeds).Error

	return err
}
