package service

import (
	"context"
	"github.com/kevinburke/twilio-go"
	"github.com/spf13/viper"
	"massimple.com/wallet-controller/internal/models"
	"net/url"
	"os"
)


var accountSid string
var authToken string
var verifyServiceId string

var twilioClient *twilio.Client

func SMSInit(){
	log.Info("Connecting to Sms")
	accountSid 		= viper.GetString("smsProvider.accountSid")
	authToken 		= viper.GetString("smsProvider.authToken")
	verifyServiceId = viper.GetString("smsProvider.verifyServiceId")
	twilioClient = twilio.NewClient(accountSid, authToken, nil)
	log.Info("Sms provider connected successfully")
}

func SendSmsCode(to models.PhoneNumber) error {
	env, _ := os.LookupEnv("ENV")
	if env == "DEV" && to == "005491100000000" {
		return nil
	}
	_, err := sendSms(to.String())
	return err
}

func sendSms(to string) (*twilio.VerifyPhoneNumber, error) {
	//_, err := twilioClient.Messages.SendMessage(FromSms, phoneNumber, "Sent via go :) ✓", nil)
	v := url.Values{}
	v.Set("To", to)
	v.Set("Channel", "sms")
	v.Set("Locale", "es")
	return twilioClient.Verify.Verifications.Create(context.Background(), verifyServiceId, v)
}

func CheckCode(phoneNumber models.PhoneNumber, code string) (bool, error) {
	env, _ := os.LookupEnv("ENV")
	if env == "DEV" {
		if code == "0000" {
			return false, nil
		}
		if code == "1111" {
			return true, nil
		}
	}
	return checkCode(phoneNumber.String(), code)
}

func checkCode(number string, code string) (bool, error) {
	v := url.Values{}
	v.Set("To", number)
	v.Set("Code", code)
	resp, err := twilioClient.Verify.Verifications.Check(context.Background(), verifyServiceId, v)
	if err != nil {
		return false, err
	}
	if resp.Status == "approved" {
		return true, nil
	}
	return false, nil
}
