package smtputil

import (
	"fmt"
	"net/smtp"
)

type SMTPClient struct {
	Email    string
	Password string
	Host     string
	Port     int
}

func (instance *SMTPClient) SendEmail(toEmail string, subject string, body string) error {
	auth := smtp.PlainAuth("", instance.Email, instance.Password, instance.Host)

	msg := "From: " + instance.Email + "\n" +
		"To: " + toEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	return smtp.SendMail(fmt.Sprintf("%s:%d", instance.Host, instance.Port), auth, instance.Email, []string{toEmail}, []byte(msg))
}

func (instance *SMTPClient) SendEmailAsync(toEmail string, subject string, body string) {
	go instance.SendEmail(toEmail, subject, body)
}
