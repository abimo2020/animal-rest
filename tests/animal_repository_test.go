package tests

import (
	e "animal-rest/errors"
	"animal-rest/mocks"
	"animal-rest/models"
	"animal-rest/repositories"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	animals = []models.Animal{{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Cat",
		Class: "Mammal",
		Legs:  4,
	}, {
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Dog",
		Class: "Mammal",
		Legs:  4,
	}}
)

func TestSaveRepository(t *testing.T) {
	var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
	var animalRepo repositories.AnimalRepository = animalRepoMock
	t.Run("Success", func(t *testing.T) {
		req := animals[0]
		animalRepoMock.On("Save", &req).Return(nil)
		err := animalRepo.Save(&req)
		assert.NotNil(t, req)
		assert.NoError(t, err)
		animalRepoMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		req := models.Animal{}
		animalRepoMock.On("Save", &req).Return(errors.New("failed to save animal"))
		err := animalRepo.Save(&req)
		assert.EqualError(t, err, "failed to save animal")
		animalRepoMock.AssertExpectations(t)
	})
}

func TestFindRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
		var animalRepo repositories.AnimalRepository = animalRepoMock
		animalRepoMock.On("Find").Return(animals, nil)
		res, err := animalRepo.Find()
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, animals, res)
		animalRepoMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
		var animalRepo repositories.AnimalRepository = animalRepoMock
		animalRepoMock.On("Find").Return(nil)
		res, err := animalRepo.Find()
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, err, e.ErrAnimalNotFound.Error())
		animalRepoMock.AssertExpectations(t)
	})
}

func TestFindByIdRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
		var animalRepo repositories.AnimalRepository = animalRepoMock
		animalRepoMock.On("FindById", uint(1)).Return(&animals[0], nil)
		res, err := animalRepo.FindById(1)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, animals[0].ID, res.ID)
		assert.Equal(t, &animals[0], res)
		animalRepoMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
		var animalRepo repositories.AnimalRepository = animalRepoMock
		animalRepoMock.On("FindById", uint(3)).Return(nil)
		res, err := animalRepo.FindById(3)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, err, e.ErrAnimalNotFound.Error())
		animalRepoMock.AssertExpectations(t)
	})
}

func TestFindByNameRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
		var animalRepo repositories.AnimalRepository = animalRepoMock
		animalRepoMock.On("FindByName", animals[0].Name).Return(&animals[0], nil)
		res, err := animalRepo.FindByName(animals[0].Name)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, animals[0].Name, res.Name)
		assert.Equal(t, &animals[0], res)
		animalRepoMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
		var animalRepo repositories.AnimalRepository = animalRepoMock
		animalRepoMock.On("FindByName", "Monkey").Return(nil)
		res, err := animalRepo.FindByName("Monkey")
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualError(t, err, e.ErrAnimalNotFound.Error())
		animalRepoMock.AssertExpectations(t)
	})
}

func TestDeleteByIdRepository(t *testing.T) {
	var animalRepoMock = &mocks.AnimalRepositoryMock{mock.Mock{}}
	var animalRepo repositories.AnimalRepository = animalRepoMock
	t.Run("Success", func(t *testing.T) {
		animalRepoMock.On("DeleteById", uint(1)).Return(nil)
		err := animalRepo.DeleteById(1)
		assert.NoError(t, err)
		animalRepoMock.AssertExpectations(t)
	})
	t.Run("Failure", func(t *testing.T) {
		animalRepoMock.On("DeleteById", uint(2)).Return(errors.New("failed to delete animal"))
		err := animalRepo.DeleteById(2)
		assert.EqualError(t, err, "failed to delete animal")
		animalRepoMock.AssertExpectations(t)
	})
}
