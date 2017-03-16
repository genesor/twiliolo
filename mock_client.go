package twiliolo

import "net/url"

type MockClient struct {
	getCall    int
	getFn      func(url.Values, string) ([]byte, error)
	postCall   int
	postFn     func(url.Values, string) ([]byte, error)
	deleteCall int
	deleteFn   func(string) error
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

func (c *MockClient) get(params url.Values, uri string) ([]byte, error) {
	c.getCall++

	return c.getFn(params, uri)
}

func (c *MockClient) post(params url.Values, uri string) ([]byte, error) {
	c.postCall++

	return c.postFn(params, uri)
}
func (c *MockClient) delete(uri string) error {
	c.deleteCall++

	return c.deleteFn(uri)
}
