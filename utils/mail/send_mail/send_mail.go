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
