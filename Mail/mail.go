package Mail

import (
	"fmt"
	"net/smtp"
	"os"
)

func Mail() {

	// Sender data.
	from := os.Getenv("from_email")
	password := os.Getenv("email_password")

	// Receiver email address.
	to := []string{
		"darirapunit@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	smtp.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
