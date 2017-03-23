package twiliolo_test

import (
	"testing"

	"github.com/Genesor/twiliolo"
	"github.com/Genesor/twiliolo/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := twiliolo.NewClient("FAKE", "FAKE_FAKE", &internal.HTTPMockClient{})

	assert.IsType(t, &twiliolo.TwilioClient{}, client)
	assert.IsType(t, &twiliolo.IncomingPhoneNumberService{}, client.IncomingPhoneNumber)
}
