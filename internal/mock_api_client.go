package internal

import (
	"net/url"

	"github.com/Genesor/twiliolo"
)

type MockAPIClient struct {
	GetCall    int
	GetFn      func(string, []twiliolo.RequestOption) ([]byte, error)
	PostCall   int
	PostFn     func(string, []twiliolo.RequestOption, url.Values) ([]byte, error)
	DeleteCall int
	DeleteFn   func(string, []twiliolo.RequestOption) error
}

func (c *MockAPIClient) AccountSid() string {
	return "TwilioloFake"
}

func (c *MockAPIClient) AuthToken() string {
	return "TwilioloFakeToken"
}

func (c *MockAPIClient) RootURL() string {
	return "http://fake.sadoma.so/"
}

func (c *MockAPIClient) Get(uri string, requestOptions []twiliolo.RequestOption) ([]byte, error) {
	c.GetCall++

	return c.GetFn(uri, requestOptions)
}

func (c *MockAPIClient) Post(uri string, requestOptions []twiliolo.RequestOption, updates url.Values) ([]byte, error) {
	c.PostCall++

	return c.PostFn(uri, requestOptions, updates)
}
func (c *MockAPIClient) Delete(uri string, requestOptions []twiliolo.RequestOption) error {
	c.DeleteCall++

	return c.DeleteFn(uri, requestOptions)
}
