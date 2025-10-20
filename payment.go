package longswipe

import "fmt"

func (c *Client) PaymentRequest(body *PaymentRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations/payment-request"
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

func (c *Client) AddressDepositRequest(body *AddressDepositRequest) (*DepositResponse, error) {
	endpoint := "/merchant-integrations/deposit-address-payment-request"
	var res DepositResponse

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

func (c *Client) DepositCharges(body *AddressDepositChargeRequest) (*ChargeEstimateResponse, error) {
	endpoint := "/merchant-integrations/request-wallet-deposit-charges"
	var charges ChargeEstimateResponse

	err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&charges,
	)

	if err != nil {
		return nil, err
	}
	return &charges, nil
}

func (c *Client) VerifyTransaction(referenceId string) (*TransactionResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations-server/verify-transaction/%s", referenceId)
	var transaction TransactionResponse

	err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&transaction,
	)

	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (c *Client) ConfirmUser(identifier string) (*ConfirmUserDetailsResponse, error) {
	endpoint := fmt.Sprintf("/merchant-integrations/confirm-user/%s", identifier)
	var userProfile ConfirmUserDetailsResponse

	err := c.doRequestAndUnmarshal(
		GET,
		endpoint,
		nil,
		&userProfile,
	)

	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (c *Client) PayoutToLongSwipeUser(body *CustomerPayout) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/payout"
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
