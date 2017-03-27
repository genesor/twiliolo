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

// HTTPClient is the interface of an HTTP client making a request
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwilioAPIClient is the struct used to make the Twilio API calls
type TwilioAPIClient struct {
	AccountSid string
	AuthToken  string
	RootURL    string
	httpClient HTTPClient
}

var _ APIClient = &TwilioAPIClient{}

// NewTwilioAPIClient instanciates a new TwilioAPIClient
func NewTwilioAPIClient(accountSid, authToken string, httpClient HTTPClient) *TwilioAPIClient {
	rootURL := ROOT + "/" + VERSION + "/Accounts/" + accountSid
	return &TwilioAPIClient{accountSid, authToken, rootURL, httpClient}
}

func (c *TwilioAPIClient) Post(uri string, requestOptions []RequestOption, values url.Values) ([]byte, error) {
	uri, err := c.buildURL(uri, requestOptions)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.AccountSid, c.AuthToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

func (c *TwilioAPIClient) Get(uri string, requestOptions []RequestOption) ([]byte, error) {
	uri, err := c.buildURL(uri, requestOptions)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.AccountSid, c.AuthToken)

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

func (c *TwilioAPIClient) Delete(uri string, requestOptions []RequestOption) error {
	uri, err := c.buildURL(uri, requestOptions)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.AccountSid, c.AuthToken)

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

func (c *TwilioAPIClient) buildURL(uri string, requestOptions []RequestOption) (string, error) {
	uri = strings.Trim(uri, "/")
	if uri == "" {
		return "", errors.New("Empty URI")
	}

	var urlStr string
	// Check for "http" because sometimes we get raw URLs from following the metadata.
	if !strings.HasPrefix(uri, "http") {
		urlStr = c.RootURL + "/" + uri
	} else {
		urlStr = uri
	}

	u, err := url.Parse(urlStr)
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
