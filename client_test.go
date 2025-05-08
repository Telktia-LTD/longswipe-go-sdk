package longswipe

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

// TestData contains all test data structures
type TestData struct {
	TestUUID       uuid.UUID
	CurrencyUUID   uuid.UUID
	CustomerUUID   uuid.UUID
	RedeemedUUID   uuid.UUID
	NetworkUUID    uuid.UUID
	CreatedAt      time.Time
	Customer       CustomerData
	Voucher        VoucherResponse
	Charges        V2PayoutDetailsResponse
	Currencies     []Currencies
	Networks       []CryptoNetworkDetails
	SuccessResp    SuccessResponse
	CustomerResp   CustomerResponse
	CustomerList   CustomersResponse
	VoucherResp    VerifyVoucherResponse
	RedemptionResp ApiResponse[struct {
		ID              uuid.UUID       `json:"id"`
		VoucherId       uuid.UUID       `json:"voucherId"`
		Amount          float64         `json:"amount"`
		Currency        CurrencyDetails `json:"currency"`
		TransactionHash string          `json:"transactionHash"`
		Timestamp       time.Time       `json:"timestamp"`
	}]
	ChargesResp         RedeemeVoucherDetailsResponse
	NetworkResp         CryptoNetworkResponse
	CurrencyResp        FetchCurrenciesResponse
	InvoiceResp         MerchantInvoiceResponse
	InvoiceCurrencyResp FetchAllAllowedInvoiceCurrencyResponse
}

func setupTestData() TestData {
	now := time.Now()
	testUUID := uuid.Must(uuid.FromString("f733b2ec-b829-4283-bf24-276014307896"))
	currencyUUID := uuid.Must(uuid.FromString("9a0470d3-f580-45b3-85c1-7f3c145540d6"))
	customerUUID := uuid.Must(uuid.FromString("08408121-ff72-4348-bdfd-41cf1c2f9e0d"))
	redeemedUUID := uuid.Must(uuid.FromString("e733b2ec-b829-4283-bf24-276014307896"))
	networkUUID := uuid.Must(uuid.FromString("b733b2ec-b829-4283-bf24-276014307896"))

	customer := CustomerData{
		ID:    testUUID,
		Email: "test@example.com",
		Name:  "Test User",
	}

	voucher := VoucherResponse{
		ID:      testUUID,
		Amount:  100.50,
		Balance: 100.50,
		Code:    "LS3130635050",
		GeneratedCurrency: CurrencyDetails{
			ID:           currencyUUID,
			Image:        "https://example.com/usd.jpg",
			Name:         "US Dollar",
			Symbol:       "$",
			Abbreviation: "USD",
			CurrencyType: "fiat",
			IsActive:     true,
		},
		WasPaidFor: true,
		IsUsed:     false,
		CreatedAt:  now,
	}

	charges := V2PayoutDetailsResponse{
		ToAmount:                               95.25,
		ProcessingFee:                          2.50,
		TotalGasAndProceesingFeeInFromCurrency: 5.25,
		TotalGasCostAndProcessingFeeInWei:      0.0005,
		ExchangeRate:                           1.05,
		PercentageCharge:                       0.025,
		IsPercentageCharge:                     true,
		ToCurrency: CurrencyDetails{
			ID:           currencyUUID,
			Name:         "US Dollar",
			Symbol:       "$",
			Abbreviation: "USD",
		},
		FromCurrency: CurrencyDetails{
			ID:           testUUID,
			Name:         "Euro",
			Symbol:       "€",
			Abbreviation: "EUR",
		},
		TotalDeductable: 5.25,
	}

	currencies := []Currencies{
		{
			ID:           testUUID,
			Image:        "https://example.com/btc.png",
			Currency:     "Bitcoin",
			Symbol:       "BTC",
			Abbreviation: "BTC",
			CurrencyType: "CRYPTO",
			IsActive:     true,
			CreatedAt:    now,
		},
		{
			ID:           currencyUUID,
			Image:        "https://example.com/eth.png",
			Currency:     "Ethereum",
			Symbol:       "ETH",
			Abbreviation: "ETH",
			CurrencyType: "CRYPTO",
			IsActive:     true,
			CreatedAt:    now,
		},
	}

	networks := []CryptoNetworkDetails{
		{
			ID:          testUUID,
			NetworkName: "Ethereum Mainnet",
			CryptoCurrencies: []CryptoCurrency{
				{
					CurrencyData: CurrencyDetails{
						ID:     currencyUUID,
						Symbol: "ETH",
					},
				},
			},
		},
	}

	invoiceCurrency := FetchAllAllowedInvoiceCurrencyResponse{
		Status:  "success",
		Message: "Invoice currencies retrieved",
		Code:    200,
		Data: []AllowedInvoiceCurrency{
			{
				Currency: CurrencyDetails{
					ID:           currencyUUID,
					Name:         "US Dollar",
					Symbol:       "$",
					Abbreviation: "USD",
				},
				IsEnabled: true,
			},
		},
	}

	invoiceResp := MerchantInvoiceResponse{
		Status:  "success",
		Message: "Invoices retrieved",
		Code:    200,
		Data: struct {
			Invoices []Invoice `json:"invoices"`
			Total    int       `json:"total"`
		}{
			Total: 1,
			Invoices: []Invoice{
				{
					ID:                testUUID,
					InvoiceNumber:     "INV-001",
					MerchantUser:      MerchantUserDatas{},
					UserId:            &customerUUID,
					InvoiceDate:       now,
					DueDate:           now.AddDate(0, 0, 30),
					TotalAmount:       1500.00,
					Status:            "pending",
					InvoiceItems:      []InvoiceItem{},
					Currency:          CurrencyDetails{},
					BlockchainNetwork: nil,
					CreatedAt:         now,
					UpdatedAt:         now,
				},
			},
		},
	}

	return TestData{
		TestUUID:     testUUID,
		CurrencyUUID: currencyUUID,
		CustomerUUID: customerUUID,
		RedeemedUUID: redeemedUUID,
		NetworkUUID:  networkUUID,
		CreatedAt:    now,
		Customer:     customer,
		Voucher:      voucher,
		Charges:      charges,
		Currencies:   currencies,
		Networks:     networks,
		SuccessResp: SuccessResponse{
			Status:  "success",
			Message: "Operation completed",
			Code:    200,
		},
		CustomerResp: CustomerResponse{
			Message: "Customer retrieved",
			Code:    200,
			Status:  "success",
			Data:    customer,
		},
		CustomerList: CustomersResponse{
			Message: "Customers retrieved",
			Code:    200,
			Status:  "success",
			Data: CustomerDetails{
				Total:     1,
				Page:      1,
				Limit:     10,
				Customers: []CustomerData{customer},
			},
		},
		VoucherResp: VerifyVoucherResponse{
			Status:  "success",
			Message: "Voucher verified",
			Code:    200,
			Data:    voucher,
		},
		RedemptionResp: ApiResponse[struct {
			ID              uuid.UUID       `json:"id"`
			VoucherId       uuid.UUID       `json:"voucherId"`
			Amount          float64         `json:"amount"`
			Currency        CurrencyDetails `json:"currency"`
			TransactionHash string          `json:"transactionHash"`
			Timestamp       time.Time       `json:"timestamp"`
		}]{
			Status:  "success",
			Message: "Voucher redeemed",
			Code:    200,
			Data: struct {
				ID              uuid.UUID       `json:"id"`
				VoucherId       uuid.UUID       `json:"voucherId"`
				Amount          float64         `json:"amount"`
				Currency        CurrencyDetails `json:"currency"`
				TransactionHash string          `json:"transactionHash"`
				Timestamp       time.Time       `json:"timestamp"`
			}{
				ID:              redeemedUUID,
				VoucherId:       testUUID,
				Amount:          100.50,
				Currency:        voucher.GeneratedCurrency,
				TransactionHash: "txn_redeem_123456",
				Timestamp:       now,
			},
		},
		ChargesResp: RedeemeVoucherDetailsResponse{
			Status:  "success",
			Message: "Charges retrieved",
			Code:    200,
			Data: RedeemVoucherDetailDataAll{
				Charges: charges,
				Voucher: voucher,
			},
		},
		NetworkResp: CryptoNetworkResponse{
			Status:  "success",
			Message: "Networks retrieved",
			Code:    200,
			Data:    networks,
		},
		CurrencyResp: FetchCurrenciesResponse{
			Message: "Currencies retrieved",
			Code:    200,
			Status:  "success",
			Data: struct {
				Currencies []Currencies `json:"currencies"`
			}{
				Currencies: currencies,
			},
		},
		InvoiceResp:         invoiceResp,
		InvoiceCurrencyResp: invoiceCurrency,
	}

}

func setupTestServer(td TestData) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.URL.Path {
		case "/merchant-integrations-server/fetch-customers":
			json.NewEncoder(w).Encode(td.CustomerList)
		case "/merchant-integrations-server/fetch-customer-by-email/johndoe@gmail.com":
			json.NewEncoder(w).Encode(td.CustomerResp)
		case "/merchant-integrations-server/add-new-customer":
			json.NewEncoder(w).Encode(td.SuccessResp)
		case "/merchant-integrations-server/update-customer":
			json.NewEncoder(w).Encode(td.SuccessResp)
		case "/merchant-integrations-server/delete-customer/" + td.TestUUID.String():
			json.NewEncoder(w).Encode(td.SuccessResp)
		case "/merchant-integrations/verify-voucher":
			json.NewEncoder(w).Encode(td.VoucherResp)
		case "/merchant-integrations/fetch-voucher-redemption-charges":
			json.NewEncoder(w).Encode(td.ChargesResp)
		case "/merchant-integrations-server/fetch-customer-by-email/nonexistent@example.com":
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{
				Status:  "error",
				Message: "Customer not found",
				Code:    404,
			})
		case "/merchant-integrations/redeem-voucher":
			var req RedeemRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{
					Status:  "error",
					Message: "Invalid request",
					Code:    400,
				})
				return
			}

			if req.VoucherCode == "INVALID_CODE" {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{
					Status:  "error",
					Message: "Invalid voucher code",
					Code:    400,
				})
				return
			}

			if req.Amount <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{
					Status:  "error",
					Message: "Amount must be positive",
					Code:    400,
				})
				return
			}

			json.NewEncoder(w).Encode(td.RedemptionResp)
		case "/merchant-integrations-server/generate-voucher-for-customer":
			var req GenerateVoucherForCustomerRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{
					Status:  "error",
					Message: "Invalid request",
					Code:    400,
				})
				return
			}

			if req.CurrencyId == uuid.Nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{
					Status:  "error",
					Message: "Currency ID is required",
					Code:    400,
				})
				return
			}

			if req.AmountToPurchase <= 0 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(ErrorResponse{
					Status:  "error",
					Message: "Amount must be positive",
					Code:    400,
				})
				return
			}

			json.NewEncoder(w).Encode(td.SuccessResp)
		case "/merchant-integrations/fetch-supported-cryptonetworks":
			json.NewEncoder(w).Encode(td.NetworkResp)
		case "/merchant-integrations/fetch-supported-currencies":
			json.NewEncoder(w).Encode(td.CurrencyResp)
		case "/merchant-integrations-server/fetch-invoice":
			json.NewEncoder(w).Encode(td.InvoiceResp)
		case "/merchant-integrations-server/fetch-all-allowed-invoice-Currency":
			json.NewEncoder(w).Encode(td.InvoiceCurrencyResp)
		case "/merchant-integrations-server/create-invoice":
			json.NewEncoder(w).Encode(td.SuccessResp)
		case "/merchant-integrations-server/approve-invoice":
			json.NewEncoder(w).Encode(td.SuccessResp)
		default:
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{
				Status:  "error",
				Message: "Endpoint not found",
				Code:    404,
			})
		}
	}))
}

func TestClient(t *testing.T) {
	td := setupTestData()
	ts := setupTestServer(td)
	defer ts.Close()

	client := NewClient(ClientConfig{
		BaseURL:    ts.URL,
		PublicKey:  "test_pk",
		PrivateKey: "test_sk",
		Timeout:    5 * time.Second,
	})

	t.Run("CustomerOperations", func(t *testing.T) {
		t.Run("GetCustomers", func(t *testing.T) {
			res, err := client.GetCustomers(&Pagination{Page: 1, Limit: 20})
			if err != nil {
				t.Fatalf("GetCustomers failed: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}

			if len(res.Data.Customers) != 1 {
				t.Errorf("Expected 1 customer, got %d", len(res.Data.Customers))
			}
		})

		t.Run("GetCustomer", func(t *testing.T) {
			res, err := client.GetCustomer("johndoe@gmail.com")
			if err != nil {
				t.Fatalf("GetCustomer failed: %v", err)
			}

			if res.Data.Email != td.Customer.Email {
				t.Errorf("Expected email %s, got %s", td.Customer.Email, res.Data.Email)
			}
		})

		t.Run("AddCustomer", func(t *testing.T) {
			res, err := client.AddCustomer(&AddNewCustomer{
				Email: "new@example.com",
				Name:  "New User",
			})
			if err != nil {
				t.Fatalf("AddCustomer failed: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}
		})

		t.Run("UpdateCustomer", func(t *testing.T) {
			res, err := client.UpdateCustomer(&UpdatCustomer{
				ID:    td.TestUUID.String(),
				Email: "updated@example.com",
				Name:  "Updated User",
			})
			if err != nil {
				t.Fatalf("UpdateCustomer failed: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}
		})

		t.Run("DeleteCustomer", func(t *testing.T) {
			res, err := client.DeleteCustomer(td.TestUUID.String())
			if err != nil {
				t.Fatalf("DeleteCustomer failed: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}
		})
	})

	t.Run("VoucherOperations", func(t *testing.T) {
		t.Run("VerifyVoucher", func(t *testing.T) {
			_, err := client.VerifyVoucher(&VerifyVoucherCodeRequest{
				VoucherCode: td.Voucher.Code,
			})
			if err != nil {
				t.Fatalf("VerifyVoucher failed: %v", err)
			}

			// if res.Data.Code != td.Voucher.Code {
			// 	t.Errorf("Expected voucher code %s, got %s", td.Voucher.Code, res.Data.Code)
			// }
		})

		t.Run("GetVoucherRedemptionCharges", func(t *testing.T) {
			res, err := client.GetVoucherRedeemptionCharges(&RedeemRequest{
				VoucherCode:            td.Voucher.Code,
				Amount:                 100.50,
				ToCurrencyAbbreviation: "USD",
			})
			if err != nil {
				t.Fatalf("GetVoucherRedemptionCharges failed: %v", err)
			}

			if res.Data.Charges.ProcessingFee != td.Charges.ProcessingFee {
				t.Errorf("Expected processing fee %f, got %f",
					td.Charges.ProcessingFee, res.Data.Charges.ProcessingFee)
			}
		})

		t.Run("RedeemVoucher", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				res, err := client.RedeemVoucher(&RedeemRequest{
					VoucherCode:            td.Voucher.Code,
					Amount:                 50.0,
					ToCurrencyAbbreviation: "USD",
				})
				if err != nil {
					t.Fatalf("RedeemVoucher failed: %v", err)
				}

				t.Logf("Voucher redemption success: %s", res.Message)
			})

			t.Run("InvalidCode", func(t *testing.T) {
				_, err := client.RedeemVoucher(&RedeemRequest{
					VoucherCode:            "INVALID_CODE",
					Amount:                 50.0,
					ToCurrencyAbbreviation: "USD",
				})
				if err == nil {
					t.Error("Expected error for invalid voucher code, got nil")
				}
			})

			t.Run("InvalidAmount", func(t *testing.T) {
				_, err := client.RedeemVoucher(&RedeemRequest{
					VoucherCode:            td.Voucher.Code,
					Amount:                 0,
					ToCurrencyAbbreviation: "USD",
				})
				if err == nil {
					t.Error("Expected error for invalid amount, got nil")
				}
			})
		})

		t.Run("GenerateVoucher", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				res, err := client.GenerateVoucher(&GenerateVoucherForCustomerRequest{
					CurrencyId:          td.CurrencyUUID,
					AmountToPurchase:    100.0,
					CustomerID:          td.CustomerUUID,
					OnChain:             true,
					BlockchainNetworkId: td.NetworkUUID,
				})
				if err != nil {
					t.Fatalf("GenerateVoucher failed: %v", err)
				}

				if res.Status != "success" {
					t.Errorf("Expected status 'success', got '%s'", res.Status)
				}
			})

			t.Run("InvalidCurrency", func(t *testing.T) {
				_, err := client.GenerateVoucher(&GenerateVoucherForCustomerRequest{
					CurrencyId:       uuid.Nil,
					AmountToPurchase: 100.0,
					CustomerID:       td.CustomerUUID,
				})
				if err == nil {
					t.Error("Expected error for invalid currency, got nil")
				}
			})

			t.Run("ZeroAmount", func(t *testing.T) {
				_, err := client.GenerateVoucher(&GenerateVoucherForCustomerRequest{
					CurrencyId:       td.CurrencyUUID,
					AmountToPurchase: 0,
					CustomerID:       td.CustomerUUID,
				})
				if err == nil {
					t.Error("Expected error for zero amount, got nil")
				}
			})
		})
	})

	t.Run("NetworkOperations", func(t *testing.T) {
		t.Run("GetAllNetwork", func(t *testing.T) {
			res, err := client.GetAllNetwork()
			if err != nil {
				t.Fatalf("GetAllNetwork failed: %v", err)
			}

			if len(res.Data) != len(td.Networks) {
				t.Errorf("Expected %d networks, got %d", len(td.Networks), len(res.Data))
			}
		})
	})

	t.Run("CurrencyOperations", func(t *testing.T) {
		t.Run("GetAllCurrency", func(t *testing.T) {
			res, err := client.GetAllCurrency()
			if err != nil {
				t.Fatalf("GetAllCurrency failed: %v", err)
			}

			if len(res.Data.Currencies) != len(td.Currencies) {
				t.Errorf("Expected %d currencies, got %d", len(td.Currencies), len(res.Data.Currencies))
			}
		})
	})

	t.Run("ErrorCases", func(t *testing.T) {
		t.Run("InvalidURL", func(t *testing.T) {
			badClient := NewClient(ClientConfig{
				BaseURL:    "http://invalid-url",
				PublicKey:  "test_pk",
				PrivateKey: "test_sk",
				Timeout:    1 * time.Microsecond,
			})

			_, err := badClient.GetCustomers(&Pagination{Page: 1, Limit: 10})
			if err == nil {
				t.Error("Expected error for invalid URL, got nil")
			}
		})

		t.Run("CustomerNotFound", func(t *testing.T) {
			_, err := client.GetCustomer("nonexistent@example.com")
			if err == nil {
				t.Error("Expected error for non-existent customer, got nil")
			}
		})
	})

	t.Run("InvoiceOperations", func(t *testing.T) {
		t.Run("FetchInvoice", func(t *testing.T) {
			res, err := client.FetchInvoice(&Pagination{
				Page:   1,
				Limit:  10,
				Search: "INV-",
			})
			if err != nil {
				t.Fatalf("FetchInvoice failed: %v", err)
			}

			if res.Data.Total != 1 {
				t.Errorf("Expected 1 invoice, got %d", res.Data.Total)
			}
		})

		t.Run("GetAllInvoiceCurrency", func(t *testing.T) {
			res, err := client.GetAllInvoiceCurrency()
			if err != nil {
				t.Fatalf("GetAllInvoiceCurrency failed: %v", err)
			}

			if len(res.Data) == 0 {
				t.Error("Expected at least one currency, got none")
			}
		})

		t.Run("CreateInvoice", func(t *testing.T) {
			req := &CreateInvoiceRequest{
				MerchantUserId: td.TestUUID,
				InvoiceDate:    time.Now(),
				DueDate:        time.Now().AddDate(0, 0, 30),
				CurrencyId:     td.CurrencyUUID,
				InvoiceItems: []InvoiceItemRequest{
					{
						Description: "Test Item",
						Quantity:    1,
						UnitPrice:   100.00,
					},
				},
			}

			res, err := client.CreateInvoice(req)
			if err != nil {
				t.Fatalf("CreateInvoice failed: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}
		})

		t.Run("ApproveInvoice", func(t *testing.T) {
			req := &ApproveInvoiceRequest{
				InvoiceID: td.TestUUID,
				OnChain:   true,
			}

			res, err := client.ApproveInvoice(req)
			if err != nil {
				t.Fatalf("ApproveInvoice failed: %v", err)
			}

			if res.Status != "success" {
				t.Errorf("Expected status 'success', got '%s'", res.Status)
			}
		})
	})
}
