package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	logToFile("logs")

	loadEnv()

	db := connectDB()

	migrateModels(db)

	app := NewRouter()
	api := app.Group("/api")
	user := api.Group("/user")
	UserRouter(user, db)

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
