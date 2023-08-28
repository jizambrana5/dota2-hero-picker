package main

import (
	"log"
	"time"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/database"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/dataset"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/rest"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	heroService := hero.NewService(database.NewRepository(database.RedisConfig{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Name,
		Timeout:  time.Duration(config.Timeout),
	}), dataset.NewRepository("./internal/pkg/repository/dataset/dataset.csv"))
	handler := rest.NewHandler(heroService)

	server := rest.Routes(handler)
	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
