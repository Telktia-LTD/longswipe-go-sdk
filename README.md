# Merchant SDK for Go

Official Golang SDK for Longswipe merchant integrations.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go SDK for interacting with the Longwipe Merchant Platform API. This library simplifies API integration by providing an easy-to-use, idiomatic interface for managing merchants, vouchers, and payments.

---

## **Purpose**

The purpose of this SDK is to:

- Abstract away the complexity of direct API interactions.
- Provide Go developers with a clean and consistent interface for the Merchant Platform.
- Ensure robust error handling, retry logic, and authentication.

---

## **Scope**

The SDK supports the following features:

1. **Authentication**

   - Support for API keys

2. **Core Functionality**

   - **Fetch_api**: Fetch API keys
   - **Voucher**: Manage voucher catalogs (verify voucher, fetch voucher redemption charges, redeem voucher, fetch admin charges, update admin charges).
   - **Invoice**: Create and fetch invoices.
   - **Customer**: Manage customer catalogs (CRUD operations for customers).
   - **Others**: Fetch supported currencies and supported crypto-networks.

3. **Utility Features**

   - **Pagination**: Simplified handling of paginated API responses.
   - **Retry Logic**: Automatic retry for transient network errors.
   - **Custom HTTP Client**: Support for custom HTTP configurations.
   - **Logging**: Built-in support for request and error logging.

4. **Developer Experience**
   - Thorough inline documentation.
   - Easy-to-understand examples.

---

## **Installation**

To install the SDK, use `go get`:

```bash
go get github.com/Telktia-LTD/longswipe-go-sdk
```

---

## **Examples**

### **Initialization**

Here’s an example of how to initialize the SDK and perform basic operations:

```go
package main

import (
	"fmt"
	"log"
	"github.com/Telktia-LTD/longswipe-go-sdk"
	"github.com/Telktia-LTD/longswipe-go-sdk/utils"
	"time"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    utils.PRODUCTION, // utils.SANDBOX
		PublicKey:  "YOUR_PUBLIC_API_KEY",
		Timeout:    5 * time.Second,
		PrivateKey: "YOUR_SECRET_API_KEY",
	})

	status, err := client.HealthCheck()
	if err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	fmt.Println("Service status:", status)

	// FETCH ALL NETWORKS
	networks, err := client.GetAllNetwork()
	if err != nil {
		log.Fatalf("network check failed: %v", err)
	}
	fmt.Println("Service status:", networks)

	// FETCH ALL NETWORKS
	currencies, err := client.GetAllCurrency()
	if err != nil {
		log.Fatalf("network check failed: %v", err)
	}

	fmt.Println("Service status:", currencies)
}
```

This example demonstrates:

1. Initializing the SDK with custom configurations.
2. Making an API request.
3. Utilizing the service layer for fetching an API key.

---

### **Further Example**

Here’s an example of how to work around the SDK and perform voucher operations:

```go
package main

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    utils.PRODUCTION,
		PublicKey:  "YOUR_PUBLIC_API_KEY",
		Timeout:    5 * time.Second,
		PrivateKey: "YOUR_SECRET_API_KEY",
	})

	// VERIFY VOUCHER
	res, err := client.VerifyVoucher(&utils.VerifyVoucherCodeRequest{VoucherCode: "LS3130635050"})
	if err != nil {
		log.Fatalf("verify voucher failed: %v", err)
	}
	fmt.Println("Service status:", res.Data.Balance)

	// FETCH VOUCHER CHARGES
	res, err := client.GetVoucherRedeemptionCharges(&utils.RedeemRequest{
		VoucherCode:            "LS3130635050",
		Amount:                 200,
		LockPin:                "",
		ToCurrencyAbbreviation: "NGN",
	})
	if err != nil {
		log.Fatalf("verify voucher failed: %v", err)
	}

	fmt.Println("Service status:", res.Data.Charges.TotalDeductable)

	// REDEEM VOUCHER CHARGES
	res, err := client.RedeemVoucher(&utils.RedeemRequest{
		VoucherCode:            "LS3130635050",
		Amount:                 100,
		LockPin:                "",
		ToCurrencyAbbreviation: "NGN",
		ReferenceId:            "123456789",
		MetaData:               map[string]string{
			"username": "YOUR USERNAME",
			"email": "Your users email"
		},
	})
	if err != nil {
		log.Fatalf("verify voucher failed: %v", err)
	}

	fmt.Println("Service status:", res)
}
```

This example demonstrates:

1. Initializing the SDK with custom configurations.
2. Fetching customer with pagination.
3. Adding new customer.
4. Updating customer.
5. Deleting customer.

You can refrence the example file for more examples

Documentation: https://developer.longswipe.com/docs/
