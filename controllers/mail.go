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
		GetSettings().SenderMail,
		GetSettings().SenderPassword,
		"smtp.gmail.com",
	)

	from := mail.Address{
		Name:    GetSettings().CompanyName + " - Jeux de Sophia",
		Address: GetSettings().SenderMail,
	}
	to := mail.Address{
		Name:    "",
		Address: toAddress,
	}
	title := "Inscription site " + GetSettings().CompanyName + " Jeux de Sophia"

	body := "Merci de votre inscription sur le site " + GetSettings().CompanyName + " des Jeux de Sophia\r\n" +
		"Votre nouveau mot de passe est " + password + "\r\n" +
		"\r\n" +
		"Note sur la protection des données personnelles :\r\n" +
		"Veuillez ne pas renseigner sur ce site des informations personnelles et/ou professionnelles portant atteinte à votre sécurité et/ou celle de votre entreprise\r\n" +
		"(identifiant, mot de passe, adresse du domicile, plaque d'immatriculation, carte de crédit, ...).\r\n"

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
		GetSettings().SenderMail,
		[]string{toAddress},
		[]byte(message),
	)

	return err
}
