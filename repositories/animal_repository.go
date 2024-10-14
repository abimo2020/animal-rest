package repositories

import (
	"animal-rest/errors"
	"animal-rest/models"
	"strings"

	"gorm.io/gorm"
)

type AnimalRepository interface {
	Save(req *models.Animal) error
	Find() ([]models.Animal, error)
	FindById(id uint) (*models.Animal, error)
	FindByName(name string) (*models.Animal, error)
	DeleteById(id uint) error
}

type animalRepository struct {
	db *gorm.DB
}

func NewAnimalRepository(db *gorm.DB) AnimalRepository {
	return &animalRepository{db}
}

func (ar *animalRepository) Save(req *models.Animal) error {
	if err := ar.db.Save(&req).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.ErrAnimalAlreadyExists
		}
		return err
	}
	return nil
}

func (ar *animalRepository) Find() ([]models.Animal, error) {
	var resp []models.Animal
	if err := ar.db.Find(&resp).Error; err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, errors.ErrAnimalNotFound
	}
	return resp, nil
}

func (ar *animalRepository) FindById(id uint) (*models.Animal, error) {
	var resp models.Animal
	if err := ar.db.First(&resp, id).Error; err != nil {
		if id <= 0 {
			return nil, errors.ErrInvalidAnimalId
		} else if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrAnimalNotFound
		}
		return nil, err
	}
	return &resp, nil
}

func (ar *animalRepository) FindByName(name string) (*models.Animal, error) {
	var resp models.Animal
	if err := ar.db.Where("name = ?", name).First(&resp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrAnimalNotFound
		}
		return nil, err
	}
	return &resp, nil
}

func (ar *animalRepository) DeleteById(id uint) error {
	resp := ar.db.Delete(&models.Animal{}, id)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return errors.ErrAnimalNotFound
	}
	return nil
}
