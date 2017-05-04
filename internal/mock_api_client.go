package internal

import (
	"net/url"

	"github.com/genesor/twiliolo/option"
)

// MockAPIClient is the mocked implementation of the APIClient interface.
type MockAPIClient struct {
	GetCall    int
	GetFn      func(string, []option.RequestOption) ([]byte, error)
	PostCall   int
	PostFn     func(string, []option.RequestOption, url.Values) ([]byte, error)
	DeleteCall int
	DeleteFn   func(string, []option.RequestOption) error
}

// Get mocked function.
func (c *MockAPIClient) Get(uri string, requestOptions []option.RequestOption) ([]byte, error) {
	c.GetCall++

	return c.GetFn(uri, requestOptions)
}

// Post mocked function.
func (c *MockAPIClient) Post(uri string, requestOptions []option.RequestOption, updates url.Values) ([]byte, error) {
	c.PostCall++

	return c.PostFn(uri, requestOptions, updates)
}

// Delete mocked function.
func (c *MockAPIClient) Delete(uri string, requestOptions []option.RequestOption) error {
	c.DeleteCall++

	return c.DeleteFn(uri, requestOptions)
}
