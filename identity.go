package gopaypal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type IdentityAccessTokenResponse struct {
	baseURL      string `json:"-"`
	TokenType    string `json:"token_type"`
	Expires      string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type IdentityUserInfoResponse struct {
	UserID          string          `json:"user_id"`
	Sub             string          `json:"sub"`
	Name            string          `json:"name"`
	GivenName       string          `json:"given_name"`
	FamilyName      string          `json:"family_name"`
	MiddleName      string          `json:"middle_name"`
	Picture         string          `json:"picture"`
	Email           string          `json:"email"`
	EmailVerified   bool            `json:"email_verified"`
	Gender          string          `json:"gender"`
	BirthDate       string          `json:"birthdate"`
	ZoneInfo        string          `json:"zoneinfo"`
	Locale          string          `json:"locale"`
	PhoneNumber     string          `json:"phone_number"`
	AccountVerified bool            `json:"verified_account"`
	Address         IdentityAddress `json:"address"`
	AccountType     string          `json:"account_type"`
	AgeRange        string          `json:"age_range"`
	PayerID         string          `json:"payer_id"`
}

type IdentityAddress struct {
	StreetAddress string `json:"street_address"`
	Locality      string `json:"locality"`
	Region        string `json:"region"`
	PostalCode    string `json:"postal_code"`
	Country       string `json:"country"`
}

// GenerateIdentityURL creates and returns an identity URL used to log-in into the PayPal services
func (c Client) GenerateIdentityURL(state string, ret string, scope []string) (*url.URL, error) {
	// Change client base URL to meet identity one
	b := c.baseURL

	if c.baseURL == SandBoxURL {
		c.baseURL = IdentitySandBoxURL
	} else {
		c.baseURL = IdentityLiveURL
	}

	// Create base URL
	idenURL, err := url.Parse(c.baseURL + IdentityURL)

	if err != nil {
		return nil, err
	}

	c.baseURL = b

	// Get URL query
	query := idenURL.Query()

	// Set client ID
	query.Add("client_id", c.clientID)

	// Set response type
	query.Add("response_type", "token")

	// Set scope
	query.Add("scope", strings.Join(scope, "+"))

	// Set state
	query.Add("state", state)

	// Set nonce value
	query.Add("nonce", CreateNonce())

	// Set return URL
	query.Add("redirect_uri", ret)

	// Set URL query
	idenURL.RawQuery = query.Encode()

	return idenURL, nil
}

// GetTokenFromRefreshToken returns the access token from the given identity refresh token
func (c *Client) GetTokenFromRefreshToken(refresh string) (*IdentityAccessTokenResponse, error) {
	// Set grant_type
	buff := []byte(fmt.Sprintf(
		"grant_type=%v&refresh_token=%v",
		"refresh_token",
		refresh,
	))

	// Create new gopaypal basic request
	req, err := c.BasicRequest(IdentityTokenURL, buff, http.MethodPost)

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

	// Convert response to JSON object
	resp := IdentityAccessTokenResponse{
		baseURL: c.baseURL,
	}

	// Unmarshal response
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetTokenFromIdentityCode returns the access token from the given identity login code
func (c *Client) GetTokenFromIdentityCode(code, ret string) (*IdentityAccessTokenResponse, error) {
	// Set grant_type
	buff := []byte(fmt.Sprintf(
		"grant_type=%v&code=%v&redirect_uri=%v",
		"authorization_code",
		code,
		ret,
	))

	// Create new gopaypal basic request
	req, err := c.BasicRequest(IdentityTokenURL, buff, http.MethodPost)

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

	// Convert response to JSON object
	resp := IdentityAccessTokenResponse{
		baseURL: c.baseURL,
	}

	// Unmarshal response
	if err := json.Unmarshal(res, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetUserInfo gets user profile attributes by the given access token
func (c Client) GetUserInfo(tkn string) (*IdentityUserInfoResponse, error) {
	// Create new gopaypal basic request
	req, err := c.BasicRequest(IdentityUserInfoURL, nil, http.MethodPost)

	if err != nil {
		return nil, err
	}

	// Set content type header
	req.Header.Set("Content-Type", "application/json")

	// Set authorization header
	req.Header.Set("Authorization", "Bearer "+tkn)

	// Execute request
	res, err := c.Execute(req)

	if err != nil {
		return nil, err
	}

	resUserInfo := IdentityUserInfoResponse{}

	// Unmarshal response
	if err := json.Unmarshal(res, &resUserInfo); err != nil {
		return nil, err
	}

	return &resUserInfo, nil
}
