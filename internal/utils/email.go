package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendOTPEmail(email string, otp string) error {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	to := []string{email}

	subject := "Your OTP for Neosurge Portfolio Management"
	body := fmt.Sprintf("Your OTP is: %s", otp)
	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", email, subject, body))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
