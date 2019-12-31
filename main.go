package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/migrations"
	"github.com/khuongnguyenBlue/vine/routes"
	"log"
)

func init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var dbErr error;
	configs.DB, dbErr = gorm.Open("postgres", configs.DBUrl(configs.BuildDBConfig()))

	if dbErr != nil {
		fmt.Println("Cannot connect to DB: ", configs.BuildDBConfig())
		panic(dbErr)
	}

	defer configs.DB.Close()

	migrations.Migrate()

	r := routes.Setup()
	r.Run()
}
