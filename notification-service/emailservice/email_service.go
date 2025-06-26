package service

type EmailService interface {
	SendEmail(v interface{}) error
}
