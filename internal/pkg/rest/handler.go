//go:generate moq -pkg mocks -out ./mocks/handler_mocks.go -skip-ensure . HeroService
package rest

import (
	"context"

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
	GetHeroBenchmark(ctx context.Context, id string) (interface{}, error)
}

func NewHandler(heroService HeroService) *Handler {
	return &Handler{
		heroService: heroService,
	}
}
