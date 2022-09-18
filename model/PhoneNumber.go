package model

import (
	"errors"
	"log"
	"regexp"
)

type PhoneNumber struct {
	Value string `json:"phoneNumber"`
}

func (phoneNumber *PhoneNumber) Validate() error {
	e164Pattern := `^\+[1-9]\d{1,14}$`
	match, err := regexp.Match(e164Pattern, []byte(phoneNumber.Value))
	if err != nil {
		log.Fatal(err.Error())
	}
	if !match {
		return errors.New("phone number must be in E.164 format")
	}
	return nil
}
