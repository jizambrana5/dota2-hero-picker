package hero

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/errors"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/logs"
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
	// 1. Get records from dataset.
	dataset, err := s.dataset.GetRecords(ctx)
	if err != nil {
		return errors.GetDataSet
	}

	// 2. Create chanel
	heroCh := make(chan []string, len(dataset)-1)
	heroErrCh := make(chan error, len(dataset)-1)
	workers := 3
	// 3. Create goroutines to process and save heroes into database concurrently.
	for i := 0; i <= workers; i++ {
		go s.saveSingleHero(ctx, heroCh, heroErrCh)
	}

	// 4. Iterate over dataset and send it to a channel to be processed.
	for i, h := range dataset {
		// skip column titles
		if i == 0 {
			continue
		}
		heroCh <- h
	}
	// 5. Close hero channel
	close(heroCh)

	// 6. Iterate over errors channel
	heroErrors := make([]error, 0)

	for j := 0; j < len(dataset)-1; j++ {
		e := <-heroErrCh
		if e != nil {
			heroErrors = append(heroErrors, e)
		}
	}
	/*for e := range heroErrCh {
		if e != nil {
			heroErrors = append(heroErrors, e)
		}
	}
	close(heroErrCh) */

	if len(heroErrors) > 0 {
		logs.Logger.Error("save_all_heroes",
			zap.Error(fmt.Errorf("failed saved heroes publication: %v", heroErrors)))
		return errors.SaveAllHeroes
	}

	return nil
}

func (s Service) saveSingleHero(ctx context.Context, heroCh chan []string, heroErrCh chan error) {
	for h := range heroCh {
		winRates, err := domain.BuildWinRates(h)
		if err != nil {
			heroErrCh <- errors.HeroSave
			continue
		}
		hero := domain.Hero{
			HeroIndex:        h[0],
			PrimaryAttribute: domain.Attribute(h[2]),
			NameInGame:       h[1],
			Role:             domain.BuildRoles(h[3]),
			WinRates:         winRates,
		}
		heroErrCh <- s.storage.SaveHero(ctx, hero)
	}
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
