package mocks

import (
	e "animal-rest/errors"
	"animal-rest/payload"

	"github.com/stretchr/testify/mock"
)

type AnimalUsecaseMock struct {
	mock.Mock
}

func (m *AnimalUsecaseMock) Create(req *payload.CreateAnimalRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *AnimalUsecaseMock) Find() ([]payload.FindAnimalResponse, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, e.ErrAnimalNotFound
	}
	return args.Get(0).([]payload.FindAnimalResponse), args.Error(1)
}
func (m *AnimalUsecaseMock) FindById(id uint) (*payload.FindAnimalResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, e.ErrAnimalNotFound
	}
	return args.Get(0).(*payload.FindAnimalResponse), args.Error(1)
}
func (m *AnimalUsecaseMock) UpdateOrCreate(id uint, req *payload.UpdateAnimalRequest) (string, error) {
	args := m.Called(id, req)
	return args.Get(0).(string), args.Error(1)
}
func (m *AnimalUsecaseMock) DeleteById(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
