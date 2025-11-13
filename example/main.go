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

	escrowId := "ESC01K9VZK9P5QCH55FGK0G5YFS6N"
	response, err := client.FetchEscrowDetails(escrowId)
	if err != nil {
		fmt.Println("Error fetching escrow details:", err)
		return
	}

	fmt.Println("Escrow Details Response:", response.Data)
	// Use response
}
