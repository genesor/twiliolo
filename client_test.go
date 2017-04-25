package twiliolo_test

import (
	"testing"

	"github.com/genesor/twiliolo"
	"github.com/genesor/twiliolo/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := twiliolo.NewClient("FAKE", "FAKE_FAKE", &internal.HTTPMockClient{})

	assert.IsType(t, &twiliolo.TwilioClient{}, client)
	assert.IsType(t, &twiliolo.IncomingPhoneNumberService{}, client.IncomingPhoneNumber)
}
