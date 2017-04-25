package twiliolo_test

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/genesor/twiliolo"
	"github.com/genesor/twiliolo/internal"
	"github.com/stretchr/testify/assert"
)

var (
	dateCreated = time.Now().Add(-48 * time.Hour)
	dateUpdated = time.Now().Add(-24 * time.Hour)
	testNumber  = twiliolo.IncomingPhoneNumber{
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
		Capabilities:         twiliolo.Capabilites{MMS: false, SMS: false, Voice: true},
		APIVersion:           twiliolo.VERSION,
		URI:                  "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers/TwiliololIncomingFake.json",
	}
)

func TestIncomingPhoneNumberUpdate(t *testing.T) {
	t.Run("OK - Success update", func(t *testing.T) {
		phoneNumber := testNumber
		phoneNumber.FriendlyName = "New Friendly Name"
		newUpdated := time.Now().Format(time.RFC1123Z)

		client := new(internal.MockAPIClient)
		client.PostFn = func(uri string, _ []twiliolo.RequestOption, params url.Values) ([]byte, error) {
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
			}`, dateCreated.Format(time.RFC1123Z), newUpdated)

			return []byte(response), nil
		}

		service := twiliolo.IncomingPhoneNumberService{Client: client}

		err := service.Update(&phoneNumber)
		assert.NoError(t, err)
		assert.Equal(t, newUpdated, phoneNumber.DateUpdated)
	})

	t.Run("NOK - Missing ID", func(t *testing.T) {
		phoneNumber := twiliolo.IncomingPhoneNumber{
			FriendlyName: "I am invalid",
		}

		client := new(internal.MockAPIClient)
		service := twiliolo.IncomingPhoneNumberService{Client: client}

		err := service.Update(&phoneNumber)

		assert.Error(t, err)
		assert.Equal(t, twiliolo.ErrIncomingPhoneMissingData, err)
	})

	t.Run("NOK - Error on API call", func(t *testing.T) {
		client := new(internal.MockAPIClient)
		client.PostFn = func(uri string, _ []twiliolo.RequestOption, params url.Values) ([]byte, error) {
			assert.Equal(t, "/IncomingPhoneNumbers/TwiliololIncomingFake.json", uri)

			return nil, errors.New("Error in API")
		}

		service := twiliolo.IncomingPhoneNumberService{Client: client}
		err := service.Update(&testNumber)

		assert.Error(t, err)
		assert.EqualError(t, err, "Error in API")
	})
}

func TestIncomingPhoneNumberGet(t *testing.T) {
	client := new(internal.MockAPIClient)
	client.GetFn = func(uri string, _ []twiliolo.RequestOption) ([]byte, error) {
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

	service := twiliolo.IncomingPhoneNumberService{Client: client}
	number, err := service.Get("TwiliololIncomingFake")

	assert.NoError(t, err)
	assert.Equal(t, testNumber, *number)
}

func TestIncomingPhoneNumberAll(t *testing.T) {
	client := new(internal.MockAPIClient)
	client.GetFn = func(uri string, requestOptions []twiliolo.RequestOption) ([]byte, error) {
		assert.Equal(t, "/IncomingPhoneNumbers.json", uri)

		if len(requestOptions) == 1 {
			assert.Equal(t, twiliolo.OptionPageSize(200), requestOptions[0])

			response := fmt.Sprintf(`
			{
				"page": 0,
				"page_size": 200,
				"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?PageSize=200",
				"first_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=0&PageSize=200",
				"previous_page_uri": null,
				"next_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=1&PageSize=200",
				"incoming_phone_numbers": [{
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
				}]
			}`, dateCreatedNumber.Format(time.RFC1123Z), dateUpdatedNumber.Format(time.RFC1123Z))

			return []byte(response), nil
		} else if len(requestOptions) == 2 && requestOptions[0] == twiliolo.OptionPage(1) {
			assert.Equal(t, twiliolo.OptionPageSize(200), requestOptions[1])

			response := fmt.Sprintf(`
			{
				"page": 1,
				"page_size": 200,
				"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=1&PageSize=200",
				"first_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=0&PageSize=200",
				"previous_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=0&PageSize=200",
				"next_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=2&PageSize=200",
				"incoming_phone_numbers": [{
					"sid": "TwiliololIncomingFake2",
					"account_sid": "TwilioloFake",
					"friendly_name": "Twiliolo Test Number 2",
					"phone_number": "+33912345679",
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
					"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers\/TwiliololIncomingFake2.json"
				}]
			}`, dateCreatedNumber.Format(time.RFC1123Z), dateUpdatedNumber.Format(time.RFC1123Z))

			return []byte(response), nil
		} else if len(requestOptions) == 2 && requestOptions[0] == twiliolo.OptionPage(2) {
			assert.Equal(t, twiliolo.OptionPageSize(200), requestOptions[1])

			response := fmt.Sprintf(`
			{
				"page": 1,
				"page_size": 200,
				"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=2&PageSize=200",
				"first_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=0&PageSize=200",
				"previous_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=1&PageSize=200",
				"next_page_uri": null,
				"incoming_phone_numbers": [{
					"sid": "TwiliololIncomingFake3",
					"account_sid": "TwilioloFake",
					"friendly_name": "Twiliolo Test Number 3",
					"phone_number": "+33912345670",
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
					"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers\/TwiliololIncomingFake3.json"
				}]
			}`, dateCreatedNumber.Format(time.RFC1123Z), dateUpdatedNumber.Format(time.RFC1123Z))

			return []byte(response), nil
		}

		return nil, errors.New("Unknown call")

	}

	service := twiliolo.IncomingPhoneNumberService{Client: client}
	phones, err := service.All()

	assert.NoError(t, err)
	assert.Equal(t, 3, client.GetCall)
	assert.Equal(t, 3, len(phones))
	assert.Equal(t, "TwiliololIncomingFake3", phones[2].Sid)
}
