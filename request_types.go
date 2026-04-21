package chippg

type createPurchaseRequest struct {
	Client                 CreatePurchaseRequestClient   `json:"client"`
	Purchase               CreatePurchaseRequestPurchase `json:"purchase"`
	BrandID                string                        `json:"brand_id"`
	SuccessRedirect        string                        `json:"success_redirect"`
	FailureRedirect        string                        `json:"failure_redirect"`
	CancelRedirect         string                        `json:"cancel_redirect"`
	SuccessCallback        *string                       `json:"success_callback,omitempty"`
	PaymentMethodWhitelist *[]string                     `json:"payment_method_whitelist,omitempty"`
}

type CreatePurchaseRequest struct {
	Client                 CreatePurchaseRequestClient   `json:"client"`
	Purchase               CreatePurchaseRequestPurchase `json:"purchase"`
	SuccessRedirect        string                        `json:"success_redirect"`
	FailureRedirect        string                        `json:"failure_redirect"`
	CancelRedirect         string                        `json:"cancel_redirect"`
	SuccessCallback        *string                       `json:"success_callback,omitempty"`
	PaymentMethodWhitelist *[]string                     `json:"payment_method_whitelist,omitempty"`
}

type CreatePurchaseRequestClient struct {
	Email string `json:"email"`
}

type CreatePurchaseRequestPurchase struct {
	Products []Product `json:"products"`
}
