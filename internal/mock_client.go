package internal

import (
	"io"
	"net/http"
	"net/url"
	"strings"

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

type HTTPMockClient struct {
	DoCall int
	DoFn   func(*http.Request) (*http.Response, error)
}

func (c *HTTPMockClient) Do(req *http.Request) (*http.Response, error) {
	c.DoCall++

	return c.DoFn(req)
}

// NewRespBodyFromString creates an io.ReadCloser from a string that is suitable for use as an
// http response body.
func NewRespBodyFromString(body string) io.ReadCloser {
	return &dummyReadCloser{strings.NewReader(body)}
}

type dummyReadCloser struct {
	body io.ReadSeeker
}

func (d *dummyReadCloser) Read(p []byte) (n int, err error) {
	n, err = d.body.Read(p)
	if err == io.EOF {
		d.body.Seek(0, 0)
	}
	return n, err
}

func (d *dummyReadCloser) Close() error {
	return nil
}
