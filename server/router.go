package server

import (
	"github.com/gofiber/fiber/v2"
)

func (server *server) Router(app *fiber.App) {

	server.personGroup(app)

}
