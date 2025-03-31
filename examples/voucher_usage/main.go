package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
	"github.com/Telktia-LTD/longswipe-go-sdk/utils"
	"github.com/gofrs/uuid"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    utils.SANDBOX,
		PublicKey:  "pk_live_odmenqa8RfE1H1E2O1C_VH8Bq9pEaxnYWLme7Rpi48E=",
		Timeout:    30 * time.Second,
		PrivateKey: "sk_live_DRKlfAZwSiLYP6Vkqzomhj6HsBgaEfzzzOFZ1Jrx-Xo=",
	})
	// GENERATE VOUCHER
	currencyID, err := uuid.FromString("9a0470d3-fh80-45b3-85c1-7f3c1455401d6")
	if err != nil {
		log.Fatalf("invalid currency ID: %v", err)
	}

	customerID, err := uuid.FromString("08408121-ft72-4348-bdfd-41cf1c2f9e0d")
	if err != nil {
		log.Fatalf("invalid customer ID: %v", err)
	}

	res, err := client.GenerateVoucher(&utils.GenerateVoucherForCustomerRequest{
		CurrencyId:       currencyID,
		AmountToPurchase: 150.00,
		CustomerID:       customerID,
		// OnChain:          false,
	})
	if err != nil {
		log.Fatalf("generate voucher failed: %v", err)
	}
	fmt.Println("Service status:", res)

}
