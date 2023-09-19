package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes(handler *Handler) *gin.Engine {
	//Set up the logger based on the environment
	logger, err := setupLogger()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// Use the custom logging middleware for all routes
	r.Use(loggingMiddleware(logger))
	// Use the MetricsMiddleware
	r.Use(MetricsMiddleware())

	// Ping
	r.GET("/ping", handler.Ping)

	// API endpoint to fetch all heroes
	r.GET("/api/heroes", handler.GetAllHeroes)
	// API endpoint to suggest a random hero based on user preferences
	r.POST("/api/hero-picker", handler.GetHeroSuggestion)
	r.GET("/api/dataset", handler.GetDataSet)
	r.POST("/api/save-heroes", handler.SaveHeroes)
	r.GET("/api/hero/:id", handler.GetHero)
	r.GET("/api/hero/:id/benchmark", handler.GetHeroBenchmark)

	//Monitoring
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}
