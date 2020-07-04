package gopaypal

import (
	"encoding/json"
	"net/http"
	"time"
)

type oauthResponse struct {
	Scope       []string  `json:"scope"`
	Nonce       string    `json:"nonce"`
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	AppID       string    `json:"app_id"`
	ExpiresIn   int       `json"expires_in"`
	Expires     time.Time `json:"-"`
}

// GetAccessToken gets the OAuth2 token from the PayPal endpoint
func (c *Client) GetAccessToken() (*oauthResponse, error) {
	// Set grant_type
	buff := []byte("grant_type=client_credentials")

	// Create new gopaypal basic request
	req, err := c.BasicRequest(OAuthURL, buff, http.MethodPost)

	if err != nil {
		return nil, err
	}

	// Set basic HTTP authentication
	req.SetBasicAuth(c.clientID, c.secret)

	// Execute request
	res, err := c.Execute(req)

	if err != nil {
		return nil, err
	}

	// Parse response as JSON
	oauthres := oauthResponse{}

	// Unmarshal response
	if json.Unmarshal(res, &oauthres); err != nil {
		return nil, err
	}

	// Set client access token
	c.AccessToken = &oauthres

	// Set token expires in time
	c.AccessToken.Expires = time.Now().Add(time.Duration(c.AccessToken.ExpiresIn))

	return &oauthres, nil
}
