package database

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/redis/go-redis/v9"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
)

func (r *Repository) GetHero(ctx context.Context, id string) (domain.Hero, error) {
	heroJSON, err := r.rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		return domain.Hero{}, err
	} else if err != nil {
		return domain.Hero{}, fmt.Errorf("failed to get hero from Redis: %w", err)
	}

	// Only attempt to deserialize if there's content.
	var hero domain.Hero
	if heroJSON != "" {
		err = json.Unmarshal([]byte(heroJSON), &hero)
		if err != nil {
			fmt.Println("Error deserializing JSON:", err)
			return domain.Hero{}, fmt.Errorf("failed to get hero from Redis: %w", err)
		}

		fmt.Printf("Retrieved hero: %+v\n", hero)
	} else {
		fmt.Println("No JSON data retrieved")
		return domain.Hero{}, fmt.Errorf("no JSON data retrieved: %w", err)
	}
	return hero, nil
}

func (r *Repository) GetAllHeroes(ctx context.Context) ([]domain.Hero, error) {
	keys, err := r.rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hero keys from Redis: %w", err)
	}

	// Sort the keys
	sort.Strings(keys)

	var heroes []domain.Hero
	for _, key := range keys {
		heroJSON, err := r.rdb.Get(ctx, key).Bytes()
		if err != nil {
			return nil, fmt.Errorf("failed to get hero from Redis: %w", err)
		}

		var hero domain.Hero
		if err := json.Unmarshal(heroJSON, &hero); err != nil {
			return nil, fmt.Errorf("failed to unmarshal hero data: %w", err)
		}

		heroes = append(heroes, hero)
	}
	return heroes, nil
}

func (r *Repository) SaveHero(ctx context.Context, hero domain.Hero) error {
	// Convert hero struct to JSON string
	heroJSON, err := json.Marshal(hero)
	if err != nil {
		fmt.Println("Error marshaling hero:", err)
		return err
	}
	// Save hero data as a Redis string value
	_, err = r.rdb.Set(ctx, hero.HeroIndex, heroJSON, 0).Result()
	if err != nil {
		fmt.Println("Error saving hero:", err)
		return err
	}

	fmt.Println("Hero saved successfully")
	return nil
}
