package usecase

import (
	"animal-rest/errors"
	"animal-rest/models"
	"animal-rest/payload"
	"animal-rest/repositories"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AnimalUsecase interface {
	Create(req *payload.CreateAnimalRequest) error
	UpdateOrCreate(id uint, req *payload.UpdateAnimalRequest) (status string, err error)
	Find() ([]payload.FindAnimalResponse, error)
	FindById(id uint) (*payload.FindAnimalResponse, error)
	DeleteById(id uint) error
}

type animalUsecase struct {
	animalRepository repositories.AnimalRepository
	redisRepository  repositories.RedisRepository
}

func NewAnimalUsecase(ar repositories.AnimalRepository, rr repositories.RedisRepository) AnimalUsecase {
	return &animalUsecase{
		animalRepository: ar,
		redisRepository:  rr,
	}
}

func (au *animalUsecase) Create(req *payload.CreateAnimalRequest) error {
	data := models.Animal{
		Name:  req.Name,
		Class: req.Class,
		Legs:  req.Legs,
	}
	if err := au.animalRepository.Save(&data); err != nil {
		return err
	}
	if err := au.redisRepository.Delete("animal"); err != nil {
		fmt.Printf("error deleting redis key: %v\n", err)
	}
	return nil
}

func (au *animalUsecase) UpdateOrCreate(id uint, req *payload.UpdateAnimalRequest) (status string, err error) {
	data, err := au.animalRepository.FindById(id)
	if err != nil {
		if err == errors.ErrAnimalNotFound {
			if _, err = au.animalRepository.FindByName(req.Name); err == nil {
				return "", errors.ErrAnimalAlreadyExists
			}
			if err = au.animalRepository.Save(&models.Animal{
				Model: gorm.Model{
					ID:        id,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Name:  req.Name,
				Class: req.Class,
				Legs:  req.Legs,
			}); err != nil {
				return
			}
			status = "create"
			if err := au.redisRepository.Delete("animal"); err != nil {
				fmt.Printf("error deleting redis key: %v\n", err)
			}
			return
		} else {
			return
		}
	}
	if req.Name != "" {
		data.Name = req.Name
	}
	if req.Class != "" {
		data.Class = req.Class
	}
	data.Legs = req.Legs
	if err = au.animalRepository.Save(data); err != nil {
		return
	}
	status = "update"
	if err := au.redisRepository.Delete("animal"); err != nil {
		fmt.Printf("error deleting redis key: %v\n", err)
	}
	return
}

func (au *animalUsecase) Find() ([]payload.FindAnimalResponse, error) {
	var resp []payload.FindAnimalResponse

	data, err := au.animalRepository.Find()
	if err != nil {
		return nil, err
	}

	redisRes, err := au.redisRepository.Get("animal")
	if err == redis.Nil {
		fmt.Println("redis key not found")
		for _, val := range data {
			resp = append(resp, payload.FindAnimalResponse{
				ID:    val.ID,
				Name:  val.Name,
				Class: val.Class,
				Legs:  val.Legs,
			})
		}
		value, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("error marshalling data: %v\n", err)
		}
		if err = au.redisRepository.Set("animal", string(value)); err != nil {
			fmt.Printf("error setting data in redis: %v\n", err)
		}
		return resp, nil
	} else if err != nil {
		fmt.Printf("error retrieving from redis: %v\n", err)
	} else {
		if err = json.Unmarshal([]byte(redisRes), &resp); err != nil {
			fmt.Printf("error unmarshalling cached data: %v\n", err)
		}
	}
	for _, val := range data {
		resp = append(resp, payload.FindAnimalResponse{
			ID:    val.ID,
			Name:  val.Name,
			Class: val.Class,
			Legs:  val.Legs,
		})
	}

	return resp, nil
}

func (au *animalUsecase) FindById(id uint) (*payload.FindAnimalResponse, error) {
	var resp payload.FindAnimalResponse
	data, err := au.animalRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	redisRes, err := au.redisRepository.Get("animal:" + string(id))
	if err == redis.Nil {
		fmt.Println("key not found")
		resp = payload.FindAnimalResponse{
			ID:    data.ID,
			Name:  data.Name,
			Class: data.Class,
			Legs:  data.Legs,
		}
		value, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("error marshalling data: %v\n", err)
		}
		stringId := strconv.Itoa(int(id))
		if err = au.redisRepository.Set("animal:"+stringId, string(value)); err != nil {
			fmt.Printf("error setting data in redis: %v\n", err)
		}
	} else if err != nil {
		fmt.Printf("error retrieving from redis: %v\n", err)
	} else {
		if err = json.Unmarshal([]byte(redisRes), &resp); err != nil {
			fmt.Printf("error unmarshalling cached data: %v\n", err)
		}
	}

	return &resp, nil
}

func (au *animalUsecase) DeleteById(id uint) error {
	if err := au.animalRepository.DeleteById(id); err != nil {
		return err
	}
	if err := au.redisRepository.Delete("animal"); err != nil {
		fmt.Printf("error deleting redis key: %v\n", err)
	}
	stringId := strconv.Itoa(int(id))
	if err := au.redisRepository.Delete("animal:" + stringId); err != nil {
		fmt.Printf("error deleting redis key: %v\n", err)
	}
	return nil
}
