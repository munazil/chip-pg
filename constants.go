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
