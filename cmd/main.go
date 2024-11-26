package cmd

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"medods/internal/handler"
	"medods/internal/models"
	"medods/internal/repository"
	"medods/internal/service"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Println(err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Println(err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(models.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
