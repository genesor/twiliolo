package twiliolo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Genesor/twiliolo"
	"github.com/Genesor/twiliolo/internal"
	"github.com/stretchr/testify/assert"
)

var (
	dateCreatedNumber = time.Now().Add(-5 * time.Hour)
	dateUpdatedNumber = time.Now().Add(-2 * time.Hour)
)

func TestGetIncomingPhoneNumberList(t *testing.T) {
	client := new(internal.MockClient)
	client.GetFn = func(uri string, _ []twiliolo.RequestOption) ([]byte, error) {
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

	list, err := twiliolo.GetIncomingPhoneNumberList(client)
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
