package longswipe

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) FetchInvoice(body *Pagination) (*MerchantInvoiceResponse, error) {
	endpoint := buildInvoiceEndpoint(body.Page, body.Limit, body.Search)
	var invoice MerchantInvoiceResponse

	err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&invoice,
	)

	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func buildInvoiceEndpoint(page int, limit int, filter string) string {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("limit", strconv.Itoa(limit))
	params.Add("filter", filter)

	return "/merchant-integrations-server/fetch-invoice?" + params.Encode()
}

func (c *Client) GetAllInvoiceCurrency() (*FetchAllAllowedInvoiceCurrencyResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/fetch-all-allowed-invoice-Currency")

	var allowedCurrency FetchAllAllowedInvoiceCurrencyResponse

	err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&allowedCurrency,
	)

	if err != nil {
		return nil, err
	}
	return &allowedCurrency, nil
}

func (c *Client) CreateInvoice(body *CreateInvoiceRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/create-invoice"
	var res SuccessResponse
	err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) ApproveInvoice(body *ApproveInvoiceRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/approve-invoice"
	var res SuccessResponse

	err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
