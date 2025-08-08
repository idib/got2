package client

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/idib/got2/internal/server"
)

// Client represents the HTTP client for the server
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new client instance
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Ping sends a request to the /ping endpoint
func (c *Client) Ping() (string, error) {
	url := fmt.Sprintf("%s/ping", c.baseURL)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to send ping request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func (c *Client) HandleChekH(_ server.HandleChekHRequest) (server.HandleChekHResponse, error) {
	url := fmt.Sprintf("%s/handleChekH", c.baseURL)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return server.HandleChekHResponse{}, fmt.Errorf("failed to send ping request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return server.HandleChekHResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return server.HandleChekHResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	return server.HandleChekHResponse{}, nil
}
