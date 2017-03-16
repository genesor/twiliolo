package twiliolo

import (
	"errors"
	"fmt"
)

var (
	// ErrIncomingPhoneListNoNextPage used when there is no next page in a list of incoming call while trying to retrieve the next page
	ErrIncomingPhoneListNoNextPage = errors.New("No NextPageURI available")
	// ErrTwilioServer used when Twilio throws a 500
	ErrTwilioServer = errors.New("Twilio Server Error")
	//ErrIncomingPhoneMissingData used when there is missing required data to perform in an IncomingPhoneNumber to perform an action
	ErrIncomingPhoneMissingData = errors.New("Missing required data in the IncomingPhoneNumber ")
)

// TwilioError is an error returned by the Twilio API
type TwilioError struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"more_info"`
}

func (e TwilioError) Error() string {
	var message string

	message = "Twilio Error, "

	if e.Status != 0 {
		message += fmt.Sprintf("Status: %d", e.Status)
	}

	if e.Code != 0 {
		message += fmt.Sprintf(", Code: %d", e.Code)
	}

	if e.Message != "" {
		message += ", Message: " + e.Message
	}

	return message
}
