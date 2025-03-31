package longswipe

import "github.com/Telktia-LTD/longswipe-go-sdk/utils"

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

func (c *Client) HealthCheck() (*HealthCheckResponse, error) {
	var response HealthCheckResponse

	err := c.doRequestAndUnmarshal(
		utils.GET,
		"/merchant-integrations-server/health",
		nil,
		&response,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
