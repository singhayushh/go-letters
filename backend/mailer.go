package backend

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
)

// SendMail Function
func SendMail(body, title, senderName, senderMail, password, receiverName, receiverMail string) error {

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		senderMail,
		password,
		smtpServer,
	)

	from := mail.Address{Name: senderName, Address: senderMail}
	to := mail.Address{Name: receiverName, Address: receiverMail}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err := smtp.SendMail(
		smtpServer+":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	return err
}
