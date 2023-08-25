package rest

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
)

type Dataset struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	// Add more fields as needed based on your dataset's structure
}

// GetAllHeroes Handler function to fetch all heroes
func (h *Handler) GetAllHeroes(c *gin.Context) {
	heroes, err := h.heroService.GetAllHeroes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get all heroes error"})
		return
	}
	c.JSON(http.StatusOK, heroes)
}

// GetHeroSuggestion Handler function to suggest a random hero based on user preferences
func (h *Handler) GetHeroSuggestion(c *gin.Context) {
	var userPreferences domain.UserPreferences
	// Bind the JSON data from the request body to userPreferences struct
	if err := c.ShouldBindJSON(&userPreferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	heroSuggestion, err := h.heroService.GetHeroSuggestion(c, userPreferences)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get hero suggestion error"})
		return
	}
	if heroSuggestion == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No hero found matching the criteria"})
	}
	c.JSON(http.StatusOK, heroSuggestion)
}

// GetDataSet Handler function to obtain the original data from the dataset
func (h *Handler) GetDataSet(c *gin.Context) {
	dataset, err := h.heroService.GetDataSet(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get data set error"})
		return
	}
	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	writer.WriteAll(dataset)
	c.JSON(http.StatusOK, dataset)
}

// SaveHeroes Handler function to save all dataset record into the database
func (h *Handler) SaveHeroes(c *gin.Context) {
	err := h.heroService.SaveHeroes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving heroes from dataset"})
		return
	}
	c.JSON(http.StatusOK, nil)
}

// GetHero Handler function to fetch a hero by index
func (h *Handler) GetHero(c *gin.Context) {
	id := c.Param("id")
	hero, err := h.heroService.GetHero(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting heroes from dataset"})
		return
	}
	c.JSON(http.StatusOK, hero)
}
