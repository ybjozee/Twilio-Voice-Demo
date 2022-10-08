package helper

import (
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/twilio/twilio-go/twiml"
	"math/rand"
	"os"
	"time"
)

func Call(phoneNumber string) (string, error) {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateCallParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetTwiml(playTwiml())

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		return "", err
	}

	response, _ := json.Marshal(*resp)
	return string(response), nil
}

func playTwiml() string {
	play := &twiml.VoicePlay{
		Url: randomAudioClip(),
	}

	verbList := []twiml.Element{play}
	playTwiml, err := twiml.Voice(verbList)
	if err != nil {
		fmt.Println(err)
	}
	return playTwiml
}

func randomAudioClip() string {
	rand.Seed(time.Now().Unix())
	audioUrls := []string{
		"https://bit.ly/3dmicn8",
		"https://bit.ly/3xxpyLI",
		"https://bit.ly/3QSGnHE",
	}
	n := rand.Int() % len(audioUrls)
	return audioUrls[n]
}
