package tests

import (
	e "animal-rest/errors"
	"animal-rest/mocks"
	"animal-rest/payload"
	"animal-rest/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createReq = payload.CreateAnimalRequest{
		Name:  "Cat",
		Class: "Mammal",
		Legs:  4,
	}
	updateReq = payload.UpdateAnimalRequest{
		Name:  "Dog",
		Class: "Mammal",
		Legs:  4,
	}
	resp = []payload.FindAnimalResponse{
		{
			ID:    1,
			Name:  "Cat",
			Class: "Mammal",
			Legs:  4,
		}, {
			ID:    2,
			Name:  "Dog",
			Class: "Mammal",
			Legs:  4,
		},
	}
)

func TestCreateUsecase(t *testing.T) {
	var animalUsecaseMock = &mocks.AnimalUsecaseMock{mock.Mock{}}
	var animalUsecase usecase.AnimalUsecase = animalUsecaseMock
	t.Run("Success", func(t *testing.T) {
		animalUsecaseMock.On("Create", &createReq).Return(nil)
		err := animalUsecase.Create(&createReq)
		assert.NotNil(t, createReq)
		assert.NoError(t, err)
		animalUsecaseMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		req := payload.CreateAnimalRequest{}
		animalUsecaseMock.On("Create", &req).Return(errors.New("failed to save animal"))
		err := animalUsecase.Create(&req)
		assert.EqualError(t, err, "failed to save animal")
		animalUsecaseMock.AssertExpectations(t)
	})
}

func TestUpdateOrCreateUsecase(t *testing.T) {
	var animalUsecaseMock = &mocks.AnimalUsecaseMock{mock.Mock{}}
	var animalUsecase usecase.AnimalUsecase = animalUsecaseMock
	t.Run("Success", func(t *testing.T) {
		animalUsecaseMock.On("UpdateOrCreate", uint(2), &updateReq).Return("create", nil)
		_, err := animalUsecase.UpdateOrCreate(2, &updateReq)
		assert.NoError(t, err)
		animalUsecaseMock.AssertExpectations(t)
	})
}

func TestFindUsecase(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var animalUsecaseMock = &mocks.AnimalUsecaseMock{mock.Mock{}}
		var animalUsecase usecase.AnimalUsecase = animalUsecaseMock
		animalUsecaseMock.On("Find").Return(resp, nil)
		res, err := animalUsecase.Find()
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, resp, res)
		animalUsecaseMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		var animalUsecaseMock = &mocks.AnimalUsecaseMock{mock.Mock{}}
		var animalUsecase usecase.AnimalUsecase = animalUsecaseMock
		animalUsecaseMock.On("Find").Return(nil)
		res, err := animalUsecase.Find()
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, err, e.ErrAnimalNotFound.Error())
		animalUsecaseMock.AssertExpectations(t)
	})
}

func TestFindByIdUsecase(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var animalUsecaseMock = &mocks.AnimalUsecaseMock{mock.Mock{}}
		var animalUsecase usecase.AnimalUsecase = animalUsecaseMock
		animalUsecaseMock.On("FindById", uint(1)).Return(&resp[0], nil)
		res, err := animalUsecase.FindById(1)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, resp[0].ID, res.ID)
		assert.Equal(t, &resp[0], res)
		animalUsecaseMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		var animalUsecaseMock = &mocks.AnimalUsecaseMock{mock.Mock{}}
		var animalUsecase usecase.AnimalUsecase = animalUsecaseMock
		animalUsecaseMock.On("FindById", uint(3)).Return(nil)
		res, err := animalUsecase.FindById(3)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, err, e.ErrAnimalNotFound.Error())
		animalUsecaseMock.AssertExpectations(t)
	})
}

func TestDeleteByIdUsecase(t *testing.T) {
	var animalUsecaseMock = &mocks.AnimalUsecaseMock{mock.Mock{}}
	var animalUsecase usecase.AnimalUsecase = animalUsecaseMock
	t.Run("Success", func(t *testing.T) {
		animalUsecaseMock.On("DeleteById", uint(1)).Return(nil)
		err := animalUsecase.DeleteById(1)
		assert.NoError(t, err)
		animalUsecaseMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		animalUsecaseMock.On("DeleteById", uint(2)).Return(errors.New("failed to delete animal"))
		err := animalUsecase.DeleteById(2)
		assert.EqualError(t, err, "failed to delete animal")
		animalUsecaseMock.AssertExpectations(t)
	})
}
