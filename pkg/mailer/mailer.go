package mailer

import ( 
	"net/smtp" 
)

type SMTPMailer struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func (m *SMTPMailer) Send(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)
	msg := []byte("Subject: " + subject + "\r\n" +
		"From: " + m.From + "\r\n" +
		"To: " + to + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body)

	return smtp.SendMail(m.Host+":"+m.Port, auth, m.From, []string{to}, msg)
} 