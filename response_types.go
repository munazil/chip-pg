package chippg

type PurchaseResponse struct {
	ID                     string              `json:"id"`
	Purchase               Purchase            `json:"purchase"`
	BrandID                string              `json:"brand_id"`
	Type                   string              `json:"type"`
	CreatedOn              int64               `json:"created_on"`
	UpdatedOn              int64               `json:"updated_on"`
	Payment                Payment             `json:"payment"`
	IssuerDetails          IssuerDetails       `json:"issuer_details"`
	TransactionData        TransactionData     `json:"transaction_data"`
	Status                 Status              `json:"status"`
	StatusHistory          []StatusHistoryItem `json:"status_history"`
	ViewedOn               int64               `json:"viewed_on"`
	CompanyID              string              `json:"company_id"`
	IsTest                 bool                `json:"is_test"`
	UserID                 *string             `json:"user_id"`
	BillingTemplateID      *string             `json:"billing_template_id"`
	ClientID               *string             `json:"client_id"`
	SendReceipt            bool                `json:"send_receipt"`
	IsRecurringToken       bool                `json:"is_recurring_token"`
	RecurringToken         *string             `json:"recurring_token"`
	SkipCapture            bool                `json:"skip_capture"`
	ForceRecurring         bool                `json:"force_recurring"`
	ReferenceGenerated     string              `json:"reference_generated"`
	Reference              string              `json:"reference"`
	Notes                  string              `json:"notes"`
	Issued                 string              `json:"issued"`
	Due                    int64               `json:"due"`
	RefundAvailability     string              `json:"refund_availability"`
	RefundableAmount       int                 `json:"refundable_amount"`
	CurrencyConversion     *CurrencyConversion `json:"currency_conversion"`
	PaymentMethodWhitelist []string            `json:"payment_method_whitelist"`
	SuccessRedirect        string              `json:"success_redirect"`
	FailureRedirect        string              `json:"failure_redirect"`
	CancelRedirect         string              `json:"cancel_redirect"`
	SuccessCallback        *string             `json:"success_callback"`
	CreatorAgent           string              `json:"creator_agent"`
	Platform               string              `json:"platform"`
	Product                string              `json:"product"`
	CreatedFromIP          string              `json:"created_from_ip"`
	InvoiceURL             *string             `json:"invoice_url"`
	CheckoutURL            string              `json:"checkout_url"`
	DirectPostURL          *string             `json:"direct_post_url"`
	MarkedAsPaid           bool                `json:"marked_as_paid"`
	OrderID                *string             `json:"order_id"`
}

type RefundPurchaseResponse struct {
	ID                 string          `json:"id"`
	Type               string          `json:"type"`
	CreatedOn          int64           `json:"created_on"`
	UpdatedOn          int64           `json:"updated_on"`
	Client             Client          `json:"client"`
	Payment            Payment         `json:"payment"`
	TransactionData    TransactionData `json:"transaction_data"`
	ReferenceGenerated string          `json:"reference_generated"`
	Reference          string          `json:"reference"`
	AccountID          string          `json:"account_id"`
	CompanyID          string          `json:"company_id"`
	IsTest             bool            `json:"is_test"`
	UserID             string          `json:"user_id"`
	BrandID            string          `json:"brand_id"`
}
