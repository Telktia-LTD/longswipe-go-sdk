package longswipe

/*
LongSwipe Go SDK Test Suite

This test suite is designed to be easily extensible for testing all SDK functionalities.

## Current Test Coverage:
- Invoice Operations (Create, Get/List, Approve)

## Adding New Test Suites:

To add tests for other functionalities like Vouchers, follow this pattern:

1. Add mock data generators (similar to generateMockInvoiceCreateRequest)
2. Add test suite function (similar to TestInvoiceOperations)
3. Add individual test functions for each operation
4. Add integration workflow tests

Example for Voucher tests:

```go
// Mock data generators
func generateMockRedeemRequest() *RedeemRequest { ... }
func generateMockVoucherResponse() *VerifyVoucherResponse { ... }

// Test suite
func TestVoucherOperations(t *testing.T) {
	t.Run("RedeemVoucher", TestRedeemVoucher)
	t.Run("VerifyVoucher", TestVerifyVoucher)
	t.Run("GenerateVoucher", TestGenerateVoucher)
}

// Individual tests
func TestRedeemVoucher(t *testing.T) { ... }
func TestVerifyVoucher(t *testing.T) { ... }
```

## Test Structure:
- Each operation has Success, Error, and Edge case scenarios
- Mock HTTP server simulates API responses
- Clean separation of test data generation
- Comprehensive assertions for both positive and negative cases
*/

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestServer represents a mock HTTP server for testing
type TestServer struct {
	server   *httptest.Server
	handlers map[string]http.HandlerFunc
}

// NewTestServer creates a new test server with configurable handlers
func NewTestServer() *TestServer {
	ts := &TestServer{
		handlers: make(map[string]http.HandlerFunc),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler, exists := ts.handlers[r.URL.Path]
		if !exists {
			http.NotFound(w, r)
			return
		}
		handler(w, r)
	})

	ts.server = httptest.NewServer(mux)
	return ts
}

func (ts *TestServer) AddHandler(path string, handler http.HandlerFunc) {
	ts.handlers[path] = handler
}

func (ts *TestServer) URL() string {
	return ts.server.URL
}

func (ts *TestServer) Close() {
	ts.server.Close()
}

// Mock data generators
func generateMockUUID() uuid.UUID {
	id, _ := uuid.NewV4()
	return id
}

func generateMockInvoiceCreateRequest() *CreateInvoiceRequest {
	return &CreateInvoiceRequest{
		FullName:     "John Doe",
		Email:        "john.doe@example.com",
		MerchantCode: "MERCH123",
		InvoiceDate:  time.Now(),
		DueDate:      time.Now().AddDate(0, 0, 30),
		InvoiceItems: []InvoiceItemRequest{
			{
				Description: "Product A",
				Quantity:    2,
				UnitPrice:   50.00,
			},
			{
				Description: "Product B",
				Quantity:    1,
				UnitPrice:   100.00,
			},
		},
		CurrencyAbbreviation:          "USD",
		BlockchainNetworkAbbreviation: "ETH",
	}
}

func generateMockSuccessResponse() *SuccessResponse {
	return &SuccessResponse{
		Status:  "success",
		Message: "Operation completed successfully",
		Code:    200,
	}
}

func generateMockInvoiceResponse() *InvoiceResponse {
	invoiceID := generateMockUUID().String()
	userID := generateMockUUID().String()
	currencyID := generateMockUUID().String()
	networkID := generateMockUUID().String()
	merchantUserID := generateMockUUID().String()
	itemID := generateMockUUID().String()

	return &InvoiceResponse{
		Status:  "success",
		Message: "Invoices retrieved successfully",
		Code:    200,
		Data: InvoiceDetails{
			Total: 1,
			Invoices: []Invoice{
				{
					ID:            invoiceID,
					InvoiceNumber: "INV-001",
					UserID:        userID,
					InvoiceDate:   "2023-12-01T10:00:00Z",
					DueDate:       "2023-12-31T23:59:59Z",
					Status:        "pending",
					TotalAmount:   200.00,
					CreatedAt:     "2023-12-01T10:00:00Z",
					UpdatedAt:     "2023-12-01T10:00:00Z",
					Currency: Currency{
						ID:           currencyID,
						Name:         "US Dollar",
						Symbol:       "$",
						Abbrev:       "USD",
						CurrencyType: "fiat",
						IsActive:     true,
						Image:        "https://example.com/usd.png",
					},
					BlockchainNetwork: BlockchainNetwork{
						ID:               networkID,
						NetworkName:      "Ethereum",
						ChainID:          "1",
						BlockExplorerURL: "https://etherscan.io",
					},
					MerchantUser: MerchantUser{
						ID:    merchantUserID,
						Name:  "Merchant Name",
						Email: "merchant@example.com",
					},
					InvoiceItems: []InvoiceItem{
						{
							ID:          itemID,
							Description: "Product A",
							Quantity:    2,
							UnitPrice:   50.00,
							TotalPrice:  100.00,
							CreatedAt:   "2023-12-01T10:00:00Z",
							UpdatedAt:   "2023-12-01T10:00:00Z",
						},
						{
							ID:          generateMockUUID().String(),
							Description: "Product B",
							Quantity:    1,
							UnitPrice:   100.00,
							TotalPrice:  100.00,
							CreatedAt:   "2023-12-01T10:00:00Z",
							UpdatedAt:   "2023-12-01T10:00:00Z",
						},
					},
				},
			},
		},
	}
}

func generateMockApproveInvoiceRequest() *ApproveInvoice {
	return &ApproveInvoice{
		InvoiceID: generateMockUUID().String(),
		OnChain:   true,
	}
}

func generateMockQueryParams() QueryParams {
	return QueryParams{
		Page:   1,
		Limit:  10,
		Filter: "pending",
	}
}

// Helper function to create a test client
func createTestClient(baseURL string) *Client {
	return NewClient(LongSwipeConfig{
		BaseURL:    baseURL,
		PublicKey:  "test-public-key",
		PrivateKey: "test-private-key",
		Timeout:    5 * time.Second,
	})
}

// Test Suite for Invoice Operations
func TestInvoiceOperations(t *testing.T) {
	// Run all invoice tests in a test suite
	t.Run("CreateInvoice", TestCreateInvoice)
	t.Run("GetInvoice", TestGetInvoice)
	t.Run("ApproveInvoice", TestApproveInvoice)
}

func TestCreateInvoice(t *testing.T) {
	testServer := NewTestServer()
	defer testServer.Close()

	t.Run("Success", func(t *testing.T) {
		// Setup mock response
		mockResponse := generateMockSuccessResponse()
		testServer.AddHandler("/merchant-integrations-server/create-invoice", func(w http.ResponseWriter, r *http.Request) {
			// Verify request method
			assert.Equal(t, http.MethodPost, r.Method)

			// Verify headers
			assert.Equal(t, "Bearer test-public-key", r.Header.Get("Authorization"))
			assert.Equal(t, "test-private-key", r.Header.Get("X-API-Private-Key"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			// Verify request body
			var requestBody CreateInvoiceRequest
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			require.NoError(t, err)
			assert.Equal(t, "John Doe", requestBody.FullName)
			assert.Equal(t, "john.doe@example.com", requestBody.Email)
			assert.Equal(t, "MERCH123", requestBody.MerchantCode)
			assert.Len(t, requestBody.InvoiceItems, 2)

			// Send response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockResponse)
		})

		// Create client and test
		client := createTestClient(testServer.URL())
		request := generateMockInvoiceCreateRequest()

		response, err := client.CreateInvoice(request)

		// Assertions
		require.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "Operation completed successfully", response.Message)
		assert.Equal(t, 200, response.Code)
	})

	t.Run("ValidationError", func(t *testing.T) {
		testServer.AddHandler("/merchant-integrations-server/create-invoice", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"status":"error","message":"Validation failed: Email is required","code":400}`))
		})

		client := createTestClient(testServer.URL())
		request := generateMockInvoiceCreateRequest()
		request.Email = "" // Invalid email

		response, err := client.CreateInvoice(request)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "400")
	})

	t.Run("NetworkError", func(t *testing.T) {
		// Create client with invalid URL
		client := createTestClient("http://invalid-url")
		request := generateMockInvoiceCreateRequest()

		response, err := client.CreateInvoice(request)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func TestGetInvoice(t *testing.T) {
	testServer := NewTestServer()
	defer testServer.Close()

	t.Run("Success", func(t *testing.T) {
		// Setup mock response
		mockResponse := generateMockInvoiceResponse()
		testServer.AddHandler("/merchant-integrations-server/fetch-invoice", func(w http.ResponseWriter, r *http.Request) {
			// Verify request method
			assert.Equal(t, http.MethodGet, r.Method)

			// Verify headers
			assert.Equal(t, "Bearer test-public-key", r.Header.Get("Authorization"))
			assert.Equal(t, "test-private-key", r.Header.Get("X-API-Private-Key"))

			// Verify query parameters
			assert.Equal(t, "1", r.URL.Query().Get("page"))
			assert.Equal(t, "10", r.URL.Query().Get("limit"))
			assert.Equal(t, "pending", r.URL.Query().Get("filter"))

			// Send response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockResponse)
		})

		// Create client and test
		client := createTestClient(testServer.URL())
		queryParams := generateMockQueryParams()

		response, err := client.GetInvoice(queryParams)

		// Assertions
		require.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "Invoices retrieved successfully", response.Message)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, 1, response.Data.Total)
		assert.Len(t, response.Data.Invoices, 1)

		// Verify invoice details
		invoice := response.Data.Invoices[0]
		assert.Equal(t, "INV-001", invoice.InvoiceNumber)
		assert.Equal(t, "pending", invoice.Status)
		assert.Equal(t, 200.00, invoice.TotalAmount)
		assert.Len(t, invoice.InvoiceItems, 2)
		assert.Equal(t, "US Dollar", invoice.Currency.Name)
		assert.Equal(t, "Ethereum", invoice.BlockchainNetwork.NetworkName)
	})

	t.Run("EmptyResult", func(t *testing.T) {
		testServer.AddHandler("/merchant-integrations-server/fetch-invoice", func(w http.ResponseWriter, r *http.Request) {
			emptyResponse := &InvoiceResponse{
				Status:  "success",
				Message: "No invoices found",
				Code:    200,
				Data: InvoiceDetails{
					Total:    0,
					Invoices: []Invoice{},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(emptyResponse)
		})

		client := createTestClient(testServer.URL())
		queryParams := generateMockQueryParams()

		response, err := client.GetInvoice(queryParams)

		// Assertions
		require.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, 0, response.Data.Total)
		assert.Len(t, response.Data.Invoices, 0)
	})

	t.Run("UnauthorizedError", func(t *testing.T) {
		testServer.AddHandler("/merchant-integrations-server/fetch-invoice", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"status":"error","message":"Unauthorized access","code":401}`))
		})

		client := createTestClient(testServer.URL())
		queryParams := generateMockQueryParams()

		response, err := client.GetInvoice(queryParams)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "401")
	})
}

func TestApproveInvoice(t *testing.T) {
	testServer := NewTestServer()
	defer testServer.Close()

	t.Run("Success", func(t *testing.T) {
		// Setup mock response
		mockResponse := generateMockSuccessResponse()
		testServer.AddHandler("/merchant-integrations-server/approve-invoice", func(w http.ResponseWriter, r *http.Request) {
			// Verify request method
			assert.Equal(t, http.MethodPost, r.Method)

			// Verify headers
			assert.Equal(t, "Bearer test-public-key", r.Header.Get("Authorization"))
			assert.Equal(t, "test-private-key", r.Header.Get("X-API-Private-Key"))
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

			// Verify request body
			var requestBody ApproveInvoice
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			require.NoError(t, err)
			assert.NotEmpty(t, requestBody.InvoiceID)
			assert.True(t, requestBody.OnChain)

			// Send response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockResponse)
		})

		// Create client and test
		client := createTestClient(testServer.URL())
		request := generateMockApproveInvoiceRequest()

		response, err := client.ApproveInvoice(request)

		// Assertions
		require.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, "success", response.Status)
		assert.Equal(t, "Operation completed successfully", response.Message)
		assert.Equal(t, 200, response.Code)
	})

	t.Run("InvoiceNotFound", func(t *testing.T) {
		testServer.AddHandler("/merchant-integrations-server/approve-invoice", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status":"error","message":"Invoice not found","code":404}`))
		})

		client := createTestClient(testServer.URL())
		request := generateMockApproveInvoiceRequest()
		request.InvoiceID = "non-existent-id"

		response, err := client.ApproveInvoice(request)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "404")
	})

	t.Run("AlreadyApproved", func(t *testing.T) {
		testServer.AddHandler("/merchant-integrations-server/approve-invoice", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(`{"status":"error","message":"Invoice already approved","code":409}`))
		})

		client := createTestClient(testServer.URL())
		request := generateMockApproveInvoiceRequest()

		response, err := client.ApproveInvoice(request)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "409")
	})
}

// Integration test example
func TestInvoiceWorkflow(t *testing.T) {
	testServer := NewTestServer()
	defer testServer.Close()

	// Setup handlers for complete workflow
	createResponse := generateMockSuccessResponse()
	listResponse := generateMockInvoiceResponse()
	approveResponse := generateMockSuccessResponse()

	testServer.AddHandler("/merchant-integrations-server/create-invoice", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(createResponse)
	})

	testServer.AddHandler("/merchant-integrations-server/fetch-invoice", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listResponse)
	})

	testServer.AddHandler("/merchant-integrations-server/approve-invoice", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(approveResponse)
	})

	client := createTestClient(testServer.URL())

	// Step 1: Create invoice
	createRequest := generateMockInvoiceCreateRequest()
	createResp, err := client.CreateInvoice(createRequest)
	require.NoError(t, err)
	assert.Equal(t, "success", createResp.Status)

	// Step 2: List invoices
	queryParams := generateMockQueryParams()
	listResp, err := client.GetInvoice(queryParams)
	require.NoError(t, err)
	assert.Equal(t, "success", listResp.Status)
	assert.Greater(t, listResp.Data.Total, 0)

	// Step 3: Approve invoice
	approveRequest := &ApproveInvoice{
		InvoiceID: listResp.Data.Invoices[0].ID,
		OnChain:   true,
	}
	approveResp, err := client.ApproveInvoice(approveRequest)
	require.NoError(t, err)
	assert.Equal(t, "success", approveResp.Status)
}

// ============================================================================
// EXAMPLE: Voucher Tests (Extensibility Demo)
// ============================================================================
// Uncomment and modify these functions to add voucher testing capability

/*
// Mock data generators for Voucher operations
func generateMockRedeemRequest() *RedeemRequest {
	return &RedeemRequest{
		VoucherCode:            "VOUCHER123",
		Amount:                 100.50,
		WalletAddress:          "0x742f35Cc6464C5C07fdF6c0B2C36D8Ca6b18Aea1",
		ToCurrencyAbbreviation: "USDT",
		ReferenceId:            generateMockUUID().String(),
		MetaData:               "test metadata",
	}
}

func generateMockVerifyVoucherRequest() *VerifyVoucherCodeRequest {
	return &VerifyVoucherCodeRequest{
		VoucherCode: "VOUCHER123",
	}
}

func generateMockGenerateVoucherRequest() *GenerateVoucherForCustomerRequest {
	return &GenerateVoucherForCustomerRequest{
		CurrencyId:       generateMockUUID(),
		AmountToPurchase: 100.00,
		CustomerID:       generateMockUUID(),
		OnChain:          true,
	}
}

// Test Suite for Voucher Operations
func TestVoucherOperations(t *testing.T) {
	t.Run("RedeemVoucher", TestRedeemVoucher)
	t.Run("VerifyVoucher", TestVerifyVoucher)
	t.Run("GenerateVoucher", TestGenerateVoucher)
	t.Run("GetVoucherRedeemptionCharges", TestGetVoucherRedeemptionCharges)
}

func TestRedeemVoucher(t *testing.T) {
	testServer := NewTestServer()
	defer testServer.Close()

	t.Run("Success", func(t *testing.T) {
		mockResponse := generateMockSuccessResponse()
		testServer.AddHandler("/merchant-integrations/redeem-voucher", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "Bearer test-public-key", r.Header.Get("Authorization"))

			var requestBody RedeemRequest
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			require.NoError(t, err)
			assert.Equal(t, "VOUCHER123", requestBody.VoucherCode)
			assert.Equal(t, 100.50, requestBody.Amount)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockResponse)
		})

		client := createTestClient(testServer.URL())
		request := generateMockRedeemRequest()

		response, err := client.RedeemVoucher(request)

		require.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, "success", response.Status)
	})

	t.Run("InvalidVoucher", func(t *testing.T) {
		testServer.AddHandler("/merchant-integrations/redeem-voucher", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"status":"error","message":"Invalid voucher code","code":400}`))
		})

		client := createTestClient(testServer.URL())
		request := generateMockRedeemRequest()
		request.VoucherCode = "INVALID"

		response, err := client.RedeemVoucher(request)

		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Contains(t, err.Error(), "400")
	})
}

func TestVerifyVoucher(t *testing.T) {
	// Similar structure to other tests...
	// Implementation would follow the same pattern
}

func TestGenerateVoucher(t *testing.T) {
	// Similar structure to other tests...
	// Implementation would follow the same pattern
}

func TestGetVoucherRedeemptionCharges(t *testing.T) {
	// Similar structure to other tests...
	// Implementation would follow the same pattern
}

// Voucher workflow integration test
func TestVoucherWorkflow(t *testing.T) {
	testServer := NewTestServer()
	defer testServer.Close()

	// Setup handlers for complete voucher workflow
	// 1. Generate voucher
	// 2. Verify voucher
	// 3. Get redemption charges
	// 4. Redeem voucher

	// Implementation would follow the same pattern as TestInvoiceWorkflow
}
*/
