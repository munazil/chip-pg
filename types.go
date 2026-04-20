package chippg

type DeliveryMethod struct {
	Method  string                 `json:"method"`
	Options map[string]interface{} `json:"options"`
}

type Client struct {
	Email                 string           `json:"email"`
	Phone                 string           `json:"phone"`
	FullName              string           `json:"full_name"`
	PersonalCode          string           `json:"personal_code"`
	LegalName             string           `json:"legal_name"`
	BrandName             string           `json:"brand_name"`
	RegistrationNumber    string           `json:"registration_number"`
	TaxNumber             string           `json:"tax_number"`
	BankAccount           string           `json:"bank_account"`
	BankCode              string           `json:"bank_code"`
	StreetAddress         string           `json:"street_address"`
	City                  string           `json:"city"`
	ZipCode               string           `json:"zip_code"`
	Country               string           `json:"country"`
	State                 string           `json:"state"`
	ShippingStreetAddress string           `json:"shipping_street_address"`
	ShippingCity          string           `json:"shipping_city"`
	ShippingZipCode       string           `json:"shipping_zip_code"`
	ShippingCountry       string           `json:"shipping_country"`
	ShippingState         string           `json:"shipping_state"`
	CC                    []string         `json:"cc"`
	BCC                   []string         `json:"bcc"`
	DeliveryMethods       []DeliveryMethod `json:"delivery_methods"`
}

type Product struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Quantity   string `json:"quantity,omitempty"`
	Discount   int    `json:"discount,omitempty"`
	TaxPercent string `json:"tax_percent,omitempty"`
	Category   string `json:"category,omitempty"`
}

type Purchase struct {
	Currency              string    `json:"currency"`
	Products              []Product `json:"products"`
	Language              string    `json:"language"`
	Notes                 string    `json:"notes"`
	Debt                  int       `json:"debt"`
	SubtotalOverride      int       `json:"subtotal_override"`
	TotalTaxOverride      int       `json:"total_tax_override"`
	TotalDiscountOverride int       `json:"total_discount_override"`
	TotalOverride         int       `json:"total_override"`
	Total                 int       `json:"total"`
	// RequestClientDetails  []interface{} `json:"request_Client_details"`
	Timezone     string `json:"timezone"`
	DueStrict    bool   `json:"due_strict"`
	EmailMessage string `json:"email_message"`
}

type IssuerBankAccount struct {
	BankAccount string `json:"bank_account"`
	BankCode    string `json:"bank_code"`
}

type IssuerDetails struct {
	BrandName          string              `json:"brand_name"`
	Website            string              `json:"website"`
	LegalName          string              `json:"legal_name"`
	RegistrationNumber string              `json:"registration_number"`
	TaxNumber          string              `json:"tax_number"`
	LegalStreetAddress string              `json:"legal_street_address"`
	LegalCountry       string              `json:"legal_country"`
	LegalCity          string              `json:"legal_city"`
	LegalZipCode       string              `json:"legal_zip_code"`
	BankAccounts       []IssuerBankAccount `json:"bank_accounts"`
}

type TransactionData struct {
	PaymentMethod string                 `json:"payment_method"`
	Extra         map[string]interface{} `json:"extra"`
	Country       string                 `json:"country"`
	Attempts      []TransactionAttempt   `json:"attempts"`
}

type TransactionAttempt struct {
	Type          string `json:"type"`
	Successful    bool   `json:"successful"`
	PaymentMethod string `json:"payment_method"`
	// Extra          map[string]interface{} `json:"extra"`
	Country        string        `json:"country"`
	ClientIP       string        `json:"client_ip"`
	ProcessingTime int64         `json:"processing_time"`
	Error          *AttemptError `json:"error"`
}

type Payment struct {
	IsOutgoing        bool   `json:"is_outgoing"`
	PaymentType       string `json:"payment_type"`
	Amount            int    `json:"amount"`
	Currency          string `json:"currency"`
	NetAmount         int    `json:"net_amount"`
	FeeAmount         int    `json:"fee_amount"`
	PendingAmount     int    `json:"pending_amount"`
	PendingUnfreezeOn int64  `json:"pending_unfreeze_on"`
	Description       string `json:"description"`
	PaidOn            int64  `json:"paid_on"`
	RemotePaidOn      int64  `json:"remote_paid_on"`
}

type AttemptError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type StatusHistoryItem struct {
	Status        string        `json:"status"`
	Timestamp     int64         `json:"timestamp"`
	RelatedObject RelatedObject `json:"related_object"`
}

type RelatedObject struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type CurrencyConversion struct {
	OriginalCurrency string  `json:"original_currency"`
	OriginalAmount   int     `json:"original_amount"`
	ExchangeRate     float64 `json:"exchange_rate"`
}
