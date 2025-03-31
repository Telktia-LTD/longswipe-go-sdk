package longswipe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func TestClient(t *testing.T) {
	// Valid UUIDs for all test cases
	testUUID := "f733b2ec-b829-4283-bf24-276014307896"
	currencyUUID := "9a0470d3-f580-45b3-85c1-7f3c145540d6"
	customerUUID := "08408121-ff72-4348-bdfd-41cf1c2f9e0d"
	redeemedUUID := "e733b2ec-b829-4283-bf24-276014307896"
	networkUUID := "b733b2ec-b829-4283-bf24-276014307896"
	currencyUUID1 := "f733b2ec-b829-4283-bf24-276014307896"
	currencyUUID2 := "9a0470d3-f580-45b3-85c1-7f3c145540d6"
	createdAt := time.Now().UTC().Format(time.RFC3339)

	// Setup test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.URL.Path {
		// Customer endpoints
		case "/merchant-integrations-server/fetch-customers":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"message": "Customers retrieved successfully",
				"code": 200,
				"status": "success",
				"data": {
					"total": 1,
					"page": 1,
					"limit": 10,
					"customer": [
						{
							"id": "` + testUUID + `",
							"email": "test@example.com",
							"name": "Test User",
							"created_at": "2023-01-01T00:00:00Z"
						}
					]
				}
			}`))

		case "/merchant-integrations-server/fetch-customer-by-email/johndoe@gmail.com":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"message": "Customer retrieved successfully",
				"code": 200,
				"status": "success",
				"customer": {
					"id": "` + testUUID + `",
					"email": "johndoe@gmail.com",
					"name": "John Doe",
					"created_at": "2023-01-01T00:00:00Z"
				}
			}`))

		case "/merchant-integrations-server/add-new-customer":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success", 
				"message": "Customer added",
				"data": {
					"id": "` + testUUID + `",
					"email": "newuser@example.com",
					"name": "New User"
				}
			}`))

		case "/merchant-integrations-server/update-customer":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success", 
				"message": "Customer updated",
				"data": {
					"id": "` + testUUID + `",
					"email": "updated@example.com",
					"name": "Updated User"
				}
			}`))

		case "/merchant-integrations-server/delete-customer/" + testUUID:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success", 
				"message": "Customer deleted"
			}`))

		// Voucher endpoints
		case "/merchant-integrations/verify-voucher":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success",
				"message": "Voucher verified",
				"code": 200,
				"data": {
					"id": "` + testUUID + `",
					"amount": 100.50,
					"balance": 100.50,
					"generatedCurrency": {
						"id": "` + currencyUUID + `",
						"image": "https://example.com/usd.jpg",
						"name": "US Dollar",
						"symbol": "$",
						"Abbreviation": "USD",
						"currencyType": "fiat",
						"isActive": true,
						"createdAt": "2023-01-01T00:00:00Z"
					},
					"code": "LS3130635050",
					"wasPaidFor": true,
					"isUsed": false,
					"createdAt": "2023-01-01T00:00:00Z",
					"createdForMerchant": true,
					"createdForExistingUser": false,
					"createdForNonExistingUser": true,
					"isLocked": false,
					"onchain": false,
					"onchainProcessing": false,
					"cryptoVoucherDetails": {
						"CodeHash": "hash123",
						"Value": "100.50",
						"Balance": "100.50",
						"Creator": "creator123",
						"IsRedeemed": false,
						"TransactionHash": "txn_123456789"
					},
					"redeemedVouchers": [
						{
							"id": "` + redeemedUUID + `",
							"redeemedUserID": null,
							"redeemerWalletAddress": null,
							"voucherID": "` + testUUID + `",
							"user": {
								"id": "` + customerUUID + `",
								"name": "John Doe",
								"email": "john.doe@example.com"
							},
							"amount": 25.25,
							"isMerchant": false,
							"createdAt": "2023-01-02T00:00:00Z"
						}
					],
					"transactionHash": "txn_verify_123",
					"metaData": "Test verification metadata"
				}
			}`))

		case "/merchant-integrations/fetch-voucher-redemption-charges":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success",
				"message": "Charges retrieved",
				"code": 200,
				"data": {
					"charges": {
						"toAmount": 95.25,
						"processingFee": 2.50,
						"totalGasAndProceesingFeeInFromCurrency": 5.25,
						"totalGasCostAndProcessingFeeInWei": 0.0005,
						"exchangeRate": 1.05,
						"percentageCharge": 0.025,
						"isPercentageCharge": true,
						"toCurrency": {
							"id": "` + currencyUUID + `",
							"image": "https://example.com/usd.jpg",
							"name": "US Dollar",
							"symbol": "$",
							"Abbreviation": "USD",
							"currencyType": "fiat",
							"isActive": true,
							"createdAt": "2023-01-01T00:00:00Z"
						},
						"fromCurrency": {
							"id": "` + testUUID + `",
							"image": "https://example.com/eur.jpg",
							"name": "Euro",
							"symbol": "€",
							"Abbreviation": "EUR",
							"currencyType": "fiat",
							"isActive": true,
							"createdAt": "2023-01-01T00:00:00Z"
						},
						"totalDeductable": 5.25
					},
					"voucher": {
						"id": "` + testUUID + `",
						"amount": 100.50,
						"balance": 95.25,
						"generatedCurrency": {
							"id": "` + currencyUUID + `",
							"image": "https://example.com/usd.jpg",
							"name": "US Dollar",
							"symbol": "$",
							"Abbreviation": "USD",
							"currencyType": "fiat",
							"isActive": true,
							"createdAt": "2023-01-01T00:00:00Z"
						},
						"code": "LS3130635050",
						"wasPaidFor": true,
						"isUsed": false,
						"createdAt": "2023-01-01T00:00:00Z",
						"createdForMerchant": true,
						"createdForExistingUser": false,
						"createdForNonExistingUser": true,
						"isLocked": false,
						"onchain": false,
						"onchainProcessing": false,
						"cryptoVoucherDetails": {
							"CodeHash": "hash123",
							"Value": "100.50",
							"Balance": "95.25",
							"Creator": "creator123",
							"IsRedeemed": false,
							"TransactionHash": "txn_123456789"
						},
						"redeemedVouchers": [
							{
								"id": "` + redeemedUUID + `",
								"redeemedUserID": null,
								"redeemerWalletAddress": null,
								"voucherID": "` + testUUID + `",
								"user": {
									"id": "` + customerUUID + `",
									"name": "John Doe",
									"email": "john.doe@example.com"
								},
								"amount": 5.25,
								"isMerchant": false,
								"createdAt": "2023-01-02T00:00:00Z"
							}
						],
						"transactionHash": "txn_123456789",
						"metaData": "Test metadata"
					}
				}
			}`))

		case "/merchant-integrations-server/fetch-customer-by-email/nonexistent@example.com":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{
				"message": "Customer not found",
				"code": 404,
				"status": "error",
				"data": null
			}`))
		case "/merchant-integrations/redeem-voucher":
			// Verify request body
			var req RedeemRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"status": "error", "message": "Invalid request"}`))
				return
			}

			if req.VoucherCode == "INVALID_CODE" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{
					"status": "error",
					"message": "Invalid voucher code",
					"code": 400
				}`))
				return
			}

			if req.Amount <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{
					"status": "error",
					"message": "Amount must be positive",
					"code": 400
				}`))
				return
			}

			// Successful response
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success",
				"message": "Voucher redeemed successfully",
				"code": 200,
				"data": {
					"id": "` + redeemedUUID + `",
					"voucherId": "` + testUUID + `",
					"amount": ` + fmt.Sprintf("%.2f", req.Amount) + `,
					"currency": {
						"id": "` + currencyUUID + `",
						"name": "US Dollar",
						"symbol": "$",
						"abbreviation": "USD"
					},
					"transactionHash": "txn_redeem_123456",
					"timestamp": "2023-01-01T00:00:00Z"
				}
			}`))
		case "/merchant-integrations-server/generate-voucher-for-customer":
			var req GenerateVoucherForCustomerRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"status":"error","code":400,"message":"Invalid request"}`))
				return
			}

			// Validate request
			if req.CurrencyId.String() == "00000000-0000-0000-0000-000000000000" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"status":"error","code":400,"message":"Currency ID is required"}`))
				return
			}

			if req.AmountToPurchase <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(`{"status":"error","code":400,"message":"Amount must be positive"}`))
				return
			}

			// Successful response - matches SuccessResponse struct
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
                "status": "success",
                "message": "Voucher generated successfully",
                "code": 200
            }`))

		case "/merchant-integrations/fetch-supported-cryptonetworks":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"status": "success",
				"message": "Networks retrieved successfully",
				"data": [
					{
						"id": "f733b2ec-b829-4283-bf24-276014307896",
						"networkName": "Ethereum Mainnet",
						"cryptoCurrencies": [
							{
								"currencyDetails": {
									"id": "9a0470d3-f580-45b3-85c1-7f3c145540d6",
									"symbol": "ETH"
								},
								"decimal": 18
							}
						]
					}
				]
			}`))

		case "/merchant-integrations/fetch-supported-currencies":
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
                "message": "Currencies retrieved successfully",
                "code": 200,
                "status": "success",
                "data": {
                    "currencies": [
                        {
                            "id": "` + currencyUUID1 + `",
                            "image": "https://example.com/btc.png",
                            "name": "Bitcoin",
                            "symbol": "BTC",
                            "Abbreviation": "BTC",
                            "currencyType": "CRYPTO",
                            "isActive": true,
                            "createdAt": "` + createdAt + `"
                        },
                        {
                            "id": "` + currencyUUID2 + `",
                            "image": "https://example.com/eth.png",
                            "name": "Ethereum",
                            "symbol": "ETH",
                            "Abbreviation": "ETH",
                            "currencyType": "CRYPTO",
                            "isActive": true,
                            "createdAt": "` + createdAt + `"
                        }
                    ]
                }
            }`))

		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{
				"message": "Endpoint not found",
				"code": 404,
				"status": "error",
				"data": null
			}`))
		}
	}))
	defer ts.Close()

	// Create test client
	client := NewClient(ClientConfig{
		BaseURL:    ts.URL,
		PublicKey:  "test_pk",
		PrivateKey: "test_sk",
		Timeout:    5 * time.Second,
	})

	// Customer tests
	t.Run("GetCustomers", func(t *testing.T) {
		res, err := client.GetCustomers(&Pagination{Page: 1, Limit: 20, Search: ""})
		if err != nil {
			t.Fatalf("GetCustomers failed: %v", err)
		}

		if res.Status != "success" {
			t.Errorf("Expected status 'success', got '%s'", res.Status)
		}

		if len(res.Data.Customers) < 1 {
			t.Fatalf("Expected at least 1 customer, got %d", len(res.Data.Customers))
		}
	})

	t.Run("GetCustomer", func(t *testing.T) {
		res, err := client.GetCustomer("johndoe@gmail.com")
		if err != nil {
			t.Fatalf("GetCustomer failed: %v", err)
		}

		if res.Status != "success" {
			t.Errorf("Expected status 'success', got '%s'", res.Status)
		}

		if res.Data.Email == "" {
			t.Fatal("Expected email to be populated, got empty string")
		}
	})

	t.Run("AddCustomer", func(t *testing.T) {
		res, err := client.AddCustomer(&AddNewCustomer{
			Email: "newuser@example.com",
			Name:  "New User",
		})

		if err != nil {
			t.Fatalf("AddCustomer failed: %v", err)
		}

		if res.Status != "success" {
			t.Errorf("Expected success status, got %s", res.Status)
		}
	})

	t.Run("UpdateCustomer", func(t *testing.T) {
		res, err := client.UpdateCustomer(&UpdatCustomer{
			ID:    testUUID,
			Email: "updated@example.com",
			Name:  "Updated User",
		})

		if err != nil {
			t.Fatalf("UpdateCustomer failed: %v", err)
		}

		if res.Status != "success" {
			t.Errorf("Expected success status, got %s", res.Status)
		}
	})

	t.Run("DeleteCustomer", func(t *testing.T) {
		res, err := client.DeleteCustomer(testUUID)

		if err != nil {
			t.Fatalf("DeleteCustomer failed: %v", err)
		}

		if res.Status != "success" {
			t.Errorf("Expected success status, got %s", res.Status)
		}
	})

	t.Run("BuildCustomerEndpoint", func(t *testing.T) {
		endpoint := buildCustomerEndpoint(2, 50, "test")
		expected := "/merchant-integrations-server/fetch-customers?limit=50&page=2&search=test"

		if endpoint != expected {
			t.Errorf("Expected endpoint %s, got %s", expected, endpoint)
		}
	})

	// Voucher tests
	t.Run("VerifyVoucher", func(t *testing.T) {
		res, err := client.VerifyVoucher(&VerifyVoucherCodeRequest{
			VoucherCode: "LS3130635050",
		})

		if err != nil {
			t.Fatalf("VerifyVoucher failed: %v", err)
		}

		if res.Status != "success" {
			t.Errorf("Expected status 'success', got '%s'", res.Status)
		}

		if res.Data.Code != "LS3130635050" {
			t.Errorf("Expected voucher code 'LS3130635050', got '%s'", res.Data.Code)
		}
	})

	t.Run("GetVoucherRedeemptionCharges", func(t *testing.T) {
		res, err := client.GetVoucherRedeemptionCharges(&RedeemRequest{
			VoucherCode:            "LS3130635050",
			Amount:                 100.50,
			ToCurrencyAbbreviation: "USD",
		})

		if err != nil {
			t.Fatalf("GetVoucherRedeemptionCharges failed: %v", err)
		}

		// Validate charges
		if res.Data.Charges.ProcessingFee != 2.50 {
			t.Errorf("Expected Processing fee 2.50, got %f", res.Data.Charges.ProcessingFee)
		}
		if res.Data.Charges.TotalDeductable != 5.25 {
			t.Errorf("Expected totalDeductable 5.25, got %f", res.Data.Charges.TotalDeductable)
		}

		// Validate voucher
		if res.Data.Voucher.Code != "LS3130635050" {
			t.Errorf("Expected voucher code 'LS3130635050', got '%s'", res.Data.Voucher.Code)
		}
		if res.Data.Voucher.Balance != 95.25 {
			t.Errorf("Expected voucher balance 95.25, got %f", res.Data.Voucher.Balance)
		}
	})

	// Error cases
	t.Run("ErrorCases", func(t *testing.T) {
		// Test with invalid URL to force error
		badClient := NewClient(ClientConfig{
			BaseURL:    "http://invalid-url",
			PublicKey:  "test_pk",
			PrivateKey: "test_sk",
			Timeout:    1 * time.Microsecond,
		})

		_, err := badClient.GetCustomers(&Pagination{Page: 1, Limit: 10})
		if err == nil {
			t.Error("Expected error for GetCustomers with invalid client, got nil")
		}

		// Test not found case
		_, err = client.GetCustomer("nonexistent@example.com")
		if err == nil {
			t.Error("Expected error for nonexistent customer, got nil")
		}

		if !strings.Contains(err.Error(), "Customer not found") {
			t.Errorf("Expected error to contain 'Customer not found', got: %v", err)
		}
	})

	t.Run("RedeemVoucher", func(t *testing.T) {
		t.Run("SuccessfulRedemption", func(t *testing.T) {
			req := &RedeemRequest{
				VoucherCode:            "VALID123",
				Amount:                 100.50,
				ToCurrencyAbbreviation: "USD",
				ReferenceId:            "order_123",
			}

			res, err := client.RedeemVoucher(req)
			if err != nil {
				t.Fatalf("RedeemVoucher failed: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}
		})

		t.Run("InvalidVoucherCode", func(t *testing.T) {
			req := &RedeemRequest{
				VoucherCode:            "INVALID_CODE",
				Amount:                 100.50,
				ToCurrencyAbbreviation: "USD",
			}

			_, err := client.RedeemVoucher(req)
			if err == nil {
				t.Fatal("Expected error for invalid voucher code, got nil")
			}
		})

		t.Run("InvalidAmount", func(t *testing.T) {
			req := &RedeemRequest{
				VoucherCode:            "VALID123",
				Amount:                 0,
				ToCurrencyAbbreviation: "USD",
			}

			_, err := client.RedeemVoucher(req)
			if err == nil {
				t.Fatal("Expected error for invalid amount, got nil")
			}
		})
	})

	t.Run("GenerateVoucher", func(t *testing.T) {
		currencyID, _ := uuid.FromString(currencyUUID)
		customerID, _ := uuid.FromString(customerUUID)
		networkID, _ := uuid.FromString(networkUUID)

		t.Run("Success", func(t *testing.T) {
			req := &GenerateVoucherForCustomerRequest{
				CurrencyId:          currencyID,
				AmountToPurchase:    100.50,
				CustomerID:          customerID,
				OnChain:             true,
				BlockchainNetworkId: networkID,
			}

			res, err := client.GenerateVoucher(req)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}
			if res.Code != 200 {
				t.Errorf("Expected code 200, got %d", res.Code)
			}
			if res.Message != "Voucher generated successfully" {
				t.Errorf("Unexpected message: %s", res.Message)
			}
		})

		t.Run("InvalidCurrency", func(t *testing.T) {
			req := &GenerateVoucherForCustomerRequest{
				CurrencyId:       uuid.Nil, // Invalid
				AmountToPurchase: 100.50,
				CustomerID:       customerID,
			}

			res, err := client.GenerateVoucher(req)
			if err == nil {
				t.Fatal("Expected error but got nil")
			}
			if res != nil {
				t.Error("Expected nil response on error")
			}
		})

		t.Run("ZeroAmount", func(t *testing.T) {
			req := &GenerateVoucherForCustomerRequest{
				CurrencyId:       currencyID,
				AmountToPurchase: 0, // Invalid
				CustomerID:       customerID,
			}

			res, err := client.GenerateVoucher(req)
			if err == nil {
				t.Fatal("Expected error but got nil")
			}
			if res != nil {
				t.Error("Expected nil response on error")
			}
		})
	})

	t.Run("Success", func(t *testing.T) {
		// Test GetAllNetwork success case separately
		t.Run("GetAllNetwork", func(t *testing.T) {
			resp, err := client.GetAllNetwork()
			if err != nil {
				t.Fatalf("GetAllNetwork failed: %v", err)
			}

			if resp.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", resp.Status)
			}
		})
	})

	t.Run("Success", func(t *testing.T) {
		resp, err := client.GetAllCurrency()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify response metadata
		if resp.Status != "success" {
			t.Errorf("Expected status 'success', got '%s'", resp.Status)
		}
		if resp.Code != 200 {
			t.Errorf("Expected code 200, got %d", resp.Code)
		}

		// Verify currencies data
		if len(resp.Data.Currencies) != 2 {
			t.Fatalf("Expected 2 currencies, got %d", len(resp.Data.Currencies))
		}

		// Verify first currency
		btc := resp.Data.Currencies[0]
		if btc.ID.String() != currencyUUID1 {
			t.Errorf("Expected currency ID %s, got %s", currencyUUID1, btc.ID)
		}
		if btc.Symbol != "BTC" {
			t.Errorf("Expected symbol 'BTC', got '%s'", btc.Symbol)
		}
		if !btc.IsActive {
			t.Error("Expected currency to be active")
		}
	})

	t.Run("EmptyResponse", func(t *testing.T) {
		// Create test server with empty response
		emptyTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{
                "message": "No currencies found",
                "code": 200,
                "status": "success",
                "data": {
                    "currencies": []
                }
            }`))
		}))
		defer emptyTS.Close()

		emptyClient := NewClient(ClientConfig{
			BaseURL:    emptyTS.URL,
			PublicKey:  "test_pk",
			PrivateKey: "test_sk",
			Timeout:    5 * time.Second,
		})

		resp, err := emptyClient.GetAllCurrency()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(resp.Data.Currencies) != 0 {
			t.Errorf("Expected empty currencies array, got %d items", len(resp.Data.Currencies))
		}
	})

	t.Run("ErrorResponse", func(t *testing.T) {
		// Create test server with error response
		errorTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{
                "message": "Internal server error",
                "code": 500,
                "status": "error"
            }`))
		}))
		defer errorTS.Close()

		errorClient := NewClient(ClientConfig{
			BaseURL:    errorTS.URL,
			PublicKey:  "test_pk",
			PrivateKey: "test_sk",
			Timeout:    5 * time.Second,
		})

		_, err := errorClient.GetAllCurrency()
		if err == nil {
			t.Fatal("Expected error but got nil")
		}
	})

}
