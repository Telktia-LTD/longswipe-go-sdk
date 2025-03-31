package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
	"github.com/Telktia-LTD/longswipe-go-sdk/utils"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    utils.SANDBOX,
		PublicKey:  "pk_live_odmenqa8RfE1H1E2O1C_VH8Bq9pEaxnYWLme7Rpi48E=", // call from your env variable
		Timeout:    5 * time.Second,
		PrivateKey: "sk_live_DRKlfAZwSiLYP6Vkqzomhj6HsBgaEfzzzOFZ1Jrx-Xo=", // call from your env variable
	})
	// FETCH ALL NETWORKS
	currencies, err := client.GetAllCurrency()
	if err != nil {
		log.Fatalf("network check failed: %v", err)
	}
	fmt.Println("Service status:", currencies.Data.Currencies)
}
