package longswipe

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

func (c *Client) HealthCheck() (*HealthCheckResponse, error) {
	var response HealthCheckResponse

	err := c.doRequestAndUnmarshal(
		GET,
		"/merchant-integrations-server/health",
		nil,
		&response,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
