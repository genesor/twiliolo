package twiliolo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/genesor/twiliolo"
	"github.com/genesor/twiliolo/internal"
	"github.com/genesor/twiliolo/option"
	"github.com/stretchr/testify/assert"
)

var (
	dateCreatedNumber = time.Now().Add(-5 * time.Hour)
	dateUpdatedNumber = time.Now().Add(-2 * time.Hour)
)

func TestIncomingPhoneNumberList(t *testing.T) {
	client := new(internal.MockAPIClient)
	client.GetFn = func(uri string, _ []option.RequestOption) ([]byte, error) {
		assert.Equal(t, "/IncomingPhoneNumbers.json", uri)

		response := fmt.Sprintf(`
		{
			"page": 0,
			"page_size": 50,
			"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json",
			"first_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=0&PageSize=50",
			"previous_page_uri": null,
			"next_page_uri": null,
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
			},{
				"sid": "TwiliololIncomingFake2",
				"account_sid": "TwilioloFake",
				"friendly_name": "Twiliolo Test Number Two",
				"phone_number": "+33987654321",
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
		}`, dateCreatedNumber.Format(time.RFC1123Z), dateUpdatedNumber.Format(time.RFC1123Z), dateCreatedNumber.Format(time.RFC1123Z), dateUpdatedNumber.Format(time.RFC1123Z))

		return []byte(response), nil
	}

	service := twiliolo.IncomingPhoneNumberService{Client: client}
	list, err := service.List()
	assert.NoError(t, err)
	assert.Equal(t, 0, list.Page)
	assert.Equal(t, 50, list.PageSize)
	assert.Equal(t, "", list.PreviousPageURI)
	assert.Equal(t, "", list.NextPageURI)
	assert.Equal(t, "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers.json?Page=0&PageSize=50", list.FirstPageURI)
	assert.Equal(t, "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers.json", list.URI)
	assert.Equal(t, "TwiliololIncomingFake", list.IncomingPhoneNumbers[0].Sid)
	assert.Equal(t, "TwiliololIncomingFake2", list.IncomingPhoneNumbers[1].Sid)
}

func TestIncomingPhoneNumberListNextPage(t *testing.T) {
	t.Run("OK - GET next page", func(t *testing.T) {
		client := new(internal.MockAPIClient)
		client.GetFn = func(uri string, requestOptions []option.RequestOption) ([]byte, error) {
			assert.Equal(t, "/IncomingPhoneNumbers.json", uri)
			assert.Equal(t, option.Page(1), requestOptions[0])
			assert.Equal(t, option.PageSize(94), requestOptions[1])

			response := fmt.Sprintf(`
			{
				"page": 1,
				"page_size": 94,
				"uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=1&PageSize=94",
				"first_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=0&PageSize=94",
				"previous_page_uri": "\/2010-04-01\/Accounts\/TwilioloFake\/IncomingPhoneNumbers.json?Page=0&PageSize=94",
				"next_page_uri": null,
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
		}

		service := twiliolo.IncomingPhoneNumberService{Client: client}
		previousList := twiliolo.IncomingPhoneNumberList{
			PageSize:    94,
			Page:        0,
			NextPageURI: "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers.json?Page=1&PageSize=94",
		}

		list, err := service.ListNextPage(&previousList)
		assert.NoError(t, err)
		assert.Equal(t, 1, list.Page)
		assert.Equal(t, 94, list.PageSize)
		assert.Equal(t, "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers.json?Page=0&PageSize=94", list.PreviousPageURI)
		assert.Equal(t, "", list.NextPageURI)
		assert.Equal(t, "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers.json?Page=0&PageSize=94", list.FirstPageURI)
		assert.Equal(t, "/2010-04-01/Accounts/TwilioloFake/IncomingPhoneNumbers.json?Page=1&PageSize=94", list.URI)
		assert.Equal(t, "TwiliololIncomingFake", list.IncomingPhoneNumbers[0].Sid)
	})
	t.Run("NOK - No next page", func(t *testing.T) {
		client := new(internal.MockAPIClient)

		service := twiliolo.IncomingPhoneNumberService{Client: client}
		previousList := twiliolo.IncomingPhoneNumberList{
			PageSize:    94,
			Page:        1,
			NextPageURI: "",
		}

		list, err := service.ListNextPage(&previousList)
		assert.Error(t, err)
		assert.Equal(t, twiliolo.ErrIncomingPhoneListNoNextPage, err)
		assert.Nil(t, list)
	})
}
