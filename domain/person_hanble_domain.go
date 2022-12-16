package domain

import "github.com/gofiber/fiber/v2"

type PersonHandle interface {
	GetPersonWithPersonID(c *fiber.Ctx) error
}
