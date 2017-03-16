package twiliolo

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Capabilites represents a Twilio Incoming Phone Number capabilities (MMS, SMS, Voice)
type Capabilites struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"SMS"`
	MMS   bool `json:"MMS"`
}

// IncomingPhoneNumber represents a Twilio Incoming Phone Number
type IncomingPhoneNumber struct {
	Sid                  string      `json:"sid"`
	AccountSid           string      `json:"account_sid"`
	FriendlyName         string      `json:"friendly_name"`
	PhoneNumber          string      `json:"phone_number"`
	VoiceURL             string      `json:"voice_url"`
	VoiceMethod          string      `json:"voice_method"`
	VoiceFallbackURL     string      `json:"voice_fallback_url"`
	VoiceFallbackMethod  string      `json:"voice_fallback_method"`
	StatusCallback       string      `json:"status_callback"`
	StatusCallbackMethod string      `json:"status_callback_method"`
	VoiceCallerIDLookup  bool        `json:"voice_caller_id_lookup"`
	VoiceApplicationSid  string      `json:"voice_application_sid"`
	DateCreated          string      `json:"date_created"`
	DateUpdated          string      `json:"date_updated"`
	SmsURL               string      `json:"sms_url"`
	SmsMethod            string      `json:"sms_method"`
	SmsFallbackURL       string      `json:"sms_fallback_url"`
	SmsFallbackMethod    string      `json:"sms_fallback_method"`
	SmsApplicationSid    string      `json:"sms_application_sid"`
	Capabilities         Capabilites `json:"capabilities"`
	Beta                 bool        `json:"beta"`
	APIVersion           string      `json:"api_version"`
	URI                  string      `json:"uri"`
}

// GetIncomingPhoneNumber performs a call to the twilio API to retrieve an Incoming Phone Number with its Sid
func GetIncomingPhoneNumber(client Client, sid string) (*IncomingPhoneNumber, error) {
	var incomingPhoneNumber *IncomingPhoneNumber

	res, err := client.get(url.Values{}, "/IncomingPhoneNumbers/"+sid+".json")
	if err != nil {
		return nil, err
	}

	incomingPhoneNumber = new(IncomingPhoneNumber)
	err = json.Unmarshal(res, incomingPhoneNumber)

	return incomingPhoneNumber, err
}

// UpdateIncomingPhoneNumber performs the update of the differents attributes of an Incoming Phone Number
func UpdateIncomingPhoneNumber(client Client, incomingPhoneNumber *IncomingPhoneNumber) error {
	if incomingPhoneNumber == nil || incomingPhoneNumber.Sid == "" {
		return ErrIncomingPhoneMissingData
	}

	updates := url.Values{}
	updates.Set("FriendlyName", incomingPhoneNumber.FriendlyName)
	updates.Set("ApiVersion", incomingPhoneNumber.APIVersion)
	updates.Set("VoiceUrl", incomingPhoneNumber.VoiceURL)
	updates.Set("VoiceMethod", incomingPhoneNumber.VoiceMethod)
	updates.Set("VoiceFallbackUrl", incomingPhoneNumber.VoiceFallbackURL)
	updates.Set("VoiceFallbackMethod", incomingPhoneNumber.VoiceFallbackMethod)
	updates.Set("VoiceCallerIdLookup", strconv.FormatBool(incomingPhoneNumber.VoiceCallerIDLookup))
	updates.Set("VoiceApplicationSid", incomingPhoneNumber.VoiceApplicationSid)
	updates.Set("StatusCallback", incomingPhoneNumber.StatusCallback)
	updates.Set("StatusCallbackMethod", incomingPhoneNumber.StatusCallbackMethod)
	updates.Set("SmsUrl", incomingPhoneNumber.SmsURL)
	updates.Set("SmsMethod", incomingPhoneNumber.SmsMethod)
	updates.Set("SmsFallbackUrl", incomingPhoneNumber.SmsFallbackURL)
	updates.Set("SmsFallbackMethod", incomingPhoneNumber.SmsFallbackMethod)
	updates.Set("AccountSid", incomingPhoneNumber.AccountSid)

	body, err := client.post(updates, "/IncomingPhoneNumbers/"+incomingPhoneNumber.Sid+".json")
	if err != nil {
		return err
	}

	var updatedIncomingPhoneNumber *IncomingPhoneNumber
	updatedIncomingPhoneNumber = new(IncomingPhoneNumber)

	err = json.Unmarshal(body, incomingPhoneNumber)
	if err != nil {
		return err
	}

	incomingPhoneNumber = updatedIncomingPhoneNumber

	return nil
}
