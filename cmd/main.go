package main

import (
	"github.com/joho/godotenv"
	"log"
	"medods/internal/handler"
	"medods/internal/models"
	"medods/internal/repository"
	"medods/internal/service"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Println(err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(models.Server)
	if err := srv.Run(os.Getenv("HOST_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
