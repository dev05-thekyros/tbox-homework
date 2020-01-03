package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type twilioSMSService struct {
}
type SMSService interface {
	SendSMS(ctx context.Context, phoneNumber string, smsContent string) error
}

var AccountSid = "AC4de4a4b0f1d142ef55b7bafb634d57aa"
var AuthToken = "9552111f312b99d831b995624887e298"

func NewTwilloSMSService() *twilioSMSService {
	return &twilioSMSService{}
}

func (*twilioSMSService) SendSMS(ctx context.Context, phoneNumber string, smsContent string) error {

	// Set account keys & information
	sentPhoneNumber := "0931317941" // mine phone number :)
	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", AccountSid)

	// Init api twilio
	msgData := url.Values{}
	msgData.Set("To", phoneNumber)
	msgData.Set("From", sentPhoneNumber)
	msgData.Set("Body", smsContent)
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(AccountSid, AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	if resp.StatusCode >= StatusCodeOK && resp.StatusCode < StatusCOdeMultiChoice {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&data); err != nil {
			return err
		} else {
			return errors.New(fmt.Sprint(data["sid"]))
		}
	}
	return nil
}
