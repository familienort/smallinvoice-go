package smallinvoice

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

// Credentials consists of client id and secret which is obtained
// in the Dashboard of smallinvoice.
type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

// AuthBundle stores all auth related data.
type AuthBundle struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

// Client stores all connection related information
// and acts as the gateway.
type Client struct {
	restyClient *resty.Client
	credentials Credentials
	authBundle  AuthBundle
}

// NewClient creates a new Smallinvoice client.
func NewClient(clientID string, clientSecret string) (*Client, error) {
	// Create the resty client
	resty := resty.New()

	// Configure resty.
	resty.SetHostURL("https://api.smallinvoice.com/v2")
	resty.SetHeader("Accept", "application/json")

	// Create basic client with resty and credentials.
	client := &Client{
		restyClient: resty,
		credentials: Credentials{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			GrantType:    "client_credentials",
		},
	}

	// Validate credentials by trying to fetch an auth bundle.
	resp, err := client.restyClient.R().
		SetBody(client.credentials).
		SetResult(client.authBundle).
		Post("/auth/access-tokens")

	// If it wasn't successful return the error.
	if err != nil {
		return &Client{}, err
	}

	// Check valid status code.
	if resp.StatusCode() != 200 {
		return &Client{}, errors.New(string(resp.Body()))
	}

	// Set the received access token.
	client.restyClient.SetAuthToken(client.authBundle.AccessToken)

	// If it was successful, return the client.
	return client, nil
}

// CreateContact creates a new contact
func (c *Client) CreateContact(contact Contact) (Contact, error) {
	resp, err := c.restyClient.R().
		SetBody(contact).
		SetResult(contact).
		Post("/contacts")

	if err != nil {
		return contact, err
	}

	if resp.StatusCode() != 201 {
		return contact, errors.New(string(resp.Body()))
	}

	return contact, nil
}
