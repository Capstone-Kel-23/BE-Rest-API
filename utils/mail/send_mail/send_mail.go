package send_mail

import (
	"github.com/Capstone-Kel-23/BE-Rest-API/utils"
)

func SendEmailVerification(to string, data interface{}) error {
	var err error
	template := "verificationEmail.html"
	subject := "Verification Email"
	err = utils.SendEmail(to, subject, data, template)
	if err != nil {
		return err
	}
	return nil
}

func SendPaymentEmail(to string, data interface{}) error {
	var err error
	template := "SendPayment.html"
	subject := "Invoice and Payment Email"
	err = utils.SendEmail(to, subject, data, template)
	if err != nil {
		return err
	}
	return nil
}

func SendPaymentCashEmail(to string, data interface{}) error {
	var err error
	template := "SendPaymentCash.html"
	subject := "Invoice and Payment Cash Email"
	err = utils.SendEmail(to, subject, data, template)
	if err != nil {
		return err
	}
	return nil
}
