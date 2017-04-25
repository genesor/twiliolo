package twiliolo_test

import (
	"fmt"
	"testing"
	"time"

	"net/url"

	"github.com/genesor/twiliolo"
	"github.com/genesor/twiliolo/internal"
	"github.com/stretchr/testify/assert"
)

func TestAvailablePhoneNumberLocal(t *testing.T) {
	client := new(internal.MockAPIClient)
	client.GetFn = func(uri string, options []twiliolo.RequestOption) ([]byte, error) {
		assert.Equal(t, twiliolo.VoiceEnabled(true), options[0])
		assert.Equal(t, "/AvailablePhoneNumbers/FR/Local.json", uri)

		response := `
		{
			"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/AvailablePhoneNumbers\/FR\/Local.json?VoiceEnabled=true",
			"available_phone_numbers": [
				{
					"friendly_name": "+3366554433",
					"phone_number": "+3366554433",
					"iso_country": "FR",
					"capabilities": {
						"voice": true,
						"SMS": true,
						"MMS": false
					},
					"address_requirements": "none",
					"beta": false
				},
				{
					"friendly_name": "+3399887799",
					"phone_number": "+3399887799",
					"iso_country": "FR",
					"capabilities": {
						"voice": true,
						"SMS": false,
						"MMS": false
					},
					"beta": false
				}
			]
		}`

		return []byte(response), nil
	}

	service := twiliolo.AvailablePhoneNumberService{Client: client}
	list, err := service.Local("FR", twiliolo.VoiceEnabled(true))

	assert.NoError(t, err)
	assert.Equal(t, 1, client.GetCall)
	assert.Equal(t, 2, len(list))
	assert.Equal(t, "+3366554433", list[0].PhoneNumber)
	assert.Equal(t, "+3366554433", list[0].FriendlyName)
	assert.Equal(t, true, list[0].Capabilities.SMS)
	assert.Equal(t, true, list[0].Capabilities.Voice)
	assert.Equal(t, "FR", list[0].ISOCountry)
	assert.Equal(t, "+3399887799", list[1].PhoneNumber)
	assert.Equal(t, "+3399887799", list[1].FriendlyName)
	assert.Equal(t, false, list[1].Capabilities.SMS)
	assert.Equal(t, true, list[1].Capabilities.Voice)
	assert.Equal(t, "FR", list[1].ISOCountry)
}

func TestAvailablePhoneNumberBuy(t *testing.T) {
	client := new(internal.MockAPIClient)
	client.PostFn = func(uri string, _ []twiliolo.RequestOption, values url.Values) ([]byte, error) {
		assert.Equal(t, "+3399887799", values.Get("PhoneNumber"))
		assert.Equal(t, "FriendlyName", values.Get("FriendlyName"))
		assert.Equal(t, "/IncomingPhoneNumbers.json", uri)

		response := fmt.Sprintf(`
			{
				"sid": "TwiliololIncomingFake",
				"account_sid": "TwilioloFake",
				"friendly_name": "FriendlyName",
				"phone_number": "+3399887799",
				"voice_url": "http://test.com",
				"voice_method": "POST",
				"voice_fallback_url": "http://fail.com",
				"voice_fallback_method": "GET",
				"status_callback": "http://status.com",
				"status_callback_method": "GET",
				"voice_caller_id_lookup": true,
				"voice_application_sid": null,
				"date_created": "%s",
				"date_updated": "%s",
				"sms_url": "http://sms.com",
				"sms_method": "GET",
				"sms_fallback_url": "http://fail-sms.com",
				"sms_fallback_method": "GET",
				"sms_application_sid": null,
				"capabilities": {
					"voice": true,
					"sms": false,
					"mms": false
				},
				"beta": false,
				"api_version": "2010-04-01",
				"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers\/TwiliololIncomingFake.json"
			}`, dateCreated.Format(time.RFC1123Z), dateCreated.Format(time.RFC1123Z))

		return []byte(response), nil
	}

	availPhone := twiliolo.AvailablePhoneNumber{
		FriendlyName:        "FriendlyName",
		PhoneNumber:         "+3399887799",
		ISOCountry:          "FR",
		AddressRequirements: "",
		Beta:                false,
		Capabilities: twiliolo.Capabilites{
			SMS:   false,
			MMS:   false,
			Voice: true,
		},
	}

	service := twiliolo.AvailablePhoneNumberService{Client: client}
	phone, err := service.Buy(&availPhone)

	assert.NoError(t, err)
	assert.Equal(t, 1, client.PostCall)
	assert.Equal(t, "+3399887799", phone.PhoneNumber)
	assert.Equal(t, "TwiliololIncomingFake", phone.Sid)
	assert.Equal(t, false, phone.Capabilities.SMS)
	assert.Equal(t, true, phone.Capabilities.Voice)
	assert.Equal(t, "FriendlyName", phone.FriendlyName)
}
