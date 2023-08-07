package rest

import (
	"math/rand"
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
	/* TODO: Implement service
	heroes, err := h.heroService.GetAllHeroes(c)
	if err != nil {
		return err
	}
	*/
	heroes := []domain.Hero{
		{HeroIndex: 1, PrimaryAttribute: "str", NameInGame: "Hero 1", Role: []domain.Role{"Carry", "Disabler"}},
		{HeroIndex: 2, PrimaryAttribute: "agi", NameInGame: "Hero 2", Role: []domain.Role{"Carry", "Escape"}},
		{HeroIndex: 3, PrimaryAttribute: "int", NameInGame: "Hero 3", Role: []domain.Role{"Support", "Nuker"}},
		// Add more heroes as needed
	}
	c.JSON(http.StatusOK, heroes)
}

// GetHeroSuggestion Handler function to suggest a random hero based on user preferences
func (h *Handler) GetHeroSuggestion(c *gin.Context) {
	/* TODO: Implement service
	hero, err := h.heroService.GetHeroSuggestion(c)
	if err != nil {
		return err
	}
	*/

	heroes := []domain.Hero{
		{HeroIndex: 1, PrimaryAttribute: "str", NameInGame: "Hero 1", Role: []domain.Role{"Carry", "Disabler"}},
		{HeroIndex: 2, PrimaryAttribute: "agi", NameInGame: "Hero 2", Role: []domain.Role{"Carry", "Escape"}},
		{HeroIndex: 3, PrimaryAttribute: "int", NameInGame: "Hero 3", Role: []domain.Role{"Support", "Nuker"}},
		// Add more heroes as needed
	}

	var userPreferences domain.UserPreferences
	// Bind the JSON data from the request body to userPreferences struct
	if err := c.ShouldBindJSON(&userPreferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Filter heroes based on user preferences
	filteredHeroes := filterHeroes(userPreferences, heroes)

	// Return a random hero from the filtered list
	if len(filteredHeroes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No hero found matching the criteria"})
	} else {
		randomHero := filteredHeroes[rand.Intn(len(filteredHeroes))]
		c.JSON(http.StatusOK, randomHero)
	}
}

// Helper function to filter heroes based on user preferences
func filterHeroes(preferences domain.UserPreferences, heroes []domain.Hero) []domain.Hero {
	var filtered []domain.Hero

	for _, hero := range heroes {
		if hero.PrimaryAttribute == preferences.PrimaryAttribute {
			if hasAllRoles(hero.Role, preferences.Roles) {
				filtered = append(filtered, hero)
			}
		}
	}

	return filtered
}

// Helper function to check if a hero has all specified roles
func hasAllRoles(heroRoles []domain.Role, userRoles []domain.Role) bool {
	for _, role := range userRoles {
		if !contains(heroRoles, role) {
			return false
		}
	}
	return true
}

// Helper function to check if a string is present in a slice of strings
func contains(slice []domain.Role, str domain.Role) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
