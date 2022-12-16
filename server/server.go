package server

import (
	"Demo-Port-n-Adapter/domain"
	"database/sql"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"

	personhandle "Demo-Port-n-Adapter/handle/person"
	personrepo "Demo-Port-n-Adapter/repository/person"
	personservice "Demo-Port-n-Adapter/service/person"
)

type (
	server struct {
		PersonHandle domain.PersonHandle
	}

	Server interface {
		Router(app *fiber.App)
		personGroup(app *fiber.App)
	}
)

// ท่าที่ 1 accept interface return struct
func NewServer(db *sql.DB, redis *redis.Client) Server {

	newPersonRepo := personrepo.NewPersonRepo(db)
	newPersonService := personservice.NewPersonService(redis, newPersonRepo)
	newPersonHandle := personhandle.NewPersonHandle(newPersonService)

	return &server{&newPersonHandle}
}

// ท่าที่ 2 accept interface return struct
// var _ Server = (*server)(nil)

// func NewServer(db *sql.DB, redis *redis.Client) *server {
// 	newPersonRepo := personrepo.NewPersonRepo(db)
// 	newPersonService := personservice.NewPersonService(redis, newPersonRepo)
// 	newPersonHandle := personhandle.NewPersonHandle(newPersonService)

// 	return &server{&newPersonHandle}
// }
