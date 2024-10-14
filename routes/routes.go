package routes

import (
	"animal-rest/config"
	"animal-rest/controllers"
	"animal-rest/middlewares"
	"animal-rest/repositories"
	"animal-rest/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LoggerMiddleware(e)
	e.Use(middleware.RemoveTrailingSlash())

	animalRepository := repositories.NewAnimalRepository(config.DB)
	redisRepository := repositories.NewRedisRepository(config.Redis)
	animalUsecase := usecase.NewAnimalUsecase(animalRepository, redisRepository)
	animalController := controllers.NewAnimalController(animalUsecase)

	e.POST("v1/animal", animalController.Create)
	e.PUT("v1/animal/:id", animalController.UpdateOrCreate)
	e.GET("v1/animal", animalController.Find)
	e.GET("v1/animal/:id", animalController.FindById)
	e.DELETE("v1/animal/:id", animalController.DeleteById)

	return e
}
