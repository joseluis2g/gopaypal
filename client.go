package gopaypal

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/go-errors"
	"io/ioutil"
	"net/http"
	"time"
)

type PayPalError struct {
	Name             string `json:"name"`
	Message          string `json:"message"`
	ErrorDescription string `json:"error_description"`
}

// Client gopaypal client for communicating with the PayPal REST API endpoints
type Client struct {
	baseURL     string
	clientID    string
	secret      string
	AccessToken *oauthResponse
}

// NewClient creates and returns a new gopaypal client with the given credentials
func NewClient(clientID, secret, base string) Client {
	return Client{
		baseURL:     base,
		clientID:    clientID,
		secret:      secret,
		AccessToken: &oauthResponse{},
	}
}

// Execute runs the given HTTP request
func (c Client) Execute(req *http.Request) ([]byte, error) {
	// Create a new HTTP client
	client := http.Client{}

	// Execute request
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	// Close response body
	defer res.Body.Close()

	// Read the whole response body
	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// If invalid request parse error
	if res.StatusCode != 200 && res.StatusCode != 201 {

		e := PayPalError{}

		// Unmarshal error message
		if err := json.Unmarshal(b, &e); err != nil {
			return nil, err
		}

		return nil, errors.New(e.Message)
	}

	return b, nil
}

// BasicRequest creates a basic request to the PayPal endpoint without the Authorization header
func (c Client) BasicRequest(endpoint string, b []byte, method string) (*http.Request, error) {
	// Wrap byte array on a io.Reader
	buff := bytes.NewBuffer(b)

	// Create HTTP request
	req, err := http.NewRequest(method, c.baseURL+endpoint, buff)

	if err != nil {
		return nil, err
	}

	// Add accept header
	req.Header.Add("Accept", "application/json")

	// Add accept language header
	req.Header.Add("Accept-Language", "en_US")

	return req, nil
}

// AuthRequest creates a basic request to the PayPal endpoint with the Authorization header set
func (c *Client) AuthRequest(endpoint string, b []byte, method string) (*http.Request, error) {
	// Create basic request
	req, err := c.BasicRequest(endpoint, b, method)

	if err != nil {
		return nil, err
	}

	// Check if access token is expired
	if time.Now().After(c.AccessToken.Expires) {

		// Get new access token
		if _, err := c.GetAccessToken(); err != nil {
			return nil, err
		}
	}

	// Set authorization header
	req.Header.Set("Authorization", "Bearer "+c.AccessToken.AccessToken)

	return req, nil
}
