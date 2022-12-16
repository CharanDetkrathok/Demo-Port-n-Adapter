package domain

import "Demo-Port-n-Adapter/model"

type PersonRepo interface {
	GetPersonWithPersonID(personId int) (*model.Person, error)
}
