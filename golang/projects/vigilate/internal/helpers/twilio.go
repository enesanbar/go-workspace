package helpers

import (
	"fmt"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/twilio/twilio-go"
)

func SendTextTwilio(to, msg string, prefs *Preferences) error {
	secret := prefs.GetPref("twilio_auth_token")
	key := prefs.GetPref("twilio_sid")

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username:   key,
		Password:   secret,
		AccountSid: key,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(prefs.GetPref("twilio_phone_number"))
	params.SetBody(msg)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}

	return nil
}
