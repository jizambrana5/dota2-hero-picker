package rest

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/lib/errors"
)

func (t *handlerSuite) Test_GetHero_ServiceError() {
	t.heroService.GetHeroFunc = func(ctx context.Context, id string) (domain.Hero, error) {
		return domain.Hero{}, errors.HeroNotFound
	}
	params := []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	MockRequest(t.ctx, params, url.Values{}, http.MethodGet)
	t.handler.GetHero(t.ctx)
	t.Equal(http.StatusNotFound, t.w.Code)
	t.Equal("{\"code\":\"hero_not_found\",\"message\":\"hero not found\"}", t.w.Body.String())

}

func (t *handlerSuite) Test_GetHero_Success() {
	//configure path params
	params := []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	MockRequest(t.ctx, params, url.Values{}, http.MethodGet)
	t.handler.GetHero(t.ctx)
	t.Equal(http.StatusOK, t.w.Code)
	t.Equal("{\"hero_index\":\"id_test\",\"primary_attr\":\"str\\t\",\"localized_name\":\"Abbadon\",\"roles\":[\"Support\"],\"win_rates\":null}", t.w.Body.String())

}
