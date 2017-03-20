package twiliolo_test

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/Genesor/twiliolo"
	"github.com/Genesor/twiliolo/internal"
	"github.com/stretchr/testify/assert"
)

const ACCOUNT_SID = "FAKE"
const AUTH_TOKEN = "FAKE_FAKE"
const ROOT_URL = twiliolo.ROOT + "/" + twiliolo.VERSION + "/Accounts/" + ACCOUNT_SID

func TestGet(t *testing.T) {
	t.Run("Basic GET", func(t *testing.T) {
		httpMock := internal.HTTPMockClient{}
		httpMock.DoFn = func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, ROOT_URL+"/TestGet", req.URL.String())

			user, pass, ok := req.BasicAuth()
			assert.Equal(t, ACCOUNT_SID, user)
			assert.Equal(t, AUTH_TOKEN, pass)
			assert.Equal(t, true, ok)

			return &http.Response{
				Status:     strconv.Itoa(200),
				StatusCode: 200,
				Body:       internal.NewRespBodyFromString("Success"),
				Header:     http.Header{},
			}, nil
		}

		client := twiliolo.NewClient(ACCOUNT_SID, AUTH_TOKEN, &httpMock)
		body, err := client.Get("/TestGet", make([]twiliolo.RequestOption, 0))

		assert.NoError(t, err)
		assert.Equal(t, []byte("Success"), body)
	})

	t.Run("Query string GET", func(t *testing.T) {
		httpMock := internal.HTTPMockClient{}
		httpMock.DoFn = func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, ROOT_URL+"/TestGet?Page=94&PageSize=42", req.URL.String())

			user, pass, ok := req.BasicAuth()
			assert.Equal(t, ACCOUNT_SID, user)
			assert.Equal(t, AUTH_TOKEN, pass)
			assert.Equal(t, true, ok)

			return &http.Response{
				Status:     strconv.Itoa(200),
				StatusCode: 200,
				Body:       internal.NewRespBodyFromString("Success"),
				Header:     http.Header{},
			}, nil
		}
		queryParms := make([]twiliolo.RequestOption, 0)
		queryParms = append(queryParms, twiliolo.Page(94))
		queryParms = append(queryParms, twiliolo.PageSize(42))

		client := twiliolo.NewClient(ACCOUNT_SID, AUTH_TOKEN, &httpMock)
		body, err := client.Get("/TestGet", queryParms)

		assert.NoError(t, err)
		assert.Equal(t, []byte("Success"), body)
	})

	t.Run("Error performing GET", func(t *testing.T) {
		httpMock := internal.HTTPMockClient{}
		httpMock.DoFn = func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, ROOT_URL+"/TestGet", req.URL.String())

			user, pass, ok := req.BasicAuth()
			assert.Equal(t, ACCOUNT_SID, user)
			assert.Equal(t, AUTH_TOKEN, pass)
			assert.Equal(t, true, ok)

			return nil, errors.New("Error perfoming the request")
		}

		client := twiliolo.NewClient(ACCOUNT_SID, AUTH_TOKEN, &httpMock)
		body, err := client.Get("/TestGet", make([]twiliolo.RequestOption, 0))

		assert.Error(t, err)
		assert.Equal(t, errors.New("Error perfoming the request"), err)
		assert.Nil(t, body)
	})

	t.Run("Error 500 GET", func(t *testing.T) {
		httpMock := internal.HTTPMockClient{}
		httpMock.DoFn = func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, ROOT_URL+"/TestGet", req.URL.String())

			user, pass, ok := req.BasicAuth()
			assert.Equal(t, ACCOUNT_SID, user)
			assert.Equal(t, AUTH_TOKEN, pass)
			assert.Equal(t, true, ok)

			return &http.Response{
				Status:     strconv.Itoa(500),
				StatusCode: 500,
				Body:       internal.NewRespBodyFromString("BOOM"),
				Header:     http.Header{},
			}, nil
		}

		client := twiliolo.NewClient(ACCOUNT_SID, AUTH_TOKEN, &httpMock)
		body, err := client.Get("/TestGet", make([]twiliolo.RequestOption, 0))

		assert.Error(t, err)
		assert.Equal(t, twiliolo.ErrTwilioServer, err)
		assert.Equal(t, []byte("BOOM"), body)
	})

	t.Run("Error 403 GET", func(t *testing.T) {
		httpMock := internal.HTTPMockClient{}
		httpMock.DoFn = func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, ROOT_URL+"/TestGet", req.URL.String())

			user, pass, ok := req.BasicAuth()
			assert.Equal(t, ACCOUNT_SID, user)
			assert.Equal(t, AUTH_TOKEN, pass)
			assert.Equal(t, true, ok)

			return &http.Response{
				Status:     strconv.Itoa(403),
				StatusCode: 403,
				Header:     http.Header{},
				Body: internal.NewRespBodyFromString(`{
					"status": 42,
					"message": "Fake message",
					"code": 403,
					"more_info": "Fake error"
				}`),
			}, nil
		}

		client := twiliolo.NewClient(ACCOUNT_SID, AUTH_TOKEN, &httpMock)
		body, err := client.Get("/TestGet", make([]twiliolo.RequestOption, 0))

		assert.Error(t, err)

		twilioError, ok := err.(*twiliolo.TwilioError)

		assert.True(t, ok)
		assert.Equal(t, 403, twilioError.Code)
		assert.Equal(t, 42, twilioError.Status)
		assert.Equal(t, "Fake message", twilioError.Message)
		assert.Equal(t, "Fake error", twilioError.MoreInfo)
		assert.NotNil(t, body)
	})

	t.Run("Error 403 - Malformated JSON GET", func(t *testing.T) {
		httpMock := internal.HTTPMockClient{}
		httpMock.DoFn = func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, ROOT_URL+"/TestGet", req.URL.String())

			user, pass, ok := req.BasicAuth()
			assert.Equal(t, ACCOUNT_SID, user)
			assert.Equal(t, AUTH_TOKEN, pass)
			assert.Equal(t, true, ok)

			return &http.Response{
				Status:     strconv.Itoa(403),
				StatusCode: 403,
				Header:     http.Header{},
				Body: internal.NewRespBodyFromString(`{
					"status": "malformated JSON"
				}`),
			}, nil
		}

		client := twiliolo.NewClient(ACCOUNT_SID, AUTH_TOKEN, &httpMock)
		body, err := client.Get("/TestGet", make([]twiliolo.RequestOption, 0))

		assert.Error(t, err)

		_, ok := err.(*twiliolo.TwilioError)

		assert.Contains(t, err.Error(), "json: cannot unmarshal string")
		assert.False(t, ok)
		assert.NotNil(t, body)
	})
}
