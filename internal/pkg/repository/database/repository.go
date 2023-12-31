package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
)

// RedisConfig holds the configuration needed to connect to Redis.
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	Timeout  time.Duration
}

type Repository struct {
	rdb Database
}

type Database interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
	Ping(ctx context.Context) *redis.StatusCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

var _ hero.Storage = (*Repository)(nil)

func NewRepository(config RedisConfig) *Repository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to Redis: %v", err))
	}

	return &Repository{
		rdb: rdb,
	}
}

func NewCustomRepository(db Database) *Repository {
	return &Repository{
		rdb: db,
	}
}
