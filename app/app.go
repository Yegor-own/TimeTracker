package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=admin password=root dbname=time-tracker port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	app := NewRouter()
	api := app.Group("/api")
	user := api.Group("/user")
	UserRouter(user, db)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
