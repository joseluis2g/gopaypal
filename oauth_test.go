package gopaypal

import (
	"flag"
	"testing"
)

var (
	clientID      string
	secret        string
	redirectURL   string
	identityToken string
	oauthToken    string
	paymentID     string
	payerID       string
)

func init() {
	// Get client ID string pointer
	flag.StringVar(&clientID, "clientid", "", "PayPal REST API client id")

	// Get secret key string pointer
	flag.StringVar(&secret, "secret", "", "PayPal REST API secret key")

	// Get redirect URL
	flag.StringVar(&redirectURL, "redirect", "", "PayPal REST API redirect uri")

	// Get identity token
	flag.StringVar(&identityToken, "iden", "", "PayPal Identity token")

	// Get OAuth token
	flag.StringVar(&oauthToken, "token", "", "PayPal OAuth token")

	// Get payment ID
	flag.StringVar(&paymentID, "paymentid", "", "PayPal payment ID")

	// Get payer ID
	flag.StringVar(&payerID, "payerid", "", "PayPal payer ID")

	flag.Parse()
}

func TestClient_GetAccessToken(t *testing.T) {
	if oauthToken != "" {
		t.SkipNow()
	}

	// Create gopaypal client
	client := NewClient(clientID, secret, SandBoxURL)

	// Try to get access token
	tokenres, err := client.GetAccessToken()

	if err != nil {
		t.Errorf("Cannot get PayPal OAuth access token: %v", err)
		t.FailNow()
	}

	// Check if response is valid
	if tokenres.AccessToken == "" {
		t.Error("Wrong access token")
		t.FailNow()
	}

	t.Logf("Your OAuth token is %v", tokenres.AccessToken)
}
