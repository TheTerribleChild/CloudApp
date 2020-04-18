package smtputil

import (
	"fmt"
	"net/smtp"
	"regexp"
)
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")


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

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}