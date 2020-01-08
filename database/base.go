package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type dbConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *dbConfig {
	dbConfig := dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return &dbConfig
}

func pgConnectUrl(dbConfig *dbConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
	)
}

func GetPgConnnection() (*gorm.DB, error) {
	return gorm.Open("postgres", pgConnectUrl(buildDBConfig()))
}
