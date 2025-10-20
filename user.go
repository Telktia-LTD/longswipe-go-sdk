package longswipe

func (c *Client) AddUser(body *AddNewUserRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/add-new-user"
	var res SuccessResponse

	_, err := c.doRequestAndUnmarshal(
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

func (c *Client) GetAllUser() (*MerchantUserResponse, error) {
	endpoint := "/merchant-integrations-server/fetch-merchant-users"
	var user MerchantUserResponse

	_, err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&user,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
