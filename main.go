package main

import (
	"log"
	"os"

	_ "time-tracker/docs"

	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {

	logToFile("logs")

	loadEnv()

	db := connectDB()

	migrateModels(db)

	app := NewRouter()
	app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("/api")

	user := api.Group("/user")
	UserRouter(user, db)

	task := api.Group("/task")
	TaskRouter(task, db)

	track := api.Group("/track")
	TrackRouter(track, db)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func logToFile(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
