package main

import (
	"fmt"
	"log"
	"os"
	"time-tracker/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (db *gorm.DB) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.Track{})
}
