package utils

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "TAGIHIN APP <nrmadi02@gmail.com>"
const CONFIG_AUTH_EMAIL = "admtagihin@gmail.com"
const CONFIG_AUTH_PASSWORD = "admintagihin123"

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(to string, subject string, data interface{}, templateFile string) error {
	var body bytes.Buffer

	tmp, err := ParseTemplateDir("templates")
	if err != nil {
		log.Fatal("Could not parse templates", err)
	}

	err = tmp.ExecuteTemplate(&body, templateFile, &data)
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body.String())

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
