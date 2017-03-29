package mock

import "github.com/genesor/twiliolo"

// IncomingPhoneNumberService is the mock of a IncomingPhoneNumberService
type IncomingPhoneNumberService struct {
	GetFn            func(string, []twiliolo.RequestOption) (*twiliolo.IncomingPhoneNumber, error)
	GetCall          int
	UpdateFn         func(*twiliolo.IncomingPhoneNumber, []twiliolo.RequestOption) error
	UpdateCall       int
	AllFn            func() ([]*twiliolo.IncomingPhoneNumber, error)
	AllCall          int
	ListFn           func([]twiliolo.RequestOption) (*twiliolo.IncomingPhoneNumberList, error)
	ListCall         int
	ListNextPageFn   func(*twiliolo.IncomingPhoneNumberList, []twiliolo.RequestOption) (*twiliolo.IncomingPhoneNumberList, error)
	ListNextPageCall int
}

func (s *IncomingPhoneNumberService) Get(sid string, requestOptions ...twiliolo.RequestOption) (*twiliolo.IncomingPhoneNumber, error) {
	s.GetCall++

	return s.GetFn(sid, requestOptions)
}

func (s *IncomingPhoneNumberService) Update(incomingPhoneNumber *twiliolo.IncomingPhoneNumber, requestOptions ...twiliolo.RequestOption) error {
	s.UpdateCall++

	return s.UpdateFn(incomingPhoneNumber, requestOptions)
}

func (s *IncomingPhoneNumberService) All() ([]*twiliolo.IncomingPhoneNumber, error) {
	s.AllCall++

	return s.AllFn()
}

func (s *IncomingPhoneNumberService) List(requestOptions ...twiliolo.RequestOption) (*twiliolo.IncomingPhoneNumberList, error) {
	s.ListCall++

	return s.ListFn(requestOptions)
}

func (s *IncomingPhoneNumberService) ListNextPage(previousList *twiliolo.IncomingPhoneNumberList, requestOptions ...twiliolo.RequestOption) (*twiliolo.IncomingPhoneNumberList, error) {
	s.ListNextPageCall++

	return s.ListNextPageFn(previousList, requestOptions)
}
