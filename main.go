package longswipe

import (
	"fmt"
	"time"
)

func main() {
	// Example usage of the LongSwipe client
	config := LongSwipeConfig{
		BaseURL:    "https://api.longswipe.com",
		PublicKey:  "your_public_key",
		PrivateKey: "your_private_key",
		Timeout:    10 * time.Second,
	}
	client := NewClient(config)

	// client.GetVoucherRedeemptionCharges(&RedeemRequest{
	// 	VoucherCode:            "LS3263655440",
	// 	Amount:                 100.0,
	// 	ToCurrencyAbbreviation: "NGN",
	// 	ReferenceId:            "ref123",
	// 	MetaData: map[string]string{
	// 		"key1": "value1",
	// 		"key2": "value2",
	// 	},
	// })
	res, err := client.VerifyVoucher(&VerifyVoucherCodeRequest{
		VoucherCode: "LS3263655440",
	})

	if err != nil {
		// Handle error
		fmt.Println("Error verifying voucher:", err)
		return
	}
	fmt.Println("Voucher verification response:", res.Data.Amount)

}
