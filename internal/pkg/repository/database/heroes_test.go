package database

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

func (t *databaseTestSuite) TestRepository_GetHero_DatabaseError() {
	t.dbMock.GetFn = func(ctx context.Context, s string) *redis.StringCmd {
		red := redis.NewStringCmd(ctx, nil, redis.Nil)
		red.SetErr(errors.New(""))
		return red
	}
	hero, err := t.repo.GetHero(context.Background(), "heroID")

	t.NotNil(err)
	t.Empty(hero)
}

func (t *databaseTestSuite) TestRepository_GetHero_WithoutData() {
	t.dbMock.GetFn = func(ctx context.Context, s string) *redis.StringCmd {
		red := redis.NewStringCmd(ctx, nil, redis.Nil)
		return red
	}
	hero, err := t.repo.GetHero(context.Background(), "heroID")

	t.NotNil(err)
	t.Empty(hero)
}

func (t *databaseTestSuite) TestRepository_GetHero_UnmarshallError() {
	t.dbMock.GetFn = func(ctx context.Context, s string) *redis.StringCmd {
		rd := redis.NewStringCmd(t.ctx, "hero_test")
		rd.SetVal(`{"hero_index": "1""}`)
		return rd
	}
	hero, err := t.repo.GetHero(context.Background(), "heroID")

	t.NotNil(err)
	t.Empty(hero)
}

func (t *databaseTestSuite) TestRepository_GetHero_Success() {
	hero, err := t.repo.GetHero(context.Background(), "heroID")
	t.Nil(err)
	t.NotEmpty(hero)
	t.Equal(1, hero.HeroIndex)
}

func (t *databaseTestSuite) TestRepository_GetAllHeroes_GetKeysError() {
	t.dbMock.KeysFn = func(ctx context.Context, s string) *redis.StringSliceCmd {
		keys := redis.NewStringSliceCmd(ctx, "")
		keys.SetErr(errors.New(""))
		return keys
	}
	heroes, err := t.repo.GetAllHeroes(context.Background())
	t.NotNil(err)
	t.Empty(heroes)
}

func (t *databaseTestSuite) TestRepository_GetAllHeroes_GetHeroError() {
	t.dbMock.GetFn = func(ctx context.Context, s string) *redis.StringCmd {
		rd := redis.NewStringCmd(t.ctx, "hero_test")
		rd.SetErr(errors.New("get hero error"))
		return rd
	}
	heroes, err := t.repo.GetAllHeroes(context.Background())
	t.NotNil(err)
	t.Empty(heroes)
}

func (t *databaseTestSuite) TestRepository_GetAllHeroes_UnmarshallError() {
	t.dbMock.GetFn = func(ctx context.Context, s string) *redis.StringCmd {
		rd := redis.NewStringCmd(t.ctx, "hero_test")
		rd.SetVal(`{"hero_index": "1""}`)
		return rd
	}
	heroes, err := t.repo.GetAllHeroes(context.Background())
	t.NotNil(err)
	t.Empty(heroes)
}

func (t *databaseTestSuite) TestRepository_GetAllHeroes_Success() {
	t.dbMock.GetFn = func(ctx context.Context, s string) *redis.StringCmd {
		rd := redis.NewStringCmd(t.ctx, "hero_test")
		switch s {
		case "hero_1":
			rd.SetVal(`{"hero_index": 1}`)
			return rd
		default:
			rd.SetVal(`{"hero_index": 2}`)
			return rd
		}
	}
	heroes, err := t.repo.GetAllHeroes(context.Background())
	t.Nil(err)
	t.NotEmpty(heroes)
}
