package rest

import (
	"context"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
)

type Handler struct {
	heroService HeroService
}

type HeroService interface {
	GetAllHeroes(ctx context.Context) ([]hero.Hero, error)
	GetHeroSuggestion(ctx context.Context) ([]hero.Hero, error)
	GetDataSet(ctx context.Context) ([][]string, error)
}

//var _ HeroService = (*investment.Service)(nil)

func NewHandler(heroService HeroService) *Handler {
	return &Handler{
		heroService: heroService,
	}
}
