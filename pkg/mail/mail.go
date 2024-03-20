package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
)

type Mail struct {
	host string
	port string
}

func NewMail(host, port string) *Mail {
	return &Mail{host, port}
}

func (m Mail) SendMail(from, to, subject, message string) error {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("To: %s\r\n", to))
	buf.WriteString(fmt.Sprintf("From: %s\r\n", from))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	buf.WriteString("\r\n")
	buf.WriteString(message + "\r\n")

	return smtp.SendMail(
		fmt.Sprintf("%s:%s", m.host, m.port),
		nil,
		from,
		[]string{to},
		buf.Bytes(),
	)
}
