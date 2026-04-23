package chippg

type PaymentMethod string

const (
	PaymentMethodDuitnowQR      PaymentMethod = "duitnow_qr"
	PaymentMethodFPX            PaymentMethod = "fpx"
	PaymentMethodFPXB2B1        PaymentMethod = "fpx_b2b1"
	PaymentMethodCryptoCoin     PaymentMethod = "crypto_coin"
	PaymentMethodDNQR           PaymentMethod = "dnqr"
	PaymentMethodMaestro        PaymentMethod = "maestro"
	PaymentMethodMastercard     PaymentMethod = "mastercard"
	PaymentMethodMPGSApplePay   PaymentMethod = "mpgs_apple_pay"
	PaymentMethodMPGSGooglePay  PaymentMethod = "mpgs_google_pay"
	PaymentMethodRazer          PaymentMethod = "razer"
	PaymentMethodRazerAtome     PaymentMethod = "razer_atome"
	PaymentMethodRazerGrabpay   PaymentMethod = "razer_grabpay"
	PaymentMethodRazerMaybankQR PaymentMethod = "razer_maybankqr"
	PaymentMethodRazerShopeepay PaymentMethod = "razer_shopeepay"
	PaymentMethodRazerTNG       PaymentMethod = "razer_tng"
	PaymentMethodVisa           PaymentMethod = "visa"
)

type Status string

const (
	CREATED         Status = "created"
	SENT            Status = "sent"
	VIEWED          Status = "viewed"
	ERROR           Status = "error"
	CANCELLED       Status = "cancelled"
	OVERDUE         Status = "overdue"
	EXPIRED         Status = "expired"
	BLOCKED         Status = "blocked"
	HOLD            Status = "hold"
	RELEASED        Status = "released"
	PENDING_RELEASE Status = "pending_release"
	PENDING_CAPTURE Status = "pending_capture"
	PREAUTHORIZED   Status = "preauthorized"
	PAID            Status = "paid"
	PENDING_EXECUTE Status = "pending_execute"
	PENDING_CHARGE  Status = "pending_charge"
	CLEARED         Status = "cleared"
	SETTLED         Status = "settled"
	CHARGEBACK      Status = "chargeback"
	PENDING_REFUND  Status = "pending_refund"
	REFUNDED        Status = "refunded"
)
