package mailer

import (
	"log"

	"gopkg.in/gomail.v2"
	"github.com/kritpi/arom-web-services/configs"
)

//go:generate mockgen -source=./mailer.go -destination=./mocks/mock_mailer.go -package=mock_mailer

type Mailer interface {
	SendEmail(to string, subject string, body string) error
}

type mailerImpl struct {
	cfg    *configs.Config
	dialer *gomail.Dialer
	message *gomail.Message
}

func NewMailer(dialer *gomail.Dialer, message *gomail.Message, cfg *configs.Config) Mailer {
	return &mailerImpl{
		dialer: dialer,
		message: message,
		cfg: cfg,
}
}

// sendEmail implements Mailer.
func (m *mailerImpl) SendEmail(to string, subject string, body string) error {
	m.message.SetHeader("From", m.cfg.EMAIL_FROM)
	m.message.SetHeader("To", to)
	m.message.SetHeader("Subject", subject)
	m.message.SetBody("text/html", body)
	log.Println(m.cfg.EMAIL_FROM)
	log.Println(to)
	log.Println(subject)
	log.Println(body)

	err := m.dialer.DialAndSend(m.message)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	return nil
}
