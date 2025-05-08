package main

import (
	"fmt"
	"log"
	"time"

	longswipe "github.com/Telktia-LTD/longswipe-go-sdk"
	"github.com/gofrs/uuid"
)

func main() {
	// Initialize the client
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    longswipe.PRODUCTION,
		PublicKey:  "pk_live_uQxOuuP-ka03sqIFgBBqbaYL9-vGkso9BOfpXUdVSIY=",
		Timeout:    10 * time.Second,
		PrivateKey: "sk_live_H-CqZ9PivCpiG9KQDyIp_G0kTYm8MH0cMpIg6CKYAVI=",
	})

	// currency, err := client.GetAllCurrency()
	// if err != nil {
	// 	log.Fatalf("Error fetching currency: %v", err)

	// }

	// users, err := client.GetAllUser()
	// if err != nil {
	// 	log.Fatalf("Error fetching user: %v", err)

	// }

	// fmt.Println("currencies", currency)
	// fmt.Println("======================================")
	// fmt.Println("users", users)

	// CREATE INVOICE
	merchantID, _ := uuid.FromString("ccdf5701-be9a-4311-a421-45e241e36d8a")
	currencyID, _ := uuid.FromString("9a0970d2-6680-45b3-85c1-7f3c189807e7")

	createInvoiceRequest := &longswipe.CreateInvoiceRequest{
		MerchantUserId: merchantID,
		InvoiceDate:    time.Now(),
		DueDate:        time.Now().AddDate(0, 0, 30), // Due in 30 days
		CurrencyId:     currencyID,
		InvoiceItems: []longswipe.InvoiceItemRequest{
			{
				Description: "Software Development Services",
				Quantity:    1,
				UnitPrice:   1000.00,
			},
			{
				Description: "Maintenance Fee",
				Quantity:    2,
				UnitPrice:   250.00,
			},
		},
	}

	createResponse, err := client.CreateInvoice(createInvoiceRequest)
	if err != nil {
		log.Fatalf("Error creating invoice: %v", err)
	}
	fmt.Printf("Invoice Creation Response: %s (Code: %d)\n", createResponse.Message, createResponse.Code)

	// FETCH ALL INVOICE
	// paginationParams := &longswipe.Pagination{
	// 	Page:   1,
	// 	Limit:  10,
	// 	Search: "INV-",
	// }

	// invoices, err := client.FetchInvoice(paginationParams)
	// if err != nil {
	// 	log.Fatalf("Error fetching invoices: %v", err)
	// }
	// fmt.Printf("Fetched Invoices: Total=%d, Page=%d\n", invoices.Data.Total, paginationParams.Page)

	// Example 2: Get all allowed invoice currencies
	// currencies, err := client.GetAllInvoiceCurrency()
	// if err != nil {
	// 	log.Fatalf("Error fetching invoice currencies: %v", err)
	// }
	// for _, currency := range currencies.Data {
	// 	fmt.Printf("Currency: %s (Enabled: %v)\n", currency.Currency.Name, currency.IsEnabled)
	// }

	// Example 3: Create a new invoice

	// // Example 4: Approve an invoice
	// invoiceID, _ := uuid.FromString("98765432-5678-5678-5678-987654321abc") // Replace with actual invoice ID
	// approveInvoiceRequest := &longswipe.ApproveInvoiceRequest{
	// 	InvoiceID: invoiceID,
	// 	OnChain:   true, // Set to true if you want blockchain processing
	// }

	// approveResponse, err := client.ApproveInvoice(approveInvoiceRequest)
	// if err != nil {
	// 	log.Fatalf("Error approving invoice: %v", err)
	// }
	// fmt.Printf("Invoice Approval Response: %s (Status: %s)\n", approveResponse.Message, approveResponse.Status)
}
