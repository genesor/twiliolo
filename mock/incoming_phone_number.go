package mock

import (
	"github.com/genesor/twiliolo"
	"github.com/genesor/twiliolo/option"
)

// IncomingPhoneNumberService is the mock of a IncomingPhoneNumberService
type IncomingPhoneNumberService struct {
	GetFn            func(string, []option.RequestOption) (*twiliolo.IncomingPhoneNumber, error)
	GetCall          int
	UpdateFn         func(*twiliolo.IncomingPhoneNumber, []option.RequestOption) error
	UpdateCall       int
	AllFn            func() ([]*twiliolo.IncomingPhoneNumber, error)
	AllCall          int
	ListFn           func([]option.RequestOption) (*twiliolo.IncomingPhoneNumberList, error)
	ListCall         int
	ListNextPageFn   func(*twiliolo.IncomingPhoneNumberList, []option.RequestOption) (*twiliolo.IncomingPhoneNumberList, error)
	ListNextPageCall int
}

func (s *IncomingPhoneNumberService) Get(sid string, requestOptions ...option.RequestOption) (*twiliolo.IncomingPhoneNumber, error) {
	s.GetCall++

	return s.GetFn(sid, requestOptions)
}

func (s *IncomingPhoneNumberService) Update(incomingPhoneNumber *twiliolo.IncomingPhoneNumber, requestOptions ...option.RequestOption) error {
	s.UpdateCall++

	return s.UpdateFn(incomingPhoneNumber, requestOptions)
}

func (s *IncomingPhoneNumberService) All() ([]*twiliolo.IncomingPhoneNumber, error) {
	s.AllCall++

	return s.AllFn()
}

func (s *IncomingPhoneNumberService) List(requestOptions ...option.RequestOption) (*twiliolo.IncomingPhoneNumberList, error) {
	s.ListCall++

	return s.ListFn(requestOptions)
}

func (s *IncomingPhoneNumberService) ListNextPage(previousList *twiliolo.IncomingPhoneNumberList, requestOptions ...option.RequestOption) (*twiliolo.IncomingPhoneNumberList, error) {
	s.ListNextPageCall++

	return s.ListNextPageFn(previousList, requestOptions)
}
