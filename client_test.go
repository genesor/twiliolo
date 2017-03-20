package twiliolo_test

import (
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

	httpMock := internal.HTTPMockClient{}
	httpMock.DoFn = func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, ROOT_URL+"/TestGet", req.URL.String())

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
}
