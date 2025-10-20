// Package longswipe provides a Go SDK for the LongSwipe API.
//
// This SDK allows you to integrate LongSwipe functionality into your Go applications.
// Official documentation for the LongSwipe API can be found at https://developer.longswipe.com/docs/
//
// Example usage:
//
// func main() {
// 	client := longswipe.NewClient(longswipe.ClientConfig{
// 		BaseURL:    utils.PRODUCTION, // utils.SANDBOX
// 		PublicKey:  "YOUR_PUBLIC_API_KEY",
// 		Timeout:    5 * time.Second,
// 		PrivateKey: "YOUR_SECRET_API_KEY",
// 	})

// 	status, err := client.HealthCheck()
// 	if err != nil {
// 		log.Fatalf("Health check failed: %v", err)
// 	}

// 	fmt.Println("Service status:", status)

// 	// FETCH ALL NETWORKS
// 	networks, err := client.GetAllNetwork()
// 	if err != nil {
// 		log.Fatalf("network check failed: %v", err)
// 	}
// 	fmt.Println("Service status:", networks)

// 	// FETCH ALL NETWORKS
// 	currencies, err := client.GetAllCurrency()
// 	if err != nil {
// 		log.Fatalf("network check failed: %v", err)
// 	}

//		fmt.Println("Service status:", currencies)
//	}
//
// For more information, visit: https://github.com/Telktia-LTD/longswipe-go-sdk
package longswipe
