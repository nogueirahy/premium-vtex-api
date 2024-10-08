package vtex

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Client struct {
	HttpClient *http.Client
	baseURL    string
}

func (c *Client) SearchProductsBySkuIds(
	ctx context.Context,
	input ReqSearchProducts,
) (*ResSearchProducts, error) {

	fullURL := fmt.Sprintf("%s/catalog_system/pub/products/search?%s", c.baseURL, input.Query)
	slog.Info("request", "url", fullURL)

	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var results *ResSearchProducts
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling response body %w", err)
	}

	return results, nil
}

func (c *Client) GetPrimeConfig(ctx context.Context) (*ResPrimeConfig, error) {
	fullURL := os.Getenv("PREMIUM_URL")
	slog.Info("request", "url", fullURL)

	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var results *ResPrimeConfig
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling response body %w", err)
	}

	return results, nil
}

func NewVtexClient(baseURL string) *Client {
	return &Client{
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
	}
}
