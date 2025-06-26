package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"notification-service/domain"
)

type PaymentConfrimationEmailSender struct{}

func NewPaymentConfrimationEmailSender() EmailService {
	return &PaymentConfrimationEmailSender{}
}

func (o PaymentConfrimationEmailSender) SendEmail(v interface{}) error {
	paymentConfirmation := v.(domain.PaymentConfirmation)

	tmpl, err := template.New("payment_confirmation.html").ParseFiles("template/payment_confirmation.html")
	if err != nil {
		log.Println("Error parsing template:", err)
	}
	var body bytes.Buffer
	err = tmpl.Execute(&body, paymentConfirmation)
	if err != nil {
		return err
	}
	msg := []byte("To: recipient@example.com\r\n" +
		"Subject: Payment Confirmation email from ecommerce site\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		body.String())
	to := []string{"receipient@eample.com"}
	err = smtp.SendMail(
		"localhost:1025",
		nil,
		"noreply@gmail.com",
		to,
		msg,
	)
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Successfully sent email")
	}
	return nil

}
