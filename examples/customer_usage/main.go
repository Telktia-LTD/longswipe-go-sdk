package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
	"github.com/gofrs/uuid"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    longswipe.PRODUCTION,
		PublicKey:  "pk_live_KIJneg1IWF0r6S6VeY8FfRuZckq2jmiBcOvgYA-hYb8=",
		Timeout:    10 * time.Second,
		PrivateKey: "sk_live_vxCrmEqSorXlcG99U9sFcRIPezmj0eKQpg_I-EBXZ50=",
	})

	// FETCH CUSTOMERS
	cus, err := client.GetCustomers(&longswipe.Pagination{
		Page:   1,
		Limit:  20,
		Search: "",
	})
	if err != nil {
		log.Fatalf("fetch customers failed: %v", err)
	}
	fmt.Println("Service status:", cus.Data.Customers[1])

	// ADD CUSTOMER
	cusRes, err := client.AddCustomer(&longswipe.AddNewCustomer{
		Name:  "john doe",
		Email: "johndoe@gmail.com",
	})
	if err != nil {
		log.Fatalf("add customer failed: %v", err)
	}
	fmt.Println("Customer added successfully:", cusRes)

	// Update Customer
	customerID, _ := uuid.FromString("c177a874-5642-4f29-ae51-e4356d376863")

	updateRes, err := client.UpdateCustomer(&longswipe.UpdatCustomer{
		ID:    customerID,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
	})
	if err != nil {
		fmt.Println("update customer failed:", err)
	}
	fmt.Println("Customer updated successfully:", updateRes)

	// DELETE CUSTOMER
	deleteRes, err := client.DeleteCustomer(customerID)
	if err != nil {
		fmt.Println("delete customer failed:", err)
	}
	fmt.Println("Customer deleted successfully:", deleteRes)
}
