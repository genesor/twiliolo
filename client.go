package twiliolo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ROOT is the baseURL of the Twilio API
const ROOT = "https://api.twilio.com"

// VERSION is the version of the Twilio API
const VERSION = "2010-04-01"

// Client is the interface of a Twilio Client
type Client interface {
	AccountSid() string
	AuthToken() string
	RootURL() string
	Get(string, []RequestOption) ([]byte, error)
	Post(string, []RequestOption, url.Values) ([]byte, error)
	Delete(string, []RequestOption) error
}

// HTTPClient is the interface of an HTTP client making a request
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwilioClient is the Twilio API client
type TwilioClient struct {
	accountSid string
	authToken  string
	rootURL    string
	httpClient HTTPClient
}

var _ Client = &TwilioClient{}

// NewClient instanciates a new TwilioClient
func NewClient(accountSid, authToken string, httpClient HTTPClient) *TwilioClient {
	rootURL := ROOT + "/" + VERSION + "/Accounts/" + accountSid
	return &TwilioClient{accountSid, authToken, rootURL, httpClient}
}

func (c *TwilioClient) Post(uri string, requestOptions []RequestOption, values url.Values) ([]byte, error) {
	uri, err := c.buildURI(uri, requestOptions)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.AccountSid(), c.AuthToken())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return body, ErrTwilioServer
		}

		twilioError := new(TwilioError)
		json.Unmarshal(body, twilioError)
		return body, twilioError

	}

	return body, err
}

func (c *TwilioClient) Get(uri string, requestOptions []RequestOption) ([]byte, error) {
	uri, err := c.buildURI(uri, requestOptions)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.AccountSid(), c.AuthToken())

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return body, ErrTwilioServer
		}
		twilioError := new(TwilioError)
		err := json.Unmarshal(body, &twilioError)
		if err != nil {
			return body, err
		}

		return body, twilioError
	}

	return body, err
}

func (c *TwilioClient) Delete(uri string, requestOptions []RequestOption) error {
	uri, err := c.buildURI(uri, requestOptions)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.AccountSid(), c.AuthToken())

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		if res.StatusCode == 500 {
			return ErrTwilioServer
		}

		twilioError := new(TwilioError)
		err := json.Unmarshal(body, &twilioError)
		if err != nil {
			return err
		}

		return twilioError
	}

	return nil
}

// AccountSid returns the AccountSid of the client
func (c *TwilioClient) AccountSid() string {
	return c.accountSid
}

// AuthToken returns the AuthToken of the client
func (c *TwilioClient) AuthToken() string {
	return c.authToken
}

// RootURL returns the RootURL of the client
func (c *TwilioClient) RootURL() string {
	return c.rootURL
}

func (c *TwilioClient) buildURI(uri string, requestOptions []RequestOption) (string, error) {
	uri = strings.Trim(uri, "/")
	if uri == "" {
		return "", errors.New("Empty URI")
	}

	parts := make([]string, 0)

	// Check for "http" because sometimes we get raw URLs from following the metadata.
	if !strings.HasPrefix(uri, "http") {
		parts = append(parts, c.RootURL())
	}

	parts = append(parts, uri)

	u, err := url.Parse(strings.Join(parts, "/"))
	if err != nil {
		return "", err
	}

	q := u.Query()
	for _, option := range requestOptions {
		key, value := option.GetValue()
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
