package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(handler *Handler) *gin.Engine {
	r := gin.Default()
	// API endpoint to fetch all heroes
	r.GET("/api/heroes", handler.GetAllHeroes)
	// API endpoint to suggest a random hero based on user preferences
	r.POST("/api/hero-picker", handler.GetHeroSuggestion)
	r.GET("/api/dataset", handler.GetDataSet)
	r.POST("/api/save-heroes", handler.SaveHeroes)
	r.GET("/api/hero/:id", handler.GetHero)
	r.GET("/api/hero/:id/benchmark", handler.GetHeroBenchmark)
	return r
}
