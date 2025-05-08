package main

import (
	"fmt"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    longswipe.PRODUCTION,
		PublicKey:  "pk_live_uQxOuuP-ka03sqIFgBBqbaYL9-vGkso9BOfpXUdVSIY=", // call from your env variable
		Timeout:    5 * time.Second,
		PrivateKey: "sk_live_H-CqZ9PivCpiG9KQDyIp_G0kTYm8MH0cMpIg6CKYAVI=", // call from your env variable
	})

	// FETCH ALL NETWORKS
	// currencies, err := client.GetAllCurrency()
	// if err != nil {
	// 	log.Fatalf("network check failed: %v", err)
	// }
	// fmt.Println("Service status:", currencies.Data.Currencies)

	// ADD CUSTOMER
	res, err := client.AddCustomer(&longswipe.AddNewCustomer{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	})

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("response", res)

}
