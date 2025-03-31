package longswipe

import "github.com/Telktia-LTD/longswipe-go-sdk/utils"

func (c *Client) GetVoucherRedeemptionCharges(body *utils.RedeemRequest) (*utils.RedeemeVoucherDetailsResponse, error) {
	endpoint := "/merchant-integrations/fetch-voucher-redemption-charges"
	var charges utils.RedeemeVoucherDetailsResponse
	err := c.doRequestAndUnmarshal(
		utils.POST,
		endpoint,
		body,
		&charges,
	)

	if err != nil {
		return nil, err
	}
	return &charges, nil
}

func (c *Client) VerifyVoucher(body *utils.VerifyVoucherCodeRequest) (*utils.VerifyVoucherResponse, error) {
	var verifyVoucher utils.VerifyVoucherResponse

	endpoint := "/merchant-integrations/verify-voucher"
	err := c.doRequestAndUnmarshal(
		utils.POST,
		endpoint,
		body,
		&verifyVoucher,
	)

	if err != nil {
		return nil, err
	}
	return &verifyVoucher, nil
}

func (c *Client) RedeemVoucher(body *utils.RedeemRequest) (*utils.SuccessResponse, error) {
	endpoint := "/merchant-integrations/redeem-voucher"
	var redeemVoucher utils.SuccessResponse

	err := c.doRequestAndUnmarshal(
		utils.POST,
		endpoint,
		body,
		&redeemVoucher,
	)

	if err != nil {
		return nil, err
	}
	return &redeemVoucher, nil
}

func (c *Client) GenerateVoucher(body *utils.GenerateVoucherForCustomerRequest) (*utils.SuccessResponse, error) {
	endpoint := "/merchant-integrations-server/generate-voucher-for-customer"
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
