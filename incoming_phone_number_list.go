package twiliolo

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// IncomingPhoneNumberList represent the response of the Twilio API when calling /IncomingPhoneNumbers.json
type IncomingPhoneNumberList struct {
	Page                 int                   `json:"page"`
	PageSize             int                   `json:"page_size"`
	URI                  string                `json:"uri"`
	FirstPageURI         string                `json:"first_page_uri"`
	NextPageURI          string                `json:"next_page_uri"`
	PreviousPageURI      string                `json:"previous_page_uri"`
	IncomingPhoneNumbers []IncomingPhoneNumber `json:"incoming_phone_numbers"`
}

// GetIncomingPhoneNumberList retrieves the first page of all the Incoming Phone Number owned
func GetIncomingPhoneNumberList(client Client, requestOptions ...RequestOption) (*IncomingPhoneNumberList, error) {
	body, err := client.Get("/IncomingPhoneNumbers.json", requestOptions)
	if err != nil {
		return nil, err
	}

	var incomingPhoneNumberList *IncomingPhoneNumberList
	incomingPhoneNumberList = new(IncomingPhoneNumberList)
	err = json.Unmarshal(body, incomingPhoneNumberList)

	return incomingPhoneNumberList, err
}

// GetNextPageIncomingPhoneNumberList retrieves the next page of a given IncomingPhoneNumberList
// If an empty NextPageURI is present in the struct it'll return an error
func GetNextPageIncomingPhoneNumberList(client Client, previousList *IncomingPhoneNumberList, requestOptions ...RequestOption) (*IncomingPhoneNumberList, error) {
	if previousList == nil || previousList.NextPageURI == "" {
		return nil, ErrIncomingPhoneListNoNextPage
	}

	params := url.Values{}
	params.Set("Page", strconv.Itoa(previousList.Page+1))
	params.Set("PageSize", strconv.Itoa(previousList.PageSize))

	body, err := client.Get("/IncomingPhoneNumbers.json", requestOptions)

	if err != nil {
		return nil, err
	}

	var incomingPhoneNumberList *IncomingPhoneNumberList
	incomingPhoneNumberList = new(IncomingPhoneNumberList)
	err = json.Unmarshal(body, incomingPhoneNumberList)

	return incomingPhoneNumberList, err
}
