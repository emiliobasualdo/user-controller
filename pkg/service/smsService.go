package service

func SendSms(phoneNumber string) error {
	log.InfoF("Sending sms to %s", phoneNumber)
	return nil
}
