package server

import "github.com/gofiber/fiber/v2"

func (server *server) personGroup(app *fiber.App) {

	app.Get("/person/:personId", server.PersonHandle.GetPersonWithPersonID)

}

