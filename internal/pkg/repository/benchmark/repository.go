//go:generate moq -pkg mocks -out ./mocks/benchmark_mock.go -skip-ensure . Requester
package benchmark

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"

	"github.com/jizambrana5/dota2-hero-picker/internal/pkg/domain/hero"
)

type (
	Requester heimdall.Client
	Config    struct {
		Timeout  time.Duration `json:"timeout"`
		BasePath string        `json:"base_path"`
		Retries  int           `json:"retries"`
	}

	Repository struct {
		BasePath string
		Client   heimdall.Client
	}
)

var _ hero.Benchmark = (*Repository)(nil)

func NewRepository(config Config) *Repository {
	initalTimeout := 2 * time.Millisecond         // Inital timeout
	maxTimeout := 90000000 * time.Millisecond     // Max time out
	exponentFactor := 2.0                         // Multiplier
	maximumJitterInterval := 2 * time.Millisecond // Max jitter interval. It must be more than 1*time.Millisecond

	// Retrier
	backoff := heimdall.NewExponentialBackoff(initalTimeout, maxTimeout, exponentFactor, maximumJitterInterval)

	// Create a new retry mechanism with the backoff
	retrier := heimdall.NewRetrier(backoff)

	// Create a new client, sets the retry mechanism, and the number of times you would like to retry
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(config.Timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(config.Retries),
	)

	return &Repository{Client: client, BasePath: config.BasePath}
}

func (r Repository) GetHeroBenchmark(ctx context.Context, id string) (interface{}, error) {
	// Create a URL object and add query parameters
	url, err := url.Parse(r.BasePath)
	if err != nil {
		// Handle error
		return nil, fmt.Errorf("error making API url: %v", err)
	}

	// Add the hero ID as a query parameter
	query := url.Query()
	query.Add("hero_id", id) // Replace "123" with the actual hero ID
	url.RawQuery = query.Encode()

	// Create an HTTP request
	req, _ := http.NewRequest("GET", url.String(), nil)

	// Make the request using Heimdall
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	// Handle the response
	if resp.StatusCode == http.StatusOK {
		// Create a map to store the JSON response
		var data map[string]interface{}

		// Parse the JSON response into the map
		if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("error decoding JSON response: %v", err)
		}

		// Now you can return the data as an interface{}
		return data, nil
	} else {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}
}
