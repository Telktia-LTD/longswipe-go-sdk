package longswipe

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Telktia-LTD/longswipe-go-sdk/utils"
)

func (c *Client) FetchInvoice(body *utils.Pagination) (*utils.MerchantInvoiceResponse, error) {
	endpoint := buildInvoiceEndpoint(body.Page, body.Limit, body.Search)
	fmt.Println("======1", endpoint)
	var invoice utils.MerchantInvoiceResponse

	err := c.doRequestAndUnmarshal(
		utils.GET,
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

func (c *Client) GetAllInvoiceCurrency() (*utils.FetchAllAllowedInvoiceCurrencyResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/fetch-all-allowed-invoice-Currency")
	fmt.Println("======2", endpoint)

	var allowedCurrency utils.FetchAllAllowedInvoiceCurrencyResponse

	err := c.doRequestAndUnmarshal(
		utils.GET,
		endpoint,
		nil,
		&allowedCurrency,
	)

	if err != nil {
		return nil, err
	}
	return &allowedCurrency, nil
}

func (c *Client) CreateInvoice(body *utils.CreateInvoiceRequest) (*utils.SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/create-invoice"
	var res utils.SuccessResponse

	err := c.doRequestAndUnmarshal(
		utils.POST,
		endpoint,
		body,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) ApproveInvoice(body *utils.ApproveInvoiceRequest) (*utils.SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/approve-invoice"
	var res utils.SuccessResponse

	err := c.doRequestAndUnmarshal(
		utils.POST,
		endpoint,
		body,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
