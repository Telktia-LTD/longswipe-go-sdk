package longswipe

func (c *Client) GetVoucherRedeemptionCharges(body *RedeemRequest) (*RedeemeVoucherDetailsResponse, error) {
	endpoint := "/merchant-integrations/fetch-voucher-redemption-charges"
	var charges RedeemeVoucherDetailsResponse
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

func (c *Client) VerifyVoucher(body *VerifyVoucherCodeRequest) (*VerifyVoucherResponse, error) {
	var verifyVoucher VerifyVoucherResponse

	endpoint := "/merchant-integrations/verify-voucher"
	err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&verifyVoucher,
	)

	if err != nil {
		return nil, err
	}
	return &verifyVoucher, nil
}

func (c *Client) RedeemVoucher(body *RedeemRequest) (*SuccessResponse, error) {
	endpoint := "/merchant-integrations/redeem-voucher"
	var redeemVoucher SuccessResponse

	err := c.doRequestAndUnmarshal(
		POST,
		endpoint,
		body,
		&redeemVoucher,
	)

	if err != nil {
		return nil, err
	}
	return &redeemVoucher, nil
}

<<<<<<< HEAD
func (c *Client) GenerateVoucherForCustomer(body *GenerateVoucherForCustomerRequest) (*SuccessResponse, error) {
=======
func (c *Client) GenerateVoucher(body *GenerateVoucherForCustomerRequest) (*SuccessResponse, error) {
>>>>>>> a11357c (invoice)
	endpoint := "/merchant-integrations-server/generate-voucher-for-customer"
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
