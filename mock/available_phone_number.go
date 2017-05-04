package mock

import (
	"github.com/genesor/twiliolo"
	"github.com/genesor/twiliolo/option"
)

// AvailablePhoneNumberService is the mock of a AvailablePhoneNumberService
type AvailablePhoneNumberService struct {
	LocalFn   func(string, []option.RequestOption) ([]twiliolo.AvailablePhoneNumber, error)
	LocalCall int
	BuyFn     func(*twiliolo.AvailablePhoneNumber, []option.RequestOption) (*twiliolo.IncomingPhoneNumber, error)
	BuyCall   int
}

// Local mocked function.
func (s *AvailablePhoneNumberService) Local(country string, requestOptions ...option.RequestOption) ([]twiliolo.AvailablePhoneNumber, error) {
	s.LocalCall++

	return s.LocalFn(country, requestOptions)
}

// Buy mocked function.
func (s *AvailablePhoneNumberService) Buy(phone *twiliolo.AvailablePhoneNumber, requestOptions ...option.RequestOption) (*twiliolo.IncomingPhoneNumber, error) {
	s.BuyCall++

	return s.BuyFn(phone, requestOptions)
}
