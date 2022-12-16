package personservice

import (
	"Demo-Port-n-Adapter/domain"
	"Demo-Port-n-Adapter/model"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type personService struct {
	Redis      *redis.Client
	PersonRepo domain.PersonRepo
}

// ท่าที่ 1 accept interface return struct
// func NewPersonService(Redis *redis.Client, PersonRepo domain.PersonRepo) domain.PersonService {
// 	return &personService{
// 		Redis:      Redis,
// 		PersonRepo: PersonRepo,
// 	}
// }

// ท่าที่ 2 accept interface return struct
var _ domain.PersonService = (*personService)(nil)
func NewPersonService(Redis *redis.Client, PersonRepo domain.PersonRepo) *personService {
	return &personService{
		Redis:      Redis,
		PersonRepo: PersonRepo,
	}
}

func (service *personService) GetPersonWithPersonID(personId int) (*model.Person, error) {

	person, err := service.PersonRepo.GetPersonWithPersonID(personId)
	if err != nil {
		return nil, fmt.Errorf("personID not found")
	}

	return person, nil

}
