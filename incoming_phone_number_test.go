package twiliolo

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	dateCreated = time.Now().Add(-5 * time.Hour)
	dateUpdated = time.Now().Add(-2 * time.Hour)
	testNumber  = IncomingPhoneNumber{
		Sid:                  "TwiliololIncomingFake",
		AccountSid:           "TwilioloFake",
		FriendlyName:         "Twiliolo Test Number",
		PhoneNumber:          "+33912345678",
		VoiceURL:             "http://test.com",
		VoiceMethod:          "POST",
		VoiceFallbackURL:     "http://fail.com",
		VoiceFallbackMethod:  "GET",
		VoiceCallerIDLookup:  true,
		VoiceApplicationSid:  "",
		StatusCallback:       "http://status.com",
		StatusCallbackMethod: "GET",
		SmsURL:               "http://sms.com",
		SmsMethod:            "GET",
		SmsFallbackURL:       "http://fail-sms.com",
		SmsFallbackMethod:    "GET",
		DateCreated:          dateCreated.Format(time.RFC1123Z),
		DateUpdated:          dateUpdated.Format(time.RFC1123Z),
		Capabilities:         Capabilites{MMS: false, SMS: false, Voice: true},
		APIVersion:           VERSION,
		URI:                  "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers/TwiliololIncomingFake.json",
	}
)

func TestUpdatePhoneNumber(t *testing.T) {
	t.Run("OK - Success update", func(t *testing.T) {
		phoneNumber := testNumber
		phoneNumber.FriendlyName = "New Friendly Name"

		client := new(MockClient)
		client.postFn = func(params url.Values, uri string) ([]byte, error) {
			assert.Equal(t, "/IncomingPhoneNumbers/TwiliololIncomingFake.json", uri)
			assert.Equal(t, "New Friendly Name", params.Get("FriendlyName"))

			response := fmt.Sprintf(`
			{
				"sid": "TwiliololIncomingFake",
				"account_sid": "TwilioloFake",
				"friendly_name": "Twiliolo Test Number",
				"phone_number": "+33912345678",
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
			}`, dateCreated.Format(time.RFC1123Z), time.Now().Format(time.RFC1123Z))

			return []byte(response), nil
		}

		err := UpdateIncomingPhoneNumber(client, &phoneNumber)
		assert.NoError(t, err)
		assert.True(t, phoneNumber.DateUpdated > testNumber.DateUpdated)
	})

	t.Run("NOK - Missing ID", func(t *testing.T) {
		phoneNumber := IncomingPhoneNumber{
			FriendlyName: "I am invalid",
		}

		client := new(MockClient)
		err := UpdateIncomingPhoneNumber(client, &phoneNumber)

		assert.Error(t, err)
		assert.Equal(t, ErrIncomingPhoneMissingData, err)
	})

	t.Run("NOK - Error on API call", func(t *testing.T) {
		client := new(MockClient)
		client.postFn = func(params url.Values, uri string) ([]byte, error) {
			assert.Equal(t, "/IncomingPhoneNumbers/TwiliololIncomingFake.json", uri)

			return nil, errors.New("Error in API")
		}

		err := UpdateIncomingPhoneNumber(client, &testNumber)

		assert.Error(t, err)
		assert.EqualError(t, err, "Error in API")
	})
}

func TestGetIncomingPhoneNumber(t *testing.T) {
	client := new(MockClient)

	client.getFn = func(params url.Values, uri string) ([]byte, error) {
		assert.Equal(t, "/IncomingPhoneNumbers/TwiliololIncomingFake.json", uri)

		response := fmt.Sprintf(`
		{
			"sid": "TwiliololIncomingFake",
			"account_sid": "TwilioloFake",
			"friendly_name": "Twiliolo Test Number",
			"phone_number": "+33912345678",
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
		}`, dateCreated.Format(time.RFC1123Z), dateUpdated.Format(time.RFC1123Z))

		return []byte(response), nil
	}

	number, err := GetIncomingPhoneNumber(client, "TwiliololIncomingFake")

	assert.NoError(t, err)
	assert.Equal(t, testNumber, *number)
}
