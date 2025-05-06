package main

import (
	"fmt"
	"log"
	"time"

	longswipe "github.com/Telktia-LTD/longswipe-go-sdk"
	"github.com/Telktia-LTD/longswipe-go-sdk/utils"
)

func main() {
	// Initialize the client
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    "http://localhost:3003",
		PublicKey:  "pk_live_yETxk8J_16k4BO4zHPhSzkX0TVTGa6dnYS8fSjijIbo=",
		Timeout:    10 * time.Second,
		PrivateKey: "sk_live_NsHt0PkXfVcKkK9vAhN8UXSgL8a1Df43egDFUq0sjQA=",
	})
	// Example 1: Fetch invoices with pagination
	paginationParams := &utils.Pagination{
		Page:   1,
		Limit:  10,
		Search: "INV-", // Search for invoices starting with INV-
	}

	invoices, err := client.FetchInvoice(paginationParams)
	if err != nil {
		log.Fatalf("Error fetching invoices: %v", err)
	}
	fmt.Printf("Fetched Invoices: Total=%d, Page=%d\n", invoices.Data.Total, paginationParams.Page)

	// Example 2: Get all allowed invoice currencies
	currencies, err := client.GetAllInvoiceCurrency()
	if err != nil {
		log.Fatalf("Error fetching invoice currencies: %v", err)
	}
	for _, currency := range currencies.Data {
		fmt.Printf("Currency: %s (Enabled: %v)\n", currency.Currency.Name, currency.IsEnabled)
	}

	// Example 3: Create a new invoice
	// merchantID, _ := uuid.FromString("12345678-1234-1234-1234-123456789abc") // Replace with actual merchant ID
	// currencyID, _ := uuid.FromString("87654321-4321-4321-4321-987654321def") // Replace with actual currency ID

	// createInvoiceRequest := &utils.CreateInvoiceRequest{
	// 	MerchantUserId: merchantID,
	// 	InvoiceDate:    time.Now(),
	// 	DueDate:        time.Now().AddDate(0, 0, 30), // Due in 30 days
	// 	CurrencyId:     currencyID,
	// 	InvoiceItems: []utils.InvoiceItemRequest{
	// 		{
	// 			Description: "Software Development Services",
	// 			Quantity:    1,
	// 			UnitPrice:   1000.00,
	// 		},
	// 		{
	// 			Description: "Maintenance Fee",
	// 			Quantity:    2,
	// 			UnitPrice:   250.00,
	// 		},
	// 	},
	// }

	// createResponse, err := client.CreateInvoice(createInvoiceRequest)
	// if err != nil {
	// 	log.Fatalf("Error creating invoice: %v", err)
	// }
	// fmt.Printf("Invoice Creation Response: %s (Code: %d)\n", createResponse.Message, createResponse.Code)

	// // Example 4: Approve an invoice
	// invoiceID, _ := uuid.FromString("98765432-5678-5678-5678-987654321abc") // Replace with actual invoice ID
	// approveInvoiceRequest := &utils.ApproveInvoiceRequest{
	// 	InvoiceID: invoiceID,
	// 	OnChain:   true, // Set to true if you want blockchain processing
	// }

	// approveResponse, err := client.ApproveInvoice(approveInvoiceRequest)
	// if err != nil {
	// 	log.Fatalf("Error approving invoice: %v", err)
	// }
	// fmt.Printf("Invoice Approval Response: %s (Status: %s)\n", approveResponse.Message, approveResponse.Status)
}
