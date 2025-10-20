package longswipe

func (c *Client) GetVoucherRedeemptionCharges(body *RedeemRequest) (*RedeemeVoucherDetailsResponse, error) {
	endpoint := "/merchant-integrations/fetch-voucher-redemption-charges"
	var charges RedeemeVoucherDetailsResponse
	_, err := c.doRequestAndUnmarshal(
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
	_, err := c.doRequestAndUnmarshal(
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

	_, err := c.doRequestAndUnmarshal(
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
