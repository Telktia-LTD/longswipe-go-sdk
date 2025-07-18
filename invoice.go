package longswipe

<<<<<<< HEAD
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
=======
import "fmt"
>>>>>>> a11357c (invoice)

func (c *Client) CreateInvoice(body *CreateInvoiceRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/create-invoice"
	var res SuccessResponse
<<<<<<< HEAD

=======
>>>>>>> a11357c (invoice)
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

<<<<<<< HEAD
func (c *Client) ApproveInvoice(body *ApproveInvoiceRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/approve-invoice"
	var res SuccessResponse

	err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&res,
=======
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
>>>>>>> a11357c (invoice)
	)

	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	return &res, nil
=======
	return &response, nil
>>>>>>> a11357c (invoice)
}
