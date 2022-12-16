package domain

import "Demo-Port-n-Adapter/model"

type PersonService interface {
	GetPersonWithPersonID(personId int) (*model.Person, error)
}
