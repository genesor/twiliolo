package twiliolo

import "net/url"

type service struct {
	Client APIClient
}

// APIClient is the interface for an HTTP Twilio Client
type APIClient interface {
	AccountSid() string
	AuthToken() string
	RootURL() string
	Get(string, []RequestOption) ([]byte, error)
	Post(string, []RequestOption, url.Values) ([]byte, error)
	Delete(string, []RequestOption) error
}

// TwilioClient is the struct containing all other services
type TwilioClient struct {
	common              service // Reuse a single struct instead of allocating one for each service on the heap.
	IncomingPhoneNumber IncomingPhoneNumberServiceInterface
}

// NewClient instanciates a new TwilioClient
func NewClient(accountSid string, authToken string, httpClient HTTPClient) *TwilioClient {
	c := TwilioClient{}
	c.common.Client = NewTwilioAPIClient(accountSid, authToken, httpClient)
	c.IncomingPhoneNumber = (*IncomingPhoneNumberService)(&c.common)

	return &c
}
