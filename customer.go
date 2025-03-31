package longswipe

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) GetCustomers(body *Pagination) (*CustomersResponse, error) {
	endpoint := buildCustomerEndpoint(body.Page, body.Limit, body.Search)
	var customers CustomersResponse

	err := c.doRequestAndUnmarshal(
		GET,
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

func (c *Client) GetCustomer(email string) (*CustomerResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/fetch-customer-by-email/%s", email)
	fmt.Println(endpoint)

	var customer CustomerResponse

	err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&customer,
	)

	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *Client) AddCustomer(body *AddNewCustomer) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/add-new-customer"
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

func (c *Client) UpdateCustomer(body *UpdatCustomer) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/update-customer"
	var res SuccessResponse

	err := c.doRequestAndUnmarshal(
		PATCH,
		endpoint,
		body,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteCustomer(customerID string) (*SuccessResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/delete-customer/%s", customerID)
	var res SuccessResponse

	err := c.doRequestAndUnmarshal(
		DELETE,
		endpoint,
		nil,
		&res,
	)

	if err != nil {
		return nil, err
	}
	return &res, nil
}
