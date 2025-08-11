package email

import (
	"fmt"
	"net/smtp"
	"os"
)

type Email struct {
	auth smtp.Auth
	addr string
	from string
}

func (e *Email) Send(to string, subject string, body string) error {
	msg := []byte("" +
		"From: " + e.from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" +
		body + "\r\n")
	err := smtp.SendMail(e.addr, e.auth, e.from, []string{to}, []byte(msg))
	if err != nil {
		fmt.Println("Error to send email", err)
	}
	return err
}

func New() *Email {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_ADDR"), os.Getenv("EMAIL_PASS"), os.Getenv("EMAIL_HOST"))
	return &Email{
		auth: auth,
		addr: os.Getenv("EMAIL_HOST") + ":587",
		from: os.Getenv("EMAIL_ADDR"),
	}
}
