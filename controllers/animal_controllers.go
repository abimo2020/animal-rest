package controllers

import (
	"animal-rest/errors"
	"animal-rest/payload"
	"animal-rest/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AnimalController interface {
	Create(c echo.Context) error
	UpdateOrCreate(c echo.Context) error
	Find(c echo.Context) error
	FindById(c echo.Context) error
	DeleteById(c echo.Context) error
}

type animalController struct {
	animalUsecase usecase.AnimalUsecase
}

func NewAnimalController(au usecase.AnimalUsecase) AnimalController {
	return &animalController{animalUsecase: au}
}

func (ac *animalController) Create(c echo.Context) error {
	var req payload.CreateAnimalRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := ac.animalUsecase.Create(&req); err != nil {
		if err == errors.ErrAnimalAlreadyExists {
			return echo.NewHTTPError(http.StatusConflict, map[string]any{
				"message": err.Error(),
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	return echo.NewHTTPError(http.StatusCreated, map[string]any{
		"message": "success to create animal",
	})
}

func (ac *animalController) UpdateOrCreate(c echo.Context) error {
	var req payload.UpdateAnimalRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	status, err := ac.animalUsecase.UpdateOrCreate(uint(id), &req)
	if err != nil {
		if err == errors.ErrAnimalAlreadyExists {
			return echo.NewHTTPError(http.StatusConflict, map[string]any{
				"message": err.Error(),
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	if status == "create" {
		return echo.NewHTTPError(http.StatusCreated, map[string]any{
			"message": "success to create animal",
		})
	}
	return echo.NewHTTPError(http.StatusOK, map[string]any{
		"message": "success to update animal",
	})
}

func (ac animalController) Find(c echo.Context) error {
	resp, err := ac.animalUsecase.Find()
	if err != nil {
		if err == errors.ErrAnimalNotFound {
			return echo.NewHTTPError(http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	return echo.NewHTTPError(http.StatusOK, map[string]any{
		"message": "success to get animals",
		"data":    resp,
	})
}

func (ac animalController) FindById(c echo.Context) error {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	resp, err := ac.animalUsecase.FindById(uint(id))
	if err != nil {
		if err == errors.ErrAnimalNotFound {
			return echo.NewHTTPError(http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	return echo.NewHTTPError(http.StatusOK, map[string]any{
		"message": "success to get animal",
		"data":    resp,
	})
}

func (ac animalController) DeleteById(c echo.Context) error {
	stringId := c.Param("id")
	id, _ := strconv.Atoi(stringId)
	err := ac.animalUsecase.DeleteById(uint(id))
	if err != nil {
		if err == errors.ErrAnimalNotFound {
			return echo.NewHTTPError(http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
	}
	return echo.NewHTTPError(http.StatusOK, map[string]any{
		"message": "success to delete animal",
	})
}
