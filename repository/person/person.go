package personrepo

import (
	"Demo-Port-n-Adapter/domain"
	"Demo-Port-n-Adapter/model"
	"database/sql"
	"fmt"
)

type personRepo struct {
	db *sql.DB
}

// ท่าที่ 1 accept interface return struct
// func NewPersonRepo(db *sql.DB) domain.PersonRepo {
// 	return &personRepo{db}
// }

// ท่าที่ 2 accept interface return struct
var _ domain.PersonRepo = (*personRepo)(nil)

func NewPersonRepo(db *sql.DB) *personRepo {
	return &personRepo{db}
}

func (repo *personRepo) GetPersonWithPersonID(personId int) (*model.Person, error) {
	var person model.Person
	queryString := getPersonWithPersonID
	err := repo.db.QueryRow(queryString, personId).Scan(
		&person.PersonID,
		&person.FirstName,
		&person.LastName,
		&person.Address,
		&person.City)
	if err != nil {
		fmt.Println("result error :: ", err)
		return &person, err
	}

	return &person, nil
}
