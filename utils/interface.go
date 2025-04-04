package utils

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserRoles string
type NetworkType string

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
	ID                        uuid.UUID         `json:"id"`
	Amount                    float64           `json:"amount"`
	Balance                   float64           `json:"balance"`
	GeneratedCurrency         CurrencyDetails   `json:"generatedCurrency"`
	Code                      string            `json:"code"`
	WasPaidFor                bool              `json:"wasPaidFor"`
	IsUsed                    bool              `json:"isUsed"`
	CreatedAt                 time.Time         `json:"createdAt"`
	CreatedForMerchant        bool              `json:"createdForMerchant"`
	CreatedForExistingUser    bool              `json:"createdForExistingUser"`
	CreatedForNonExistingUser bool              `json:"createdForNonExistingUser"`
	IsLocked                  bool              `json:"isLocked"`
	OnChain                   bool              `json:"onchain"`
	OnChainProcessing         bool              `json:"onchainProcessing"`
	CryptoVoucherDetails      *CryptoVoucher    `json:"cryptoVoucherDetails"`
	RedeemedVouchers          []RedeemedVoucher `json:"redeemedVouchers"`
	TransactionHash           string            `json:"transactionHash"`
	MetaData                  string            `json:"metaData"`
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

type CustomersResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Data    struct {
		Total     int64          `json:"total"`
		Page      int            `json:"page"`
		Limit     int            `json:"limit"`
		Customers []CustomerData `json:"customer"`
	} `json:"data"`
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
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type GenerateVoucherForCustomerRequest struct {
	BlockchainNetworkId uuid.UUID `json:"blockchainNetworkId" validate:"omitempty,uuid"`
	CurrencyId          uuid.UUID `json:"currencyId" validate:"required"`
	AmountToPurchase    float64   `json:"amountToPurchase" validate:"required"`
	CustomerID          uuid.UUID `json:"customerId" validate:"required"`
	OnChain             bool      `json:"onChain" validate:"omitempty"`
}
