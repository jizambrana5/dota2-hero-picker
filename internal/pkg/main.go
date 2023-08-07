package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/database"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/dataset"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/rest"
)

func main() {
	r := gin.Default()
	heroService := hero.NewService(database.NewRepository(database.RedisConfig{}), dataset.NewRepository("./internal/pkg/repository/dataset2/dataset.csv"))
	handler := rest.NewHandler(heroService)
	// API endpoint to fetch all heroes
	r.GET("/api/heroes", handler.GetAllHeroes)

	// API endpoint to suggest a random hero based on user preferences
	r.POST("/api/hero-picker", handler.GetHeroSuggestion)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
