package longswipe

import "fmt"

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

func (c *Client) GetInvoice(query QueryParams) (*InvoiceResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/fetch-invoice?page=%d&limit=%d&filter=%s", query.Page, query.Limit, query.Filter)
	var response InvoiceResponse
	err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		query,
		&response,
	)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ApproveInvoice(request *ApproveInvoice) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/approve-invoice"
	var response SuccessResponse
	err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		request,
		&response,
	)

	if err != nil {
		return nil, err
	}
	return &response, nil
}
