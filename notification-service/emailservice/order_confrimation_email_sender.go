package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"notification-service/domain"
)

type OrderConfrimationEmailSender struct{}

func NewOrderConfrimationEmailSender() EmailService {
	return &OrderConfrimationEmailSender{}
}

func (o OrderConfrimationEmailSender) SendEmail(v interface{}) error {
	orderConfirmation := v.(domain.OrderConfirmation)
	funcMap := template.FuncMap{
		"mul": func(a, b float64) float64 {
			return a * b
		},
	}
	tmpl, err := template.New("order_confirmation.html").Funcs(funcMap).ParseFiles("template/order_confirmation.html")
	if err != nil {
		log.Println("Error parsing template:", err)
	}
	var body bytes.Buffer
	err = tmpl.Execute(&body, orderConfirmation)
	if err != nil {
		return err
	}
	msg := []byte("To: recipient@example.com\r\n" +
		"Subject: Order Confirmation email from ecommerce site\r\n" +
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
