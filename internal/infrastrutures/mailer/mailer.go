package mailer

import (
	"log"

	"github.com/go-gomail/gomail"
	"github.com/kritpi/arom-web-services/configs"
)

//go:generate mockgen -source=./mailer.go -destination=./mocks/mock_mailer.go -package=mock_mailer

type Mailer interface {
	SendEmail(to string, subject string, body string) error
}

type mailerImpl struct {
	cfg    configs.Config
	dialer *gomail.Dialer
}

func NewMailer(host string, port int, username string, password string) Mailer {
	return &mailerImpl{
		dialer: gomail.NewDialer(host, port, username, password),
	}
}

// sendEmail implements Mailer.
func (m *mailerImpl) SendEmail(to string, subject string, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.cfg.EMAIL_FROM)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	d := gomail.NewDialer(m.cfg.SMTP_HOST, m.cfg.SMTP_PORT, m.cfg.EMAIL_FROM, m.cfg.EMAIL_PASSWORD)

	// ส่งอีเมล
	err := d.DialAndSend(msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	return nil
}
