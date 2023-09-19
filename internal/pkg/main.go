package main

import (
	"fmt"
	"log"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/logs"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/benchmark"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/database"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/dataset"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/rest"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Panic(err, "cannot load config HOLA")
	}
	// Set the custom validator for Gin
	err = rest.SetupValidators()
	if err != nil {
		panic(err)
	}

	// Initialize the logger based on the environment
	logs.InitLogger("development")
	defer logs.Logger.Sync() //nolint

	heroService := hero.NewService(database.NewRepository(database.RedisConfig{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.Name,
		Timeout:  config.Redis.Timeout,
	}), dataset.NewRepository(config.Dataset.Path), benchmark.NewRepository(benchmark.Config{
		Timeout:  config.Client.Timeout,
		BasePath: config.Client.BasePath,
		Retries:  config.Client.Retries,
	}))
	handler := rest.NewHandler(heroService)

	server := rest.Routes(handler)
	err = server.Run(fmt.Sprintf("%s%s", ":", config.App.Port))
	if err != nil {
		panic(err)
	}
}
