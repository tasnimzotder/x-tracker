package services

import (
	"backend/utils"
	"fmt"
	"github.com/twilio/twilio-go"
	twilioAPI "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
)

func SendMessages(message string) {
	config, err := utils.LoadConfig("./../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.TwilioAccountSID,
		Password: config.TwilioAuthToken,
	})

	params := &twilioAPI.CreateMessageParams{}

	params.SetFrom(fmt.Sprintf("whatsapp:+%s", config.TwilioPhoneNumberFrom))
	params.SetTo(fmt.Sprintf("whatsapp:+%s", config.TwilioPhoneNumberTo))
	params.SetBody(message)

	rsp, err := twilioClient.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("error sending message: %s", err.Error())
	}

	if rsp.Body != nil {
		log.Println(*rsp.Body)
	} else {
		log.Println("no message sent")
	}
}
