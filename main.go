package main

import (
	"animal-rest/config"
	"animal-rest/routes"
	"animal-rest/util"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("error loading .env file")
	}
	config.InitDB()
	config.InitMigration()
	config.RedisInit()
}

func main() {
	e := routes.New()
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}
	e.Logger.Fatal(e.Start(":8080"))
}
