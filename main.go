package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/khuongnguyenBlue/vine/configs"
	"github.com/khuongnguyenBlue/vine/migrations"
	"github.com/khuongnguyenBlue/vine/routes"
	"github.com/khuongnguyenBlue/vine/seeds"
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
	configs.DB.LogMode(true)

	if dbErr != nil {
		panic(dbErr)
	}

	defer configs.DB.Close()

	migrations.Migrate()

	c := make(chan error)
	go seeds.All(c)
	log.Println("wait")
	if seedErr := <-c; seedErr != nil {
		log.Println("Failed to seeding")
	} else {
		log.Println("Finished seeding")
	}

	r := routes.Setup()
	r.Run()
}
