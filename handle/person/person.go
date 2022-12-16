package personhandle

import (
	"Demo-Port-n-Adapter/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type personHandle struct {
	PersonService domain.PersonService
}

// ท่าที่ 1 accept interface return struct
// func NewPersonHandle(PersonService domain.PersonService) personHandle {
// 	return personHandle{PersonService: PersonService}
// }

// ท่าที่ 2 accept interface return struct
var _ domain.PersonHandle = (*personHandle)(nil)

func NewPersonHandle(PersonService domain.PersonService) personHandle {
	return personHandle{PersonService: PersonService}
}

func (handle *personHandle) GetPersonWithPersonID(c *fiber.Ctx) error {

	personId, _ := strconv.Atoi(c.Params("personId"))
	response, err := handle.PersonService.GetPersonWithPersonID(personId)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
