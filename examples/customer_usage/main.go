package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    longswipe.SANDBOX,
		PublicKey:  "pk_live_odmenqa8RfE1H1E2O1C_VH8Bq9pEaxnYWLme7Rpi48E=",
		Timeout:    10 * time.Second,
		PrivateKey: "sk_live_DRKlfAZwSiLYP6Vkqzomhj6HsBgaEfzzzOFZ1Jrx-Xo=",
	})

	// FETCH CUSTOMERS
	res, err := client.GetCustomers(&longswipe.Pagination{
		Page:   1,
		Limit:  20,
		Search: "",
	})
	if err != nil {
		log.Fatalf("fetch customers failed: %v", err)
	}
	fmt.Println("Service status:", res)
}
