package internal

import (
	"io"
	"net/http"
	"strings"
)

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
