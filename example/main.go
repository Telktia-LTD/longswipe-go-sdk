package main

import (
	"fmt"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    "http://localhost:8888",
		PublicKey:  "pk_live_lpKM1dGecmlxkZrqrTmNErz2EnRNjog5-EBYzLiswDI=",
		PrivateKey: "sk_live_N8TrEXa-4PtQo2AyIkW8Qg2wcQQ9N8Oa2_yrAUyH7Dc=",
		Timeout:    30 * time.Second,
	})

	res, err := client.AccountBalance("USDT")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Payout Response:", res.Message)
	fmt.Println("Merchant:", res.Data.Merchant)
	fmt.Println("Balance:", res.Data.Balance)
	fmt.Println("Currency:", res.Data.Currency)

	// Use response
}
