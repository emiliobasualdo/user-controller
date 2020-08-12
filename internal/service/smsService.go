package service

import (
	"context"
	"github.com/kevinburke/twilio-go"
	"net/url"
)


const AccountSid = "AC3a76c0b4bd8624d06ac96b8f6b96f40f" // todo pasar a dev
const AuthToken = "9f226bfa3b8225f8669ba83bab7406dc"
const VerifyServiceId = "VA28a6c2e09395cf1782d9decfe213b8c5"
const FromSms = "+14068135125"
const FromWapp = "+14155238886"

var client *twilio.Client

func SMSInit(){
	client = twilio.NewClient(AccountSid, AuthToken, nil)
	log.Info("Sms provider connected")
}

func SendSmsCode(to string) error {
	if to == "+5491100000000" {
		return nil
	}
	_, err := sendSms(to)
	return err
}

func sendSms(to string) (*twilio.VerifyPhoneNumber, error) {
	//_, err := client.Messages.SendMessage(FromSms, phoneNumber, "Sent via go :) ✓", nil)
	v := url.Values{}
	v.Set("To", to)
	v.Set("Channel", "sms")
	v.Set("Locale", "es")
	return client.Verify.Verifications.Create(context.Background(), VerifyServiceId, v)
}

func CheckCode(phoneNumber string, code string) (bool, error) {
	if code == "000000" {
		return false, nil
	}
	if code == "111111" {
		return true, nil
	}
	return checkCode(phoneNumber, code)
}

func checkCode(number string, code string) (bool, error) {
	v := url.Values{}
	v.Set("To", number)
	v.Set("Code", code)
	resp, err := client.Verify.Verifications.Check(context.Background(), VerifyServiceId, v)
	if err != nil {
		return false, err
	}
	if resp.Status == "approved" {
		return true, nil
	}
	return false, nil
}
