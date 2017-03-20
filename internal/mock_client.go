package internal

import (
	"net/url"

	"github.com/Genesor/twiliolo"
)

type MockClient struct {
	GetCall    int
	GetFn      func(string, []twiliolo.RequestOption) ([]byte, error)
	PostCall   int
	PostFn     func(string, []twiliolo.RequestOption, url.Values) ([]byte, error)
	DeleteCall int
	DeleteFn   func(string, []twiliolo.RequestOption) error
}

func (c *MockClient) AccountSid() string {
	return "TwilioloFake"
}

func (c *MockClient) AuthToken() string {
	return "TwilioloFakeToken"
}

func (c *MockClient) RootURL() string {
	return "http://fake.sadoma.so/"
}

func (c *MockClient) Get(uri string, requestOptions []twiliolo.RequestOption) ([]byte, error) {
	c.GetCall++

	return c.GetFn(uri, requestOptions)
}

func (c *MockClient) Post(uri string, requestOptions []twiliolo.RequestOption, updates url.Values) ([]byte, error) {
	c.PostCall++

	return c.PostFn(uri, requestOptions, updates)
}
func (c *MockClient) Delete(uri string, requestOptions []twiliolo.RequestOption) error {
	c.DeleteCall++

	return c.DeleteFn(uri, requestOptions)
}
