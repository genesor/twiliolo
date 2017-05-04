package twiliolo

import (
	"encoding/json"
	"net/url"

	"github.com/genesor/twiliolo/option"
)

// AvailablePhoneNumberServiceInterface is the interface of a IncomingPhoneNumberService
type AvailablePhoneNumberServiceInterface interface {
	Local(string, ...option.RequestOption) ([]AvailablePhoneNumber, error)
	Buy(*AvailablePhoneNumber, ...option.RequestOption) (*IncomingPhoneNumber, error)
}

// AvailablePhoneNumberService handles communication with the Incoming Phone Number related methods.
type AvailablePhoneNumberService service

// AvailablePhoneNumber represents a Twilio Incoming Phone Number.
type AvailablePhoneNumber struct {
	FriendlyName        string      `json:"friendly_name"`
	PhoneNumber         string      `json:"phone_number"`
	ISOCountry          string      `json:"iso_country"`
	AddressRequirements string      `json:"address_requirements"`
	Capabilities        Capabilites `json:"capabilities"`
	Beta                bool        `json:"beta"`
	// US & CA Only
	Lata       string `json:"lata"`
	RateCenter string `json:"rate_center"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Region     string `json:"region"`
	PostalCode string `json:"postal_code"`
}

type searchAvailablePhoneNumber struct {
	AvailablePhoneNumbers []AvailablePhoneNumber `json:"available_phone_numbers"`
	URI                   string                 `json:"uri"`
}

// Local performs a call to the twilio API to retrieve Incoming Phone Numbers
// available with the given params
// Doc: https://www.twilio.com/docs/api/rest/available-phone-numbers#local-instance
func (s *AvailablePhoneNumberService) Local(countryCode string, requestOptions ...option.RequestOption) ([]AvailablePhoneNumber, error) {
	var search searchAvailablePhoneNumber

	res, err := s.Client.Get("/AvailablePhoneNumbers/"+countryCode+"/Local.json", requestOptions)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &search)

	return search.AvailablePhoneNumbers, err
}

// Buy performs the update of the differents attributes of an Incoming Phone Number.
// In case of a number with an address requirement you need to use the
// web UI to buy one first
// Doc: https://www.twilio.com/docs/api/rest/incoming-phone-numbers#list-post
func (s *AvailablePhoneNumberService) Buy(availablePhoneNumber *AvailablePhoneNumber, requestOptions ...option.RequestOption) (*IncomingPhoneNumber, error) {

	updates := url.Values{}
	updates.Set("PhoneNumber", availablePhoneNumber.PhoneNumber)
	updates.Set("FriendlyName", availablePhoneNumber.FriendlyName)

	body, err := s.Client.Post("/IncomingPhoneNumbers.json", requestOptions, updates)
	if err != nil {
		return nil, err
	}

	var incomingPhoneNumber IncomingPhoneNumber

	err = json.Unmarshal(body, &incomingPhoneNumber)
	if err != nil {
		return nil, err
	}

	return &incomingPhoneNumber, nil
}
