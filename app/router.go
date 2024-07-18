package main

import (
	"time-tracker/app/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func NewRouter() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(), logger.New())

	return app
}

func UserRouter(app fiber.Router, db *gorm.DB) {
	app.Post("/createUser", handler.CreateUser(db))
	app.Get("/getUser")
	app.Patch("/updateUser")
}
