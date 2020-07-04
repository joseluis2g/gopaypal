package gopaypal

import (
	"testing"
)

func TestClient_CreatePayment(t *testing.T) {
	// Create gopaypal client
	client := NewClient(clientID, secret, SandBoxURL)

	// Try to get access token
	_, err := client.GetAccessToken()

	if err != nil {
		t.Errorf("Cannot get PayPal OAuth access token: %v", err)
		t.FailNow()
	}

	// Create payment
	res, err := client.CreatePayment(Payment{
		Intent: "sale",
		Payer: Payer{
			PaymentMethod: "paypal",
		},
		Transactions: []Transaction{
			Transaction{
				Amount: Amount{
					Total:    "10",
					Currency: "EUR",
					Details: Details{
						SubTotal: "10",
					},
				},
				Description: "gopaypal create payment test",
				Custom:      "gopaypal-test",
			},
		},
		RedirectURL: RedirectURL{
			ReturnURL: redirectURL,
			CancelURL: redirectURL,
		},
	})

	if err != nil {
		t.Errorf("Cannot create PayPal payment: %v", err)
		t.FailNow()
	}

	// Check if payment ID is valid
	if res.ID == "" {
		t.Error("Invalid PayPal create payment response", err)
		t.FailNow()
	}

	// Display payment link to user
	t.Logf("This is the PayPal approve payment URL %v", res.Links[1].Href)
	t.Log("Use the flags -paymentid and -payerid to test the execution of a PayPal payment")
}

func TestClient_ExecutePayment(t *testing.T) {
	// Check for necessary flags
	if paymentID == "" || payerID == "" {
		t.SkipNow()
	}

	// Create gopaypal client
	client := NewClient(clientID, secret, SandBoxURL)

	// Try to get access token
	_, err := client.GetAccessToken()

	if err != nil {
		t.Errorf("Cannot get PayPal OAuth access token: %v", err)
		t.FailNow()
	}

	// Try to execute the payment
	res, err := client.ExecutePayment(paymentID, payerID)

	if err != nil {
		t.Errorf("Cannot execute PayPal payment: %v", err)
		t.FailNow()
	}

	// Check if payment is approved
	if res.State != "approved" {
		t.Errorf("Payment is not executed. Got %v state. Expected %v", res.State, "approved")
		t.FailNow()
	}
}
