package service

import (
	"fmt"
	"net/smtp"
	"simple-order-stock-manager/model"
)

func (c *Config) sendEmail(sendEmailDetails model.SendEmailRequest) error {
	from := sendEmailDetails.From
	password := sendEmailDetails.Password
	to := []string{sendEmailDetails.To}
	message := []byte(sendEmailDetails.Template)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Failed to send email, error: ", err)
		return err
	}
	return nil
}
