package main

import "gorm.io/gorm"

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
