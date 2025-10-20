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

	res, err := client.PayoutToLongSwipeUser(&longswipe.CustomerPayout{
		Amount:                        5,
		MetaData:                      "{\"order_id\":\"6735\"}",
		ToCurrencyAbbreviation:        "USDT",
		FromCurrencyAbbreviation:      "USDT",
		LongswipeUsernameOrEmail:      "festechdev@gmail.com",
		ReferenceId:                   "payout-ref-" + time.Now().Format("20060102150405"),
		BlockchainNetworkAbbreviation: "ETH",
	})

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Payout Response:", res.Message)

	// Use response
}
