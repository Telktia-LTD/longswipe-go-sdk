package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
	"github.com/gofrs/uuid"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    longswipe.PRODUCTION,
		PublicKey:  "pk_live_KIJneg1IWF0r6S6VeY8FfRuZckq2jmiBcOvgYA-hYb8=",
		Timeout:    100 * time.Second,
		PrivateKey: "sk_live_vxCrmEqSorXlcG99U9sFcRIPezmj0eKQpg_I-EBXZ50=",
	})
	// VERIFY VOUCHER
	res, err := client.VerifyVoucher(&longswipe.VerifyVoucherCodeRequest{
		VoucherCode: "LS9115933307",
	})
	if err != nil {
		log.Fatalf("verify voucher failed: %v", err)
	}
	fmt.Println("Service status:", res.Data.Balance)

	// FETCH VOUCHER CHARGE
	voucherRes, err := client.GetVoucherRedeemptionCharges(&longswipe.RedeemRequest{
		VoucherCode:            "LS9115933307",
		Amount:                 100.00,
		LockPin:                "4617",
		ToCurrencyAbbreviation: "NGN",
	})
	if err != nil {
		log.Fatalf("fetch voucher charges failed: %v", err)
	}
	fmt.Println("Voucher charges:", voucherRes.Data.Voucher.GeneratedCurrency.ID)

	// REDEEM VOUCHER
	redeemRes, err := client.RedeemVoucher(&longswipe.RedeemRequest{
		VoucherCode:            "LS9115933307",
		Amount:                 500.00,
		LockPin:                "4617",
		ToCurrencyAbbreviation: "NGN",
	})
	if err != nil {
		log.Fatalf("redeem voucher failed: %v", err)
	}
	fmt.Println("Voucher redeemed:", redeemRes)

	// GENERATE VOUCHER FOR CUSTOMER
	currencyID, err := uuid.FromString("9a0970d2-6680-45b3-85c1-7f3c189807e7")
	if err != nil {
		log.Fatalf("invalid currency ID: %v", err)
	}

	customerID, err := uuid.FromString("c52e511b-1f86-43ca-baf7-817d530a7c86")
	if err != nil {
		log.Fatalf("invalid customer ID: %v", err)
	}

	generateVoucher, err := client.GenerateVoucherForCustomer(&longswipe.GenerateVoucherForCustomerRequest{
		CurrencyId:       currencyID,
		AmountToPurchase: 500.00,
		CustomerID:       customerID,
	})
	if err != nil {
		log.Fatalf("generate voucher failed: %v", err)
	}
	fmt.Println("Service status:", generateVoucher)

}
