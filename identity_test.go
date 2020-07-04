package gopaypal

import (
	"strings"
	"testing"
)

func TestClient_GenerateIdentityURL(t *testing.T) {
	// Create gopaypal client
	client := NewClient(clientID, secret, SandBoxURL)

	// Define URL scope
	scope := []string{"email"}

	// Generate identity URL
	url, err := client.GenerateIdentityURL("test", redirectURL, scope)

	if err != nil {
		t.Errorf("Cannot generate identity URL: %v", err)
		t.FailNow()
	}

	// Analyze the returned URL
	if url.Query().Get("client_id") != client.clientID {
		t.Error("Invalid URL client ID on identity URL")
		t.FailNow()
	}

	if url.Query().Get("scope") != strings.Join(scope, "+") {
		t.Error("Invalid URL scope")
		t.FailNow()
	}
}

func TestClient_GetTokenFromIdentityCode(t *testing.T) {
	// Create gopaypal client
	client := NewClient(clientID, secret, SandBoxURL)

	// Check if token is set
	if identityToken == "" {

		// Define URL scope
		scope := []string{"email"}

		// Generate identity URL
		url, err := client.GenerateIdentityURL("test", redirectURL, scope)

		if err != nil {
			t.Errorf("Cannot generate identity URL: %v", err)
			t.FailNow()
		}

		// Print URL to user for testing purposes
		t.Logf("Use this identity URL: %v", url.String())
		t.Log("Restart the test passing the identity token with the flag -iden")

		t.SkipNow()
	}

	// Get token from user code
	res, err := client.GetTokenFromIdentityCode(identityToken, redirectURL)

	if err != nil {
		t.Errorf("Cannot get token from identity code: %v", err)
		t.FailNow()
	}

	if res.AccessToken == "" {
		t.Error("Invalid identity access token")
		t.FailNow()
	}

	if res.TokenType != "Bearer" {
		t.Errorf("Unexpected access token type. Got %v expected %v", res.TokenType, "Bearer")
		t.FailNow()
	}

	// Test get user info call
	getUserInfoFromAccessToken(client, t)

	// Test refresh token call
	getTokenFromRefreshTokenTest(client, res.RefreshToken, t)

}

func getUserInfoFromAccessToken(client Client, t *testing.T) {
	// Get application access token first
	if _, err := client.GetAccessToken(); err != nil {
		t.Errorf("Cannot get access token: %v", err)
		t.FailNow()
	}

	// Get user information attributes
	info, err := client.GetUserInfo()

	if err != nil {
		t.Errorf("Cannot get user info attributes from access token: %v", err)
		t.FailNow()
	}

	// Check for valid user information attributes
	if info.UserID == "" {
		t.Error("Invalid user payer ID")
		t.FailNow()
	}
}

func getTokenFromRefreshTokenTest(client Client, refreshToken string, t *testing.T) {
	// Generate new access token by refresh token
	res, err := client.GetTokenFromRefreshToken(refreshToken)

	if err != nil {
		t.Errorf("Cannot get token from refresh token: %v", err)
		t.FailNow()
	}

	if res.AccessToken == "" {
		t.Error("Invalid identity access token")
		t.FailNow()
	}

	if res.TokenType != "Bearer" {
		t.Errorf("Unexpected access token type. Got %v expected %v", res.TokenType, "Bearer")
		t.FailNow()
	}
}
