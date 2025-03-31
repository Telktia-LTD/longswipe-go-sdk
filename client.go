package longswipe

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ClientConfig struct {
	BaseURL    string
	PublicKey  string
	PrivateKey string
	Timeout    time.Duration
}

type Client struct {
	baseURL    string
	publicKey  string
	privateKey string
	httpClient *http.Client
}

func NewClient(config ClientConfig) *Client {
	if config.Timeout == 0 {
		config.Timeout = 10 * time.Second
	}

	return &Client{
		baseURL:    config.BaseURL,
		publicKey:  config.PublicKey,
		privateKey: config.PrivateKey,
		httpClient: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		},
	}
}

func (c *Client) doRequest(method, path string, body interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.publicKey)
	req.Header.Set("X-API-Private-Key", c.privateKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "LongSwipe-Go-SDK/v1")
	req.Header.Set("X-Forwarded-Proto", "https")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return resp, nil
}

func (c *Client) doRequestAndUnmarshal(method, path string, requestBody, responseStruct interface{}) error {
	resp, err := c.doRequest(method, path, requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(responseStruct); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
