package rest

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/errors"
)

// GetHero Handler function to fetch a hero by index
func (h *Handler) GetHero(c *gin.Context) {
	id := c.Param("id")
	hero, err := h.heroService.GetHero(c, id)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, hero)
}

// GetAllHeroes Handler function to fetch all heroes
func (h *Handler) GetAllHeroes(c *gin.Context) {
	heroes, err := h.heroService.GetAllHeroes(c)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, heroes)
}

// GetDataSet Handler function to obtain the original data from the dataset
func (h *Handler) GetDataSet(c *gin.Context) {
	dataset, err := h.heroService.GetDataSet(c)
	if err != nil {
		handleError(c, err)
		return
	}
	csvBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(csvBuffer)
	err = writer.WriteAll(dataset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "write all dataset"})
		return
	}
	c.JSON(http.StatusOK, dataset)
}

// GetHeroSuggestion Handler function to suggest a random hero based on user preferences
func (h *Handler) GetHeroSuggestion(c *gin.Context) {
	var userPreferences domain.UserPreferences
	// Bind the JSON data from the request body to userPreferences struct
	if err := c.ShouldBindJSON(&userPreferences); err != nil {
		// Handle validation errors
		errorMsgs := handleShouldBindJSONErrors(err.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMsgs,
		})
		return
	}

	heroSuggestion, err := h.heroService.GetHeroSuggestion(c, userPreferences)
	if err != nil {
		handleError(c, err)
		return
	}
	if heroSuggestion == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No hero found matching the criteria"})
		return
	}
	c.JSON(http.StatusOK, heroSuggestion)
}

// SaveHeroes Handler function to save all dataset record into the database
func (h *Handler) SaveHeroes(c *gin.Context) {
	err := h.heroService.SaveHeroes(c)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetHeroBenchmark(c *gin.Context) {
	id := c.Param("id")
	hero, err := h.heroService.GetHeroBenchmark(c, id)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, hero)
}

func (h *Handler) GetFullHeroInfo(c *gin.Context) {
	id := c.Param("id")
	hero, err := h.heroService.GetFullHeroInfo(c, id)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, hero)
}

func handleError(c *gin.Context, err error) {
	if customErr, ok := err.(errors.CustomError); ok {
		c.JSON(customErr.HTTPCode(), gin.H{"code": customErr.InternalCode(), "message": customErr.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
