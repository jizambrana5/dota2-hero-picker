//go:generate moq -pkg mocks -out ./mocks/handler_mocks.go -skip-ensure . HeroService
package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
)

type Handler struct {
	heroService HeroService
}

type HeroService interface {
	GetHero(ctx context.Context, id string) (domain.Hero, error)
	GetAllHeroes(ctx context.Context) ([]domain.Hero, error)
	GetDataSet(ctx context.Context) ([][]string, error)
	GetHeroSuggestion(ctx context.Context, preferences domain.UserPreferences) ([]domain.Hero, error)
	SaveHeroes(ctx context.Context) error
	GetHeroBenchmark(ctx context.Context, heroID string) (interface{}, error)
	GetFullHeroInfo(ctx context.Context, heroID string) (domain.FullHeroInfo, error)
}

func NewHandler(heroService HeroService) *Handler {
	return &Handler{
		heroService: heroService,
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
