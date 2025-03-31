package longswipe

import "github.com/Telktia-LTD/longswipe-go-sdk/utils"

func (c *Client) GetAllNetwork() (*utils.CryptoNetworkResponse, error) {
	endpoint := "/merchant-integrations/fetch-supported-cryptonetworks"
	var networks utils.CryptoNetworkResponse

	err := c.doRequestAndUnmarshal(
		utils.GET,
		endpoint,
		nil,
		&networks,
	)

	if err != nil {
		return nil, err
	}
	return &networks, nil
}

func (c *Client) GetAllCurrency() (*utils.FetchCurrenciesResponse, error) {
	endpoint := "/merchant-integrations/fetch-supported-currencies"
	var currencies *utils.FetchCurrenciesResponse
	err := c.doRequestAndUnmarshal(
		utils.GET,
		endpoint,
		nil,
		&currencies,
	)

	if err != nil {
		return nil, err
	}
	return currencies, nil
}
