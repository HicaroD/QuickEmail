package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net/smtp"
	"strings"
)

type ServiceAddress struct {
	host string
	port string
}

func (service_address ServiceAddress) get_full_service_address() string {
	return service_address.host + ":" + service_address.port
}

type User struct {
	email string
}

type Message struct {
	subject      string
	message_body string
}

type EmailSender struct {
	service_address ServiceAddress
	user            User
	message         Message
}

func (email_sender EmailSender) get_email_message(recipient []string, message Message) []byte {
	msg := fmt.Sprintf("To: %s\r\n"+"Subject: %s\r\n"+"\r\n"+"%s\r\n",
		recipient,
		message.subject,
		message.message_body)

	return []byte(msg)
}

func (email_sender EmailSender) send_email(auth smtp.Auth, message Message, recipient []string) {
	err := smtp.SendMail(
		email_sender.service_address.get_full_service_address(),
		auth,
		email_sender.user.email,
		recipient,
		email_sender.get_email_message(recipient, message),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func (email_sender EmailSender) authenticate_host(password string) smtp.Auth {
	auth := smtp.PlainAuth("",
		email_sender.user.email,
		password,
		email_sender.service_address.host,
	)
	return auth
}

func parse_all_command_line_arguments() (*string, *string, *string) {
	topic := flag.String("topic", "", "The topic of the e-mail")
	message_body := flag.String("send", "", "The actual message that you want to send")
	recipient := flag.String("to", "", "The recipient e-mail")

	flag.Parse()

	return topic, message_body, recipient
}

func ask_for_user_email() string {
	var email string

	fmt.Print("Insert your e-mail: ")
	fmt.Scanln(&email)

	return email
}

func ask_for_user_password() string {
	fmt.Print("Insert your password (INVISIBLE INPUT): ")
	password, err := terminal.ReadPassword(0)

	if err != nil {
		log.Fatal(err)
	}

	return string(password)
}

func main() {
	gmail_address := ServiceAddress{"smtp.gmail.com", "587"}

	topic, message_body, recipient := parse_all_command_line_arguments()
	email, password := ask_for_user_email(), ask_for_user_password()

	user := User{email}
	message := Message{*topic, *message_body}

	email_sender := EmailSender{
		service_address: gmail_address,
		user:            user,
		message:         message,
	}

	auth := email_sender.authenticate_host(strings.TrimSpace(password))
	email_sender.send_email(auth, message, []string{*recipient})

	fmt.Printf("E-mail successfully sent to %s", *recipient)
}
