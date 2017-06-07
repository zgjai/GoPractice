package main

import (
	"github.com/scorredoira/email"
	"log"
	"net/smtp"
	"net/mail"
)

func main() {
	// compose the message
	m := email.NewMessage("Hi", "this is the body")
	m.From = mail.Address{Name: "From", Address: "guijiang.zhang@qunar.com"}
	m.To = []string{"guijiang.zhang@qunar.com"}
	
	// add attachments
	if err := m.Attach("/Users/zhangguijiang/Downloads/1.png"); err != nil {
		log.Fatal(err)
	}
	
	// send it
	auth := smtp.PlainAuth("", "", "", "mta2.corp.qunar.com")
	if err := email.Send("mta2.corp.qunar.com:25", auth, m); err != nil {
		log.Fatal(err)
	}
}
