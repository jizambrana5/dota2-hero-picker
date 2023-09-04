package database

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockRedisClient struct {
	GetFn  func(context.Context, string) *redis.StringCmd
	KeysFn func(context.Context, string) *redis.StringSliceCmd
	PingFn func(ctx context.Context) *redis.StatusCmd
	SetFn  func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return m.GetFn(ctx, key)
}

func (m *MockRedisClient) Keys(ctx context.Context, key string) *redis.StringSliceCmd {
	return m.KeysFn(ctx, key)
}

func (m *MockRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	return m.PingFn(ctx)
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return m.SetFn(ctx, key, value, expiration)
}

type databaseTestSuite struct {
	suite.Suite
	ctx    context.Context
	repo   *Repository
	dbMock *MockRedisClient
}

func (t *databaseTestSuite) SetupTest() {
	t.ctx = context.Background()
	t.dbMock = &MockRedisClient{
		GetFn: func(ctx context.Context, s string) *redis.StringCmd {
			rd := redis.NewStringCmd(t.ctx, "hero_test")
			rd.SetVal(`{"hero_index": "1"}`)
			return rd
		},
		KeysFn: func(ctx context.Context, s string) *redis.StringSliceCmd {
			heroesKeys := []string{`hero_1`, `hero_2`}
			rd := redis.NewStringSliceCmd(ctx, "all_herores")
			rd.SetVal(heroesKeys)
			return rd
		},
	}
	t.repo = NewCustomRepository(t.dbMock)
}

func (t *databaseTestSuite) Test_NewRepository_Panic() {
	defer func() { recover() }() // no lint
	NewRepository(RedisConfig{})
	t.Errorf(nil, "should have panicked")
}
func (t *databaseTestSuite) Test_NewRepository_Success() {
	t.dbMock = new(MockRedisClient)
	t.dbMock.PingFn = func(ctx context.Context) *redis.StatusCmd {
		return redis.NewStatusCmd(t.ctx)
	}
	defer func() {
		r := recover()
		assert.NotNil(t.T(), r)
		assert.Contains(t.T(), r.(string), "failed to connect to Redis")
	}()
	repo := NewRepository(RedisConfig{
		Addr:     "localhost:1111",
		Password: "",
		DB:       0,
		Timeout:  time.Second,
	})
	t.NotNil(repo)
}

func TestDatabase(t *testing.T) {
	suite.Run(t, new(databaseTestSuite))
}
