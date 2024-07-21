package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Output struct {
	Data  any
	Error string
}

func handleErrors(c *fiber.Ctx, err error) error {
	log.Println(err)
	c.Status(http.StatusBadRequest)

	return c.JSON(Output{
		Data:  nil,
		Error: err.Error(),
	})
}
