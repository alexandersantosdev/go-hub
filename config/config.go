package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB_USER string
var DB_PASS string
var DB_NAME string
var DB_HOST string
var DB_PORT string
var SERVER_PORT string

type Config struct {
	DB          *gorm.DB
	SERVER_PORT string
}

func loadEnvs() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	DB_USER = os.Getenv("DB_USER")

	if DB_USER == "" {
		DB_USER = "postgres"
	}

	DB_PASS = os.Getenv("DB_PASS")

	if DB_PASS == "" {
		DB_PASS = "postgres"
	}

	DB_NAME = os.Getenv("DB_NAME")

	if DB_NAME == "" {
		DB_NAME = "test"
	}

	DB_HOST = os.Getenv("DB_HOST")

	if DB_HOST == "" {
		DB_HOST = "localhost"
	}

	DB_PORT = os.Getenv("DB_PORT")

	if DB_PORT == "" {
		DB_PORT = "5432"
	}

	SERVER_PORT = os.Getenv("SERVER_PORT")

	if SERVER_PORT == "" {
		SERVER_PORT = "3000"
	}

	return nil
}

func conectDB() error {

	dsn := fmt.Sprintf("host=%s user%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}

func migrateDB() error {
	return nil
}

func (c *Config) NewConfig() *Config {

	var err error = loadEnvs()
	if err != nil {
		panic(err)
	}

	err = conectDB()
	if err != nil {
		panic(err)
	}

	err = migrateDB()
	if err != nil {
		panic(err)
	}

	return &Config{
		DB:          DB,
		SERVER_PORT: SERVER_PORT,
	}

}
