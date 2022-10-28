package emails

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type Receiver struct {
	Name     string
	Surname  string
	Mail     []string
	Birthday string
}

func SendAll() error {
	// Sender data.
	from := "zhaikapan@gmail.com"
	password := "ndmjfsukuhirwjsg"

	// Receiver email address.
	to := []Receiver{
		{"Alfarabi", "Maken", []string{"s60180073@gse.cs.msu.ru"}, "14.05.2001"},
		//{"John", "Kash", []string{"fara2001@mail.ru"}, "12.05.1999"},
		//{"Sayan", "Kinzhitayev", []string{"s60180063@gse.cs.msu.ru"}, "19.04.2001"},
		//{"Madi", "Albek", []string{"albek.madi13@gmail.com"}, "13.05.2000"},
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template/email.html")
	var body bytes.Buffer

	//sending...
	for i := 0; i < len(to); i++ {
		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))
		t.Execute(&body, struct {
			Name    string
			Surname string
			Message string
		}{
			Name:    to[i].Name,
			Surname: to[i].Surname,
			Message: to[i].Birthday,
		})

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to[i].Mail, body.Bytes())
		if err != nil {
			return err
		}
		body = bytes.Buffer{}
	}
	return nil
}
