package usecase

import (
	"test-api/app/person/model"
	"test-api/app/person/repository"
)

type PersonUsecase interface {
	GetAll() ([]*model.Person, error)
	GetById(int) (*model.Person, error)
	Create(string, string) (*model.Person, error)
	Remove(int) error
	Update(int, string, string) (*model.Person, error)
}

// check whether all the methods defined in the interface are implemented
var _ PersonUsecase = &personUsecase{}

type personUsecase struct {
	repo repository.PersonRepositoryInterface
}

func NewPersonUsecase(repo repository.PersonRepositoryInterface) PersonUsecase {
	return &personUsecase{
		repo: repo,
	}
}

func (o *personUsecase) GetAll() ([]*model.Person, error) {
	persons, err := o.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (o *personUsecase) GetById(id int) (*model.Person, error) {
	person, err := o.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (o *personUsecase) Create(firstName, lastName string) (*model.Person, error) {
	person, err := o.repo.Create(firstName, lastName)

	if err != nil {
		return nil, err
	}
	return person, nil
}

func (o *personUsecase) Update(id int, firstName, lastName string) (*model.Person, error) {
	person, err := o.repo.Update(id, firstName, lastName)

	if err != nil {
		return nil, err
	}
	return person, nil
}

func (o *personUsecase) Remove(id int) error {
	err := o.repo.Remove(id)

	if err != nil {
		return err
	}
	return nil
}
