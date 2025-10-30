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

// doRequest performs the HTTP request and returns the status code and body bytes.
// It never returns a non-nil *http.Response (to avoid leaking bodies); instead it
// reads the body fully and returns the bytes so callers can decide how to handle it.
func (c *Client) doRequest(method, path string, body interface{}) (int, []byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.publicKey)
	req.Header.Set("X-API-Private-Key", c.privateKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "LongSwipe-Go-SDK/v1")
	req.Header.Set("X-Forwarded-Proto", "https")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		// Try to parse a JSON error with a `message` field and return that message
		var parsed map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &parsed); err == nil {
			if msg, ok := parsed["message"].(string); ok && msg != "" {
				return resp.StatusCode, bodyBytes, fmt.Errorf("%s", msg)
			}
		}

		// fallback: return the raw body string
		return resp.StatusCode, bodyBytes, fmt.Errorf("%s", string(bodyBytes))
	}

	return resp.StatusCode, bodyBytes, nil
}

func (c *Client) doRequestAndUnmarshal(method, path string, requestBody, responseStruct interface{}) (int, error) {
	status, bodyBytes, err := c.doRequest(method, path, requestBody)
	if err != nil {
		// even on error we may have bodyBytes with API message; return status and error
		return status, err
	}

	if responseStruct == nil {
		// caller doesn't want the body decoded
		return status, nil
	}

	if err := json.Unmarshal(bodyBytes, responseStruct); err != nil {
		return status, fmt.Errorf("failed to decode response: %w", err)
	}

	return status, nil
}
