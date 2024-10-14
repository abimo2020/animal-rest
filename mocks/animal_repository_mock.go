package mocks

import (
	e "animal-rest/errors"
	"animal-rest/models"

	"github.com/stretchr/testify/mock"
)

type AnimalRepositoryMock struct {
	mock.Mock
}

func (m *AnimalRepositoryMock) Save(req *models.Animal) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *AnimalRepositoryMock) Find() ([]models.Animal, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, e.ErrAnimalNotFound
	}
	return args.Get(0).([]models.Animal), args.Error(1)
}
func (m *AnimalRepositoryMock) FindById(id uint) (*models.Animal, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, e.ErrAnimalNotFound
	}
	return args.Get(0).(*models.Animal), args.Error(1)
}
func (m *AnimalRepositoryMock) FindByName(name string) (*models.Animal, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, e.ErrAnimalNotFound
	}
	return args.Get(0).(*models.Animal), args.Error(1)
}

func (m *AnimalRepositoryMock) DeleteById(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
