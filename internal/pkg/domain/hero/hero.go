package hero

import (
	"context"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/errors"
)

func (s Service) GetHero(ctx context.Context, id string) (domain.Hero, error) {
	hero, err := s.storage.GetHero(ctx, id)
	if err != nil {
		return domain.Hero{}, errors.HeroNotFound
	}
	return hero, nil
}

func (s Service) GetAllHeroes(ctx context.Context) ([]domain.Hero, error) {
	heroes, err := s.storage.GetAllHeroes(ctx)
	if err != nil {
		return nil, errors.HeroGetAll
	}
	return heroes, nil
}

func (s Service) GetHeroSuggestion(ctx context.Context, userPreferences domain.UserPreferences) ([]domain.Hero, error) {
	heroes, err := s.storage.GetAllHeroes(ctx)
	if err != nil {
		return nil, errors.HeroGetAll
	}
	// Filter heroes based on user preferences
	heroesToSort := filterHeroes(userPreferences, heroes)
	return heroesToSort.SortHeroesByWinRate(userPreferences.RankName), nil
}

func (s Service) GetDataSet(ctx context.Context) ([][]string, error) {
	records, err := s.dataset.GetRecords(ctx)
	if err != nil {
		return nil, errors.GetDataSet
	}
	return records, nil
}

func (s Service) SaveHeroes(ctx context.Context) (err error) {
	// get records from dataset
	dataset, err := s.dataset.GetRecords(ctx)
	if err != nil {
		return errors.GetDataSet
	}

	// iterate over dataset, build hero and save it in db
	for i, h := range dataset {
		// skip column titles
		if i == 0 {
			continue
		}
		winRates, err := domain.BuildWinRates(h)
		if err != nil {
			return errors.HeroSave
		}
		hero := domain.Hero{
			HeroIndex:        h[0],
			PrimaryAttribute: domain.Attribute(h[2]),
			NameInGame:       h[1],
			Role:             domain.BuildRoles(h[3]),
			WinRates:         winRates,
		}
		err = s.storage.SaveHero(ctx, hero)
		if err != nil {
			return errors.HeroSave
		}
	}
	return nil
}

func (s Service) GetHeroBenchmark(ctx context.Context, id string) (interface{}, error) {
	return s.benchmark.GetHeroBenchmark(ctx, id)
}

// Helper function to filter heroes based on user preferences
func filterHeroes(preferences domain.UserPreferences, heroes []domain.Hero) domain.Heroes {
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
