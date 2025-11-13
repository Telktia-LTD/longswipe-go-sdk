package longswipe

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserRoles string
type NetworkType string
type MERCHANTROLES string
type TransactionType string
type TransactionStatus string

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
	Data    T      `json:"data,omitempty"`
}

type RedeemRequest struct {
	VoucherCode            string  `json:"voucherCode" validate:"required"`
	Amount                 float64 `json:"amount" validate:"required"`
	LockPin                string  `json:"lockPin" validate:"omitempty"`
	WalletAddress          string  `json:"walletAddress" validate:"omitempty"`
	ToCurrencyAbbreviation string  `json:"toCurrencyAbbreviation" validate:"omitempty"`
	ReferenceId            string  `json:"referenceId" validate:"omitempty"`
	MetaData               string  `json:"metaData" validate:"omitempty"`
}

type UserResponse struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Surname       string    `json:"surname"`
	Othernames    string    `json:"otherNames"`
	RegChannel    string    `json:"regChannel"`
	ExternalID    string    `json:"externalID"`
	Role          UserRoles `json:"role"`
	IsActive      bool      `json:"isActive"`
	EmailVerified bool      `json:"emailVerified"`
	Avatar        string    `json:"avatar"`
	IsPinSet      bool      `json:"isPinSet"`
}

type VoucherResponse struct {
	ID                        uuid.UUID       `json:"id"`
	Amount                    float64         `json:"amount"`
	Balance                   float64         `json:"balance"`
	GeneratedCurrency         CurrencyDetails `json:"generatedCurrency"`
	Code                      string          `json:"code"`
	WasPaidFor                bool            `json:"wasPaidFor"`
	IsUsed                    bool            `json:"isUsed"`
	CreatedAt                 time.Time       `json:"createdAt"`
	CreatedForMerchant        bool            `json:"createdForMerchant"`
	CreatedForExistingUser    bool            `json:"createdForExistingUser"`
	CreatedForNonExistingUser bool            `json:"createdForNonExistingUser"`
	IsLocked                  bool            `json:"isLocked"`
	OnChain                   bool            `json:"onchain"`
	OnChainProcessing         bool            `json:"onchainProcessing"`
	CryptoVoucherDetails      *CryptoVoucher  `json:"cryptoVoucherDetails"`
	TransactionHash           string          `json:"transactionHash"`
	MetaData                  string          `json:"metaData"`
}
type RedeemedVoucher struct {
	ID                    uuid.UUID    `json:"id"`
	RedeemedUserID        *uuid.UUID   `json:"redeemedUserID"`
	RedeemerWalletAddress *string      `json:"redeemerWalletAddress"`
	VoucherID             uuid.UUID    `json:"voucherID"`
	User                  UserResponse `json:"user"`
	Amount                float64      `json:"amount"`
	IsMerchant            bool         `json:"isMerchant"`
	CreatedAt             time.Time    `json:"createdAt"`
}
type CryptoVoucher struct {
	CodeHash        string
	Value           string
	Balance         string
	Creator         string
	IsRedeemed      bool
	TransactionHash string
}

type V2PayoutDetailsResponse struct {
	Amount                                 float64         `json:"swapAmount"` // Amount requested to swap
	ToAmount                               float64         `json:"toAmount"`   // Amount to receive
	ProcessingFee                          float64         `json:"processingFee"`
	TotalGasAndProceesingFeeInFromCurrency float64         `json:"totalGasAndProceesingFeeInFromCurrency"`
	TotalGasCostAndProcessingFeeInWei      float64         `json:"totalGasCostAndProcessingFeeInWei"`
	ExchangeRate                           float64         `json:"exchangeRate"`
	PercentageCharge                       float64         `json:"percentageCharge"`
	IsPercentageCharge                     bool            `json:"isPercentageCharge"`
	ToCurrency                             CurrencyDetails `json:"toCurrency"`
	FromCurrency                           CurrencyDetails `json:"fromCurrency"`
	TotalDeductable                        float64         `json:"totalDeductable"`
}

type RedeemVoucherDetailDataAll struct {
	Charges V2PayoutDetailsResponse `json:"charges"`
	Voucher VoucherResponse         `json:"voucher"`
}

type RedeemeVoucherDetailsResponse = ApiResponse[RedeemVoucherDetailDataAll]

type VoucherPurchaseChargesDetails struct {
	Amount                            float64 `json:"amount"`
	AmountInWei                       float64 `json:"amountInWei"`
	GasPriceInWei                     float64 `json:"gasPriceInWei"`
	GasLimitInWei                     float64 `json:"gasLimitInWei"`
	TotalGasCostInWei                 float64 `json:"totalGasCostInWei"`
	ProcessingFeeInWei                float64 `json:"processingFeeInWei"`
	BalanceAfterChargesInWei          float64 `json:"balanceAfterChargesInWei"`
	TotalGasCostAndProcessingFeeInWei float64 `json:"totalGasCostAndProcessingFeeInWei"`
	TotalGasCost                      float64 `json:"totalGasCost"`
	ProcessingFee                     float64 `json:"processingFee"`
	TotalGasCostAndProcessingFee      float64 `json:"totalGasCostAndProcessingFee"`
	BalanceAfterCharges               float64 `json:"balanceAfterCharges"`
}

type VerifyVoucherCodeRequest struct {
	VoucherCode string `json:"voucherCode" validate:"required"`
}

type VerifyVoucherResponse = ApiResponse[VoucherResponse]

type CryptoNetworkDetails struct {
	ID               uuid.UUID        `json:"id"`
	NetworkName      string           `json:"networkName"`
	RpcUrl           string           `json:"rpcUrl"`
	ChainID          string           `json:"chainID"`
	BlockExplorerUrl string           `json:"blockExplorerUrl"`
	CryptoCurrencies []CryptoCurrency `json:"cryptocurrencies"`
	NetworkType      NetworkType      `json:"networkType"`
}

type CurrencyDetails struct {
	ID           uuid.UUID `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	Abbreviation string    `json:"Abbreviation"`
	CurrencyType string    `json:"currencyType"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt"`
}

type CryptoCurrency struct {
	ID                       uuid.UUID       `json:"id"`
	CurrencyData             CurrencyDetails `json:"currencyData"`
	CurrencyAddress          string          `json:"currencyAddress"`
	LongswipeContractAddress string          `json:"longswipeContractAddress"`
	CurrencyName             string          `json:"currencyName"`
	CurrencyDecimals         string          `json:"currencyDecimals"`
	NetworkID                uuid.UUID       `json:"networkID"`
	Status                   bool            `json:"status"`
}

type CryptoNetworkResponse = ApiResponse[[]CryptoNetworkDetails]

type FetchCurrenciesResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Currencies []Currencies `json:"currencies"`
	} `json:"data"`
}

type Currencies struct {
	ID           uuid.UUID `json:"id"`
	Image        string    `json:"image"`
	Currency     string    `json:"currency"`
	Symbol       string    `json:"symbol"`
	Abbreviation string    `json:"abbreviation"`
	IsActive     bool      `json:"isActive"`
	CurrencyType string    `json:"currencyType"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Pagination struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type CustomerData struct {
	ID         uuid.UUID `json:"id"`
	MerchantID uuid.UUID `json:"merchantID"`
	Name       string    `json:"name" validate:"required"`
	Email      string    `json:"email" validate:"email,required"`
}

type CustomerDetails struct {
	Total     int64          `json:"total"`
	Page      int            `json:"page"`
	Limit     int            `json:"limit"`
	Customers []CustomerData `json:"customer"`
}

type CustomersResponse struct {
	Message string          `json:"message"`
	Code    int             `json:"code"`
	Status  string          `json:"status"`
	Data    CustomerDetails `json:"data"`
}

type CustomerResponse struct {
	Message string       `json:"message"`
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Data    CustomerData `json:"customer"`
}

type AddNewCustomer struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type UpdatCustomer struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Name  string    `json:"name" validate:"required"`
	Email string    `json:"email" validate:"email,required"`
}

type GenerateVoucherForCustomerRequest struct {
	BlockchainNetworkId uuid.UUID `json:"blockchainNetworkId" validate:"omitempty,uuid"`
	CurrencyId          uuid.UUID `json:"currencyId" validate:"required"`
	AmountToPurchase    float64   `json:"amountToPurchase" validate:"required"`
	CustomerID          uuid.UUID `json:"customerId" validate:"required"`
	OnChain             bool      `json:"onChain" validate:"omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type CreateInvoiceRequest struct {
	FullName            string               `json:"fullName" validate:"required"`
	Email               string               `json:"email" validate:"email,required"`
	MerchantCode        string               `json:"merchantCode" validate:"required"`
	InvoiceDate         time.Time            `json:"invoiceDate" validate:"required"`
	DueDate             time.Time            `json:"dueDate" validate:"required"`
	InvoiceItems        []InvoiceItemRequest `json:"invoiceItems" validate:"required"`
	CurrencyId          uuid.UUID            `json:"currencyId" validate:"required"`
	BlockchainNetworkId uuid.UUID            `json:"blockchainNetworkId" validate:"omitempty"`
}

type InvoiceItemRequest struct {
	Description string  `json:"description" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	UnitPrice   float64 `json:"unitPrice" validate:"required"`
}

type ApproveInvoiceRequest struct {
	InvoiceID uuid.UUID `json:"invoiceID" validate:"required"`
	OnChain   bool      `json:"onChain" validate:"omitempty"`
}

type MerchantInvoiceResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Invoices []Invoice `json:"invoices"`
		Total    int       `json:"total"`
	} `json:"data"`
}

type NetworkDetails struct {
	ID               uuid.UUID `json:"id"`
	NetworkName      string    `json:"networkName"`
	ChainID          string    `json:"chainID"`
	BlockExplorerUrl string    `json:"blockExplorerUrl"`
}

type MerchantUserDatas struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type InvoiceItem struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unitPrice"`
	TotalPrice  float64   `json:"totalPrice"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Invoice struct {
	ID                uuid.UUID       `json:"id"`
	InvoiceNumber     string          `json:"invoiceNumber"`
	UserId            *uuid.UUID      `json:"userId"`
	Email             string          `json:"email"`
	FullName          string          `json:"fullName"`
	InvoiceDate       time.Time       `json:"invoiceDate"`
	DueDate           time.Time       `json:"dueDate"`
	TotalAmount       float64         `json:"totalAmount"`
	Status            string          `json:"status"`
	InvoiceItems      []InvoiceItem   `json:"invoiceItems"`
	Currency          CurrencyDetails `json:"currency"`
	BlockchainNetwork *NetworkDetails `json:"blockchainNetwork"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
}

type AllowedInvoiceCurrency struct {
	Currency CurrencyDetails `json:"currency" validate:"omitempty"`
}

type FetchAllAllowedInvoiceCurrencyResponse struct {
	Message string                   `json:"message"`
	Code    int                      `json:"code"`
	Status  string                   `json:"status"`
	Data    []AllowedInvoiceCurrency `json:"data" validate:"omitempty,uuid"`
}

type AddNewUserRequest struct {
	Name  string        `json:"name" validate:"required"`
	Email string        `json:"email" validate:"email,required"`
	Role  MERCHANTROLES `json:"role" validate:"required"`
}

type MerchantUserData struct {
	ID         uuid.UUID     `json:"id"`
	MerchantID uuid.UUID     `json:"merchantID"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Team       string        `json:"team"`
	Role       MERCHANTROLES `json:"role"`
}

type MerchantUserResponse struct {
	Message string             `json:"message"`
	Code    int                `json:"code"`
	Status  string             `json:"status"`
	Data    []MerchantUserData `json:"data"`
}

type PaymentRequest struct {
	Amount         float64                `json:"amount"`
	Currency       string                 `json:"currency"`
	UserIdentifier string                 `json:"user_identifier"`
	Metadata       map[string]interface{} `json:"metadata"`
	ReferenceID    string                 `json:"reference_id"`
}

type AddressDepositRequest struct {
	Amount                      float64                `json:"amount"`
	BlockchainNetworkID         string                 `json:"blockchainNetworkId"`
	CurrencyAbbreviation        string                 `json:"currency_abbreviation"`
	Metadata                    map[string]interface{} `json:"metadata"`
	PayWithCurrencyAbbreviation string                 `json:"pay_with_currency_abbreviation"`
	ReferenceID                 string                 `json:"reference_id"`
}

type DepositResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		ID                      string  `json:"id"`
		Address                 string  `json:"address"`
		AmountToDeposit         float64 `json:"amountToDeposit"`
		ExpiresAt               string  `json:"expiresAt"`   // ISO 8601 date string
		DateCreated             string  `json:"dateCreated"` // ISO 8601 date string
		BlockchainNetworkDetail struct {
			ID               string `json:"id"`
			NetworkName      string `json:"networkName"`
			ChainID          string `json:"chainID"`
			BlockExplorerURL string `json:"blockExplorerUrl"`
			NetworkType      string `json:"networkType"`
			NetworkLogo      string `json:"networkLogo"`
		} `json:"blockchainNetworkDetail"`
	} `json:"data"`
}

type AddressDepositChargeRequest struct {
	Amount                      float64 `json:"amount"`
	BlockchainNetworkID         string  `json:"blockchainNetworkId"`
	CurrencyAbbreviation        string  `json:"currency_abbreviation"`
	PayWithCurrencyAbbreviation string  `json:"pay_with_currency_abbreviation"`
}

type ChargeEstimateResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		SwapAmount                             float64        `json:"swapAmount"`
		ToAmount                               float64        `json:"toAmount"`
		ProcessingFee                          float64        `json:"processingFee"`
		TotalGasAndProcessingFeeInFromCurrency float64        `json:"totalGasAndProceesingFeeInFromCurrency"`
		TotalGasCostAndProcessingFeeInWei      float64        `json:"totalGasCostAndProcessingFeeInWei"`
		ExchangeRate                           float64        `json:"exchangeRate"`
		PercentageCharge                       float64        `json:"percentageCharge"`
		IsPercentageCharge                     bool           `json:"isPercentageCharge"`
		ToCurrency                             CurrencyDetail `json:"toCurrency"`
		FromCurrency                           CurrencyDetail `json:"fromCurrency"`
		TotalDeductable                        float64        `json:"totalDeductable"`
	} `json:"data"`
}

type CurrencyDetail struct {
	ID           string `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Abbreviation string `json:"Abbreviation"`
	CurrencyType string `json:"currencyType"`
	IsActive     bool   `json:"isActive"`
	CreatedAt    string `json:"createdAt"`
}

type TransactionResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		ID              string         `json:"id"`
		UserID          *string        `json:"userId"` // Use pointer to handle string or null
		ReferenceID     string         `json:"referenceId"`
		Amount          float64        `json:"amount"`
		Title           string         `json:"title"`
		Message         string         `json:"message"`
		ChargedAmount   float64        `json:"chargedAmount"`
		ChargeType      string         `json:"chargeType"`
		Type            string         `json:"type"`
		Status          string         `json:"status"`
		Currency        CurrencyDetail `json:"currency"`
		CreatedAt       string         `json:"createdAt"` // ISO 8601 date string
		UpdatedAt       string         `json:"updatedAt"` // ISO 8601 date string
		TransactionHash string         `json:"transactionHash"`
		ApplicationName string         `json:"applicationName"`
		ReferenceHash   string         `json:"referenceHash"`
		MetaData        string         `json:"metaData"`
	} `json:"data"`
}

type ConfirmUserDetails struct {
	Fullname string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Phone    string `json:"phone" validate:"required"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
}

type ConfirmUserDetailsResponse struct {
	Message string             `json:"message"`
	Code    int                `json:"code"`
	Status  string             `json:"status"`
	Data    ConfirmUserDetails `json:"data"`
}

type CustomerPayout struct {
	Amount                        float64 `json:"amount" validate:"required"`
	MetaData                      string  `json:"metaData" validate:"omitempty"`
	FromCurrencyAbbreviation      string  `json:"fromCurrencyAbbreviation" validate:"omitempty"`
	ToCurrencyAbbreviation        string  `json:"toCurrencyAbbreviation" validate:"omitempty"`
	ReferenceId                   string  `json:"referenceId" validate:"omitempty"`
	LongswipeUsernameOrEmail      string  `json:"longswipeUsernameOrEmail" validate:"omitempty"`
	BlockchainNetworkAbbreviation string  `json:"blockchainNetworkAbbreviation" validate:"omitempty"`
}

type Transactions struct {
	ID              uuid.UUID         `json:"id"`
	UserID          *uuid.UUID        `json:"userId"`
	ReferenceID     string            `json:"referenceId"`
	Amount          float64           `json:"amount"`
	Title           string            `json:"title"`
	Message         string            `json:"message"`
	ChargedAmount   float64           `json:"chargedAmount"`
	ChargeType      TransactionType   `json:"chargeType"`
	Type            TransactionType   `json:"type"`
	Status          TransactionStatus `json:"status"`
	Currency        CurrencyDetails   `json:"currency"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
	TransactionHash string            `json:"transactionHash"`
	ApplicationName string            `json:"applicationName"`
	ReferenceHash   string            `json:"referenceHash"`
	MetaData        string            `json:"metaData"`
}
type PaginationInfo struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}
type TransactionListResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Transactions []Transactions `json:"transactions"`
		Pagination   PaginationInfo `json:"pagination"`
	} `json:"data"`
}

type MerchantBalanceDetails struct {
	Merchant PublicMerchantResponse `json:"merchant"`
	Balance  float64                `json:"balance"`
	Currency CurrencyDetails        `json:"currency"`
}

type PublicBalanceResponse struct {
	Message string                 `json:"message"`
	Code    int                    `json:"code"`
	Status  string                 `json:"status"`
	Data    MerchantBalanceDetails `json:"data"`
}

type PublicMerchantResponse struct {
	CompanyName         string `json:"companyName"`
	TradingName         string `json:"tradingName"`
	MerchantDescription string `json:"merchantDescription"`
	MerchantCode        string `json:"merchantCode"`
	Avatar              string `json:"avatar"`
}

type EscrowPublicRequest struct {
	Amount                    float64 `json:"amount" validate:"required,gt=0"`
	CurrencyAbbreviation      string  `json:"currency_abbreviation" validate:"required"`
	ApproverEmail             string  `json:"approver_email" validate:"required,email"`
	ApproverAuthorizationCode string  `json:"approver_authorization_code" validate:"required"`
}

type EscrowInitialDetails struct {
	EscrowId string `json:"escrow_id"`
}

type EscrowInitialResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Code    int                  `json:"code"`
	Data    EscrowInitialDetails `json:"data"`
}

type UpdateEscrowStatusRequest struct {
	Status   string `json:"status" validate:"required,oneof=COMPLETED CANCELLED FAILED APPROVED"`
	EscrowId string `json:"escrow_id" validate:"required"`
}

type AllEscrowDetail struct {
	EscrowId   string            `json:"escrow_id"`
	Amount     float64           `json:"amount"`
	Balance    float64           `json:"balance"`
	Currency   CurrencyDetails   `json:"currency"`
	Status     TransactionStatus `json:"status"`
	IsReleased bool              `json:"is_released"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type AllEscrowResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		Escrows    []AllEscrowDetail `json:"escrows"`
		Pagination PaginationInfo    `json:"pagination_info"`
	} `json:"data"`
}

type EscrowDetailResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Code    int          `json:"code"`
	Data    EscrowDetail `json:"data"`
}

type EscrowDetail struct {
	EscrowId        string               `json:"escrow_id"`
	Amount          float64              `json:"amount"`
	Status          TransactionStatus    `json:"status"`
	Balance         float64              `json:"balance"`
	IsReleased      bool                 `json:"is_released"`
	ApplicationName string               `json:"application_name"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	Recipients      []RecipientDetail    `json:"recipients"`
	Approver        RoleDetail           `json:"approver"`
	Transactions    []EscrowTransactions `json:"transactions"`
	FundRelease     []FundReleaseDetail  `json:"fund_release"`
}

type RecipientDetail struct {
	Email         string                 `json:"email"`
	Role          string                 `json:"role"`
	PayoutDetails RecipientPayoutDetails `json:"payout_details,omitempty"`
}

type RecipientPayoutDetails struct {
	Currency    EscrowCurrencyDetails `json:"currency,omitempty"`
	Fiat        FiatDetails           `json:"fiat_details,omitempty"`
	Crypto      Crypto                `json:"crypto_details,omitempty"`
	IsConfirmed bool                  `json:"is_confirmed"`
}

type EscrowCurrencyDetails struct {
	Image        string `json:"image"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Abbreviation string `json:"abbreviation"`
	CurrencyType string `json:"currencyType"`
}

type FiatDetails struct {
	BankName      string `json:"bank"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
}

type Crypto struct {
	NetworkName         string `json:"network_name,omitempty"`
	NetworkAbbreviation string `json:"network_abbreviation,omitempty"`
	NetworkLogo         string `json:"network_logo,omitempty"`
	WalletAddress       string `json:"wallet_address,omitempty"`
}

type RoleDetail struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type EscrowTransactions struct {
	ReferenceID     string                `json:"referenceId"`
	Amount          float64               `json:"amount"`
	Title           string                `json:"title"`
	Message         string                `json:"message"`
	ChargedAmount   float64               `json:"chargedAmount"`
	ChargeType      TransactionType       `json:"chargeType"`
	Type            TransactionType       `json:"type"`
	Status          TransactionStatus     `json:"status"`
	Currency        EscrowCurrencyDetails `json:"currency"`
	CreatedAt       time.Time             `json:"createdAt"`
	UpdatedAt       time.Time             `json:"updatedAt"`
	TransactionHash string                `json:"transactionHash"`
	ApplicationName string                `json:"applicationName"`
	ReferenceHash   string                `json:"referenceHash"`
	MetaData        string                `json:"metaData"`
}

type FundReleaseDetail struct {
	Id              int64                 `json:"id"`
	Amount          float64               `json:"amount"`
	Status          TransactionStatus     `json:"status"`
	Charges         float64               `json:"charges"`
	TotalDeductable float64               `json:"total_deductable"`
	IsReleased      bool                  `json:"is_released"`
	Approved        bool                  `json:"approved"`
	Currency        EscrowCurrencyDetails `json:"currency"`
	CreatedAt       time.Time             `json:"created_at"`
	UpdatedAt       time.Time             `json:"updated_at"`
}

type RequestOtp struct {
	EscrowId string `json:"escrow_id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type AddEscrowRecipient struct {
	EscrowId          string `json:"escrow_id" validate:"required"`
	AuthorizationCode string `json:"authorization_code" validate:"omitempty"`
	Email             string `json:"email" validate:"omitempty,email"`
}

type RequestFundReleaseRequest struct {
	EscrowId          string  `json:"escrow_id" validate:"required"`
	Amount            float64 `json:"amount" validate:"required,gt=0"`
	RecipientEmail    string  `json:"recipient_email" validate:"required,email"`
	MetaData          string  `json:"metadata" validate:"omitempty"`
	Otp               string  `json:"otp" validate:"required"`
	AuthorizationCode string  `json:"authorization_code" validate:"required"`
}

type ConfirmFundRelease struct {
	EscrowId          string `json:"escrow_id" validate:"required"`
	FundRequestId     int64  `json:"fund_request_id" validate:"required"`
	Otp               string `json:"otp" validate:"required"`
	AuthorizationCode string `json:"authorization_code" validate:"required"`
}

type SystemRelease struct {
	EscrowId      string `json:"escrow_id" validate:"required"`
	FundRequestId int64  `json:"fund_request_id" validate:"required"`
}

type UpdateAuthorizationCodeRequest struct {
	EscrowId             string `json:"escrow_id" validate:"required"`
	EscrowUserEmail      string `json:"escrow_user_email" validate:"required"`
	NewAuthorizationCode string `json:"new_authorization_code" validate:"required"`
	Otp                  string `json:"otp" validate:"omitempty"`
	OldAuthorizationCode string `json:"old_authorization_code" validate:"omitempty"`
}

type AddPayoutDetailsRequest struct {
	EscrowId                      string `json:"escrow_id" validate:"required"`
	Email                         string `json:"email" validate:"required,email"`
	CurrencyAbbreviation          string `json:"currency_abbreviation" validate:"required"`
	WalletAddress                 string `json:"wallet_address" validate:"omitempty"`
	BankCode                      string `json:"bank_code" validate:"omitempty"`
	AccountNumber                 string `json:"account_number" validate:"omitempty"`
	AccountName                   string `json:"account_name" validate:"omitempty"`
	SortCode                      string `json:"sort_code" validate:"omitempty"`
	BlockchainNetworkAbbreviation string `json:"blockchain_network_abbreviation" validate:"omitempty"`
	Save                          bool   `json:"save" validate:"omitempty"`
}

type RecipientAccountDetailsResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Code    int                     `json:"code"`
	Data    RecipientAccountDetails `json:"data"`
}

type RecipientAccountDetails struct {
	Email     string                `json:"email"`
	Role      string                `json:"role"`
	Confirmed bool                  `json:"confirmed"`
	Currency  EscrowCurrencyDetails `json:"currency"`
	Fiat      FiatDetails           `json:"account_details,omitempty"`
	Crypto    Crypto                `json:"crypto_details,omitempty"`
}

type FundRequestDetails struct {
	EscrowId       string         `json:"escrow_id"`
	PaymentLink    string         `json:"payment_link"`
	PaymentDetails PaymentDetails `json:"payment_details"`
}

type FundRequestResponse struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Code    int                `json:"code"`
	Data    FundRequestDetails `json:"data"`
}

type PaymentDetails struct {
	ID                   uuid.UUID         `json:"id"`
	ApplicationName      string            `json:"application_name"`
	Amount               float64           `json:"amount"`
	Status               TransactionStatus `json:"status"`
	CurrencyAbbreviation string            `json:"currency_abbreviation"`
	CreatedAt            time.Time         `json:"created_at"`
	Identifier           string            `json:"identifier"`
	CurrencyDetails      *CurrencyDetails  `json:"currency"`
}
