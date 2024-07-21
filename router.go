package main

import (
	"time-tracker/handler"

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
	app.Get("/labourCosts", handler.LabourCosts(db))
	app.Post("/createUser", handler.CreateUser(db))
	app.Patch("/updateUser", handler.UpdateUser(db))
	app.Delete("/deleteUser", handler.DeliteUser(db))
}

func TaskRouter(app fiber.Router, db *gorm.DB) {
	app.Post("/createTask", handler.CreateTask(db))
	app.Patch("/updateTask", handler.UpdateTask(db))
	app.Delete("/deleteTask", handler.DeliteTask(db))
}

func TrackRouter(app fiber.Router, db *gorm.DB) {
	app.Post("/createTrack", handler.CreateTrack(db))
	app.Patch("/updateTrack", handler.UpdateTrack(db))
	app.Delete("/deleteTrack", handler.DeliteTrack(db))
}
