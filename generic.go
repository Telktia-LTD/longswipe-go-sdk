package longswipe

func (c *Client) GetAllNetwork() (*CryptoNetworkResponse, error) {
	endpoint := "/merchant-integrations/fetch-supported-cryptonetworks"
	var networks CryptoNetworkResponse

	_, err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&networks,
	)

	if err != nil {
		return nil, err
	}
	return &networks, nil
}

func (c *Client) GetAllCurrency() (*FetchCurrenciesResponse, error) {
	endpoint := "/merchant-integrations/fetch-supported-currencies"
	var currencies FetchCurrenciesResponse
	_, err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&currencies,
	)

	if err != nil {
		return nil, err
	}
	return &currencies, nil
}
