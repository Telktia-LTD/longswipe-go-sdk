package longswipe

import "fmt"

func (c *Client) InitiateEscrow(body *EscrowPublicRequest) (*EscrowInitialResponse, error) {
	endpoint := "/merchant-integrations-server/initiate-escrow"
	var result EscrowInitialResponse

	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) UpdateEscrow(body *UpdateEscrowStatusRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/update-escrow"
	var result SuccessResponse

	_, err := c.doRequestAndUnmarshal(
		PATCH,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) FundRequest(escrowId string) (*FundRequestResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/escrow-fund-request/%s", escrowId)
	var result FundRequestResponse

	_, err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) FetchAllEscrow() (*AllEscrowResponse, error) {
	endpoint := "/merchant-integrations-server/fetch-all-escrows"
	var result AllEscrowResponse

	_, err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) FetchEscrowDetails(escrowId string) (*EscrowDetailResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/escrow-details/%s", escrowId)
	var result EscrowDetailResponse

	_, err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) RequestOtp(body RequestOtp) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/escrow-request-otp"
	var result SuccessResponse

	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) AddRecipient(body AddEscrowRecipient) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/add-escrow-recipient"
	var result SuccessResponse

	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteRecipient(escrowId, email string) (*SuccessResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/delete-escrow-user/%s?email=%s", escrowId, email)

	var result SuccessResponse
	_, err := c.doRequestAndUnmarshal(
		DELETE,
		endpoint,
		nil,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) RequestFundRelease(body RequestFundReleaseRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/escrow-fund-release-request"

	var result SuccessResponse
	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) FundRelease(body ConfirmFundRelease) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/escrow-fund-release"

	var result SuccessResponse
	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SystemRelease(body SystemRelease) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/escrow-system-fund-release"

	var result SuccessResponse
	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ResetAuthorizationCode(body UpdateAuthorizationCodeRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/escrow-reset-authorization-code"

	var result SuccessResponse
	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) AddRecipientPayoutDetails(body AddPayoutDetailsRequest) (*RecipientAccountDetailsResponse, error) {
	endpoint := "/merchant-integrations-server/escrow-recipient-payout-details"

	var result RecipientAccountDetailsResponse
	_, err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&result,
	)

	if err != nil {
		return nil, err
	}
	return &result, nil
}
