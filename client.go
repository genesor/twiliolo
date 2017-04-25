package twiliolo

import (
	"net/url"

	"github.com/genesor/twiliolo/option"
)

type service struct {
	Client APIClient
}

// APIClient is the interface for an HTTP Twilio Client
type APIClient interface {
	Get(string, []option.RequestOption) ([]byte, error)
	Post(string, []option.RequestOption, url.Values) ([]byte, error)
	Delete(string, []option.RequestOption) error
}

// TwilioClient is the struct containing all other services
type TwilioClient struct {
	common               service // Reuse a single struct instead of allocating one for each service on the heap.
	IncomingPhoneNumber  IncomingPhoneNumberServiceInterface
	AvailablePhoneNumber AvailablePhoneNumberServiceInterface
}

// NewClient instanciates a new TwilioClient
func NewClient(accountSid string, authToken string, httpClient HTTPClient) *TwilioClient {
	c := TwilioClient{}
	c.common.Client = NewTwilioAPIClient(accountSid, authToken, httpClient)
	c.IncomingPhoneNumber = (*IncomingPhoneNumberService)(&c.common)
	c.AvailablePhoneNumber = (*AvailablePhoneNumberService)(&c.common)

	return &c
}
