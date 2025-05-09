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
		PublicKey:  "pk_live_KIJneg1IWF0r6S6VeY8FfRuZckq2jmiBcOvgYA-hYb8=",
		Timeout:    100 * time.Second,
		PrivateKey: "sk_live_vxCrmEqSorXlcG99U9sFcRIPezmj0eKQpg_I-EBXZ50=",
	})

	currency, err := client.GetAllCurrency()
	if err != nil {
		log.Fatalf("Error fetching currency: %v", err)
	}
	fmt.Println("Fetched currency:", currency)

	users, err := client.GetAllUser()
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)

	}
	fmt.Println("Fetched currency:", users)

	// CREATE INVOICE
	currencyID, _ := uuid.FromString("9a0970d2-6680-45b3-85c1-7f3c189807e7")

	createInvoiceRequest := &longswipe.CreateInvoiceRequest{
		FullName:     "John Doe",
		Email:        "john.doe@example.com",
		MerchantCode: "LS84816",
		InvoiceDate:  time.Now(),
		DueDate:      time.Now().AddDate(0, 0, 30), // Due in 30 days
		CurrencyId:   currencyID,
		InvoiceItems: []longswipe.InvoiceItemRequest{
			{
				Description: "Software Development",
				Quantity:    1,
				UnitPrice:   1000.00,
			},
			{
				Description: "Maintenance Services",
				Quantity:    2,
				UnitPrice:   250.00,
			},
		},
	}

	createResponse, err := client.CreateInvoice(createInvoiceRequest)
	if err != nil {
		log.Fatalf("Error creating invoice: %v", err.Error())
	}
	fmt.Printf("Invoice Creation Response: %s (Code: %d)\n", createResponse.Message, createResponse.Code)

	// FETCH ALL INVOICE
	paginationParams := &longswipe.Pagination{
		Page:   1,
		Limit:  10,
		Search: "",
	}

	invoices, err := client.FetchInvoice(paginationParams)
	if err != nil {
		log.Fatalf("Error fetching invoices: %v", err)
	}
	fmt.Println("Fetched invoice:", invoices.Data.Invoices[0].ID)

	// Example 2: Get all allowed invoice currencies
	currencies, err := client.GetAllInvoiceCurrency()
	if err != nil {
		log.Fatalf("Error fetching invoice currencies: %v", err)
	}
	fmt.Println("Fetched currencies:", currencies.Data[0].Currency.Name)

	// // Example 4: Approve an invoice
	invoiceID, _ := uuid.FromString("edcf9ed6-9341-4f95-9ba6-739351acdc5c") // Replace with actual invoice ID
	approveInvoiceRequest := &longswipe.ApproveInvoiceRequest{
		InvoiceID: invoiceID,
	}
	approveResponse, err := client.ApproveInvoice(approveInvoiceRequest)
	if err != nil {
		log.Fatalf("Error approving invoice: %v", err)
	}
	fmt.Printf("Invoice Approval Response: %s (Status: %s)\n", approveResponse.Message, approveResponse.Status)
}
