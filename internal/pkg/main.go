package main

import (
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/database"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/dataset"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/rest"
)

func main() {
	heroService := hero.NewService(database.NewRepository(database.RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Timeout:  100000000000000,
	}), dataset.NewRepository("./internal/pkg/repository/dataset/dataset.csv"))
	handler := rest.NewHandler(heroService)

	server := rest.Routes(handler)
	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
