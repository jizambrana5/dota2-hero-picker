package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain"
	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/rest/mocks"
)

type handlerSuite struct {
	suite.Suite
	ctx         *gin.Context
	heroService *mocks.HeroServiceMock
	handler     *Handler
	w           *httptest.ResponseRecorder
}

func (t *handlerSuite) SetupTest() {
	t.w = httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	t.ctx, _ = gin.CreateTestContext(t.w)
	t.ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	t.heroService = &mocks.HeroServiceMock{
		GetAllHeroesFunc: func(ctx context.Context) ([]domain.Hero, error) {
			return []domain.Hero{{
				HeroIndex:        "1",
				PrimaryAttribute: "str",
				NameInGame:       "Abbadon",
				Role:             []domain.Role{"Support"},
			}}, nil
		},
		GetDataSetFunc: func(ctx context.Context) ([][]string, error) {
			return [][]string{{"ID", "Name", "Primary Attribute", "Roles"}, {"1", "Abbado", "str", "Support"}}, nil
		},
		GetHeroFunc: func(ctx context.Context, id string) (domain.Hero, error) {
			return domain.Hero{
				HeroIndex:        "id_test",
				PrimaryAttribute: "str	",
				NameInGame:       "Abbadon",
				Role:             []domain.Role{"Support"},
			}, nil
		},
		GetHeroSuggestionFunc: func(ctx context.Context, preferences domain.UserPreferences) ([]domain.Hero, error) {
			return []domain.Hero{{
				HeroIndex:        "1",
				PrimaryAttribute: "str",
				NameInGame:       "Abbadon",
				Role:             []domain.Role{"Support"},
			}}, nil
		},
		SaveHeroesFunc: func(ctx context.Context) error {
			return nil
		},
	}
	t.handler = NewHandler(t.heroService)
}

func (t *handlerSuite) Test_NewHandler() {
	t.NotNil(NewHandler(t.heroService))
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func MockRequest(c *gin.Context, params gin.Params, u url.Values, method string) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}
