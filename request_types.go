package chippg

type createPurchaseRequest struct {
	Client                 ClientDetails    `json:"client"`
	Purchase               PurchaseDetails  `json:"purchase"`
	BrandID                string           `json:"brand_id"`
	SuccessRedirect        *string          `json:"success_redirect,omitempty"`
	FailureRedirect        *string          `json:"failure_redirect,omitempty"`
	CancelRedirect         *string          `json:"cancel_redirect,omitempty"`
	SuccessCallback        *string          `json:"success_callback,omitempty"`
	PaymentMethodWhitelist *[]PaymentMethod `json:"payment_method_whitelist,omitempty"`
}

type PurchaseOption struct {
	Client                 ClientDetails    `json:"client"`
	Purchase               PurchaseDetails  `json:"purchase"`
	SuccessRedirect        *string          `json:"success_redirect,omitempty"`
	FailureRedirect        *string          `json:"failure_redirect,omitempty"`
	CancelRedirect         *string          `json:"cancel_redirect,omitempty"`
	SuccessCallback        *string          `json:"success_callback,omitempty"`
	PaymentMethodWhitelist *[]PaymentMethod `json:"payment_method_whitelist,omitempty"`
}

type ClientDetails struct {
	Email    string  `json:"email"`
	FullName *string `json:"full_name,omitempty"`
}

type PurchaseDetails struct {
	Products []Product `json:"products"`
}
