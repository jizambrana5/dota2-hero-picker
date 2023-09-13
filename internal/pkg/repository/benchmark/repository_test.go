package benchmark

import (
	"bytes"
	"context"
	"embed"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/repository/benchmark/mocks"
)

type BenchmarkSuite struct {
	suite.Suite
	ctx           context.Context
	requester     *mocks.RequesterMock
	repo          *Repository
	mockResponses map[string][]byte
}

//go:embed mocks/responses
var fsPay embed.FS

func (t *BenchmarkSuite) SetupSuite() {
	t.ctx = context.Background()
	t.mockResponses = make(map[string][]byte)
	dir, err := fsPay.ReadDir("mocks/responses")
	t.Nil(err)
	for _, f := range dir {
		if f.IsDir() {
			break
		}
		b, err := fsPay.ReadFile("mocks/responses/" + f.Name())
		t.Nil(err)
		s := strings.Split(f.Name(), ".")
		t.mockResponses[s[0]] = b
	}
	t.requester = &mocks.RequesterMock{}
	t.repo = NewRepository(Config{
		Timeout:  1000000,
		BasePath: "https://test.com",
		Retries:  4,
	})
	t.repo.Client = t.requester
}

func TestBenchmark(t *testing.T) {
	suite.Run(t, new(BenchmarkSuite))
}

func (t *BenchmarkSuite) Test_NewRepository() {
	t.NotNil(NewRepository(Config{
		Timeout:  0,
		BasePath: "",
		Retries:  0,
	}))
}

func (t *BenchmarkSuite) Test_GetBenchMark_Fail() {
	// Given
	r, _ := http.NewRequestWithContext(t.ctx, http.MethodGet, "", nil)
	t.requester.DoFunc = func(req *http.Request) (*http.Response, error) {
		b := io.NopCloser(bytes.NewReader([]byte("")))
		return &http.Response{
			Request:    r,
			Body:       b,
			StatusCode: 500,
		}, nil
	}
	resp, err := t.repo.GetHeroBenchmark(t.ctx, "1")
	t.NotNil(err)
	t.Empty(resp)
	t.EqualError(err, "API request failed with status code: 500")
}

func (t *BenchmarkSuite) Test_GetBenchMark_JsonError() {
	// Given
	r, _ := http.NewRequestWithContext(t.ctx, http.MethodGet, "", nil)
	t.requester.DoFunc = func(req *http.Request) (*http.Response, error) {
		b := io.NopCloser(bytes.NewReader([]byte("}{")))
		return &http.Response{
			Request:    r,
			Body:       b,
			StatusCode: 200,
		}, nil
	}
	resp, err := t.repo.GetHeroBenchmark(t.ctx, "1")
	t.NotNil(err)
	t.Empty(resp)
	t.EqualError(err, "error decoding JSON response: invalid character '}' looking for beginning of value")
}

func (t *BenchmarkSuite) Test_GetBenchMarkSuccess() {
	// Given
	r, _ := http.NewRequestWithContext(t.ctx, http.MethodGet, "", nil)
	t.requester.DoFunc = func(req *http.Request) (*http.Response, error) {
		response, ok := t.mockResponses["success"]
		t.True(ok)
		return &http.Response{
			Request:    r,
			Body:       io.NopCloser(bytes.NewReader(response)),
			StatusCode: 200,
		}, nil
	}
	resp, err := t.repo.GetHeroBenchmark(t.ctx, "1")
	t.Nil(err)
	t.NotEmpty(resp)
}
