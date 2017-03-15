package controllers

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
)

//SendMail send registration mail with new password created
func SendMail(toAddress string, password string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"cgijds2017@gmail.com",
		"04mai2013",
		"smtp.gmail.com",
	)

	from := mail.Address{
		Name:    "CGI Jeux de Sophia",
		Address: "cgijds2017@gmail.com",
	}
	to := mail.Address{
		Name:    "",
		Address: toAddress,
	}
	title := "Inscription site CGI Jeux de Sophia"

	body := "Merci de votre inscription sur le site CGI des jeux de Sophia\r\nVotre nouveau mot de passe est " + password

	// Fill header data
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"cgijds2017@gmail.com",
		[]string{toAddress},
		[]byte(message),
	)

	return err
}
