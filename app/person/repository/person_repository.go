package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"test-api/app/person/model"
	"test-api/infrastructure/db"
)

type PersonRepository struct {
	*gorm.DB
}

// NewPersonRepository will create an object that represent the Repository interface
func NewPersonRepository(DB *gorm.DB) PersonRepositoryInterface {
	return &PersonRepository{
		DB: db.Get().DB,
	}
}

type PersonRepositoryInterface interface {
	GetAll() ([]*model.Person, error)
	GetById(int) (*model.Person, error)
	Create(string, string) (*model.Person, error)
	Remove(int) error
	Update(int, string, string) (*model.Person, error)
}

// check whether all the methods defined in the interface are implemented
var _ PersonRepositoryInterface = &PersonRepository{}

// GetPersons ...
func (c *PersonRepository) GetAll() ([]*model.Person, error) {
	persons := []*model.Person{}

	if err := c.DB.Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}

// GetById ...
func (c *PersonRepository) GetById(personID int) (*model.Person, error) {
	person := model.Person{}

	if err := c.DB.First(&person, personID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &person, nil
}

// Create ...
func (c *PersonRepository) Create(firstName, lastName string) (*model.Person, error) {

	newPerson := model.Person{
		FirstName: firstName,
		LastName:  lastName,
	}

	err := c.DB.Create(&newPerson).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &newPerson, nil
}

// Remove ...
func (c *PersonRepository) Remove(personID int) error {

	if err := c.DB.Where("id = ?", personID).Delete(&model.Person{}).Error; err != nil {
		return err
	}

	return nil
}

// Update ...
func (c *PersonRepository) Update(personID int, firstName, lastName string) (*model.Person, error) {

	err := c.DB.Table("person").Where("id = ?", personID).Updates(map[string]interface{}{
		"first_name": firstName,
		"last_name":  lastName,
	}).Error

	if err != nil {
		return nil, err
	}

	res, _ := c.GetById(personID)
	return res, nil
}
