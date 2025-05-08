package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Telktia-LTD/longswipe-go-sdk"
)

func main() {
	client := longswipe.NewClient(longswipe.ClientConfig{
		BaseURL:    longswipe.PRODUCTION,
		PublicKey:  "pk_live_uQxOuuP-ka03sqIFgBBqbaYL9-vGkso9BOfpXUdVSIY=",
		Timeout:    10 * time.Second,
		PrivateKey: "sk_live_H-CqZ9PivCpiG9KQDyIp_G0kTYm8MH0cMpIg6CKYAVI=",
	})

	// ADD USER
	// res, err := client.AddUser(&longswipe.AddNewUserRequest{Name: "Testing", Email: "festech1426@gmail.com", Role: "USER"})
	// if err != nil {
	// 	log.Fatalf("Add User failed: %v", err)
	// }
	// fmt.Println("Service status:", res)

	// ADD USER
	res, err := client.GetAllUser()
	if err != nil {
		log.Fatalf("fetch users failed: %v", err)
	}
	fmt.Println("Service status:", res)
}
