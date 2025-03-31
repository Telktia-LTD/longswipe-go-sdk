package longswipe

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/Telktia-LTD/longswipe-go-sdk/utils"
)

func (c *Client) GetCustomers(body *utils.Pagination) (*utils.CustomersResponse, error) {
	endpoint := buildCustomerEndpoint(body.Page, body.Limit, body.Search)
	var customers utils.CustomersResponse

	err := c.doRequestAndUnmarshal(
		utils.GET,
		endpoint,
		nil,
		&customers,
	)

	if err != nil {
		return nil, err
	}
	return &customers, nil
}

func buildCustomerEndpoint(page int, limit int, search string) string {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("limit", strconv.Itoa(limit))
	params.Add("search", search)

	return "/merchant-integrations-server/fetch-customers?" + params.Encode()
}

func (c *Client) GetCustomer(email string) (*utils.CustomerResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/fetch-customer-by-email/%s", email)
	fmt.Println(endpoint)

	var customer utils.CustomerResponse

	err := c.doRequestAndUnmarshal(
		utils.GET,
		endpoint,
		nil,
		&customer,
	)

	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *Client) AddCustomer(body *utils.AddNewCustomer) (*utils.SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/add-new-customer"
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

func (c *Client) UpdateCustomer(body *utils.UpdatCustomer) (*utils.SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/update-customer"
	var res utils.SuccessResponse

	err := c.doRequestAndUnmarshal(
		utils.PATCH,
		endpoint,
		body,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteCustomer(customerID string) (*utils.SuccessResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/delete-customer/%s", customerID)
	var res utils.SuccessResponse

	err := c.doRequestAndUnmarshal(
		utils.DELETE,
		endpoint,
		nil,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
