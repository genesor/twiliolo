package twiliolo

import (
	"encoding/json"
	"fmt"
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
	get(url.Values, string) ([]byte, error)
	post(url.Values, string) ([]byte, error)
	delete(string) error
}

// TwilioClient is the Twilio API client
type TwilioClient struct {
	accountSid string
	authToken  string
	rootURL    string
}

var _ Client = &TwilioClient{}

// NewClient instanciates a new TwilioClient
func NewClient(accountSid, authToken string) *TwilioClient {
	rootURL := ROOT + "/" + VERSION + "/Accounts/" + accountSid
	return &TwilioClient{accountSid, authToken, rootURL}
}

func (c *TwilioClient) post(values url.Values, uri string) ([]byte, error) {
	req, err := http.NewRequest("POST", c.buildURI(uri), strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.AccountSid(), c.AuthToken())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
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

func (c *TwilioClient) get(queryParams url.Values, uri string) ([]byte, error) {
	var params *strings.Reader

	if queryParams == nil {
		queryParams = url.Values{}
	}

	params = strings.NewReader(queryParams.Encode())
	req, err := http.NewRequest("GET", c.buildURI(uri), params)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.AccountSid(), c.AuthToken())
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
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

func (c *TwilioClient) delete(uri string) error {
	req, err := http.NewRequest("DELETE", c.buildURI(uri), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.AccountSid(), c.AuthToken())
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 204 {
		return fmt.Errorf("Non-204 returned from server for DELETE: %d", res.StatusCode)
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

func (c *TwilioClient) buildURI(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	newParts := make([]string, 0, len(parts))
	// Check for "http" because sometimes we get raw URLs from following the metadata.
	if !strings.HasPrefix(parts[0], "http") {
		newParts = append(newParts, c.RootURL())
	}
	for _, p := range parts {
		p = strings.Trim(p, "/")
		if p == "" {
			continue
		}
		newParts = append(newParts, p)
	}
	return strings.Join(newParts, "/")
}
