package gopaypal

const (
	SandBoxURL          = "https://api.sandbox.paypal.com"
	LiveURL             = "https://api.paypal.com"
	IdentitySandBoxURL  = "https://www.sandbox.paypal.com"
	IdentityLiveURL     = "https://www.paypal.com"
	IdentityUserInfoURL = "/v1/identity/openidconnect/userinfo/?schema=openid"
	PaymentCreateURL    = "/v1/payments/payment"
	PaymentExecuteURL   = "/v1/payments/payment/%v/execute"
	PaymentInfoURL      = "/v1/payments/payment/%v"
	OAuthURL            = "/v1/oauth2/token"
	IdentityURL         = "/signin/authorize"
	IdentityTokenURL    = "/v1/identity/openidconnect/tokenservice"
	nonceLength         = 7
)
