package mock

import "github.com/genesor/twiliolo"

// NewMockClient instanciates a new TwilioClient with mocked services
func NewMockClient() *twiliolo.TwilioClient {
	c := twiliolo.TwilioClient{}
	c.IncomingPhoneNumber = &IncomingPhoneNumberService{}
	c.AvailablePhoneNumber = &AvailablePhoneNumberService{}

	return &c
}
