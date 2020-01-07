package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/khuongnguyenBlue/vine/controllers"
	"github.com/khuongnguyenBlue/vine/database"
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
	db, err := database.GetPgConnnection()
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	database.Migrate(db)
	database.SeedData(db)

	controller := controllers.NewController(db)
	r := routes.Setup(controller)
	r.Run()
}
