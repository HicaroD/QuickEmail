package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net/smtp"
	"strings"
)

func parse_all_command_line_arguments() (string, string, string, string) {
	username := flag.String("from", "", "Your username")
	topic := flag.String("topic", "", "The topic of the e-mail")
	message_body := flag.String("send", "", "The actual message that you want to send")
	recipient := flag.String("to", "", "The recipient's e-mail")

	flag.Parse()

	if *username == "" || *topic == "" || *message_body == "" || *recipient == "" {
		log.Fatalf("Every argument should be passed. Please, see 'Help' section on github.com/HicaroD/QuickEmail/#help")
	}

	return *username, *topic, *message_body, *recipient
}

const SMTP_PORT = "587"

type ServiceAddress struct {
	host string
	port string
}

func (service_address ServiceAddress) get_full_service_address() string {
	return service_address.host + ":" + service_address.port
}

type User struct {
	name  string
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

func (email_sender EmailSender) get_email_message(recipient []string) []byte {
	msg := fmt.Sprintf("From: %s <%s>\r\nTo: <%s>\r\n"+"Subject: %s\r\n"+"\r\n"+"%s\r\n",
		email_sender.user.name,
		email_sender.user.email,
		recipient,
		email_sender.message.subject,
		email_sender.message.message_body)

	return []byte(msg)
}

func (email_sender EmailSender) send_email(auth smtp.Auth, recipient []string) error {
	err := smtp.SendMail(
		email_sender.service_address.get_full_service_address(),
		auth,
		email_sender.user.email,
		recipient,
		email_sender.get_email_message(recipient),
	)

	return err
}

func (email_sender EmailSender) authenticate_host(password string) smtp.Auth {
	auth := smtp.PlainAuth("",
		email_sender.user.email,
		password,
		email_sender.service_address.host,
	)
	return auth
}

func ask_for_user_email() string {
	var email string

	fmt.Print("Insert your e-mail: ")
	fmt.Scanln(&email)

	return email
}

func ask_for_user_password() (string, error) {
	fmt.Print("Insert your password (INVISIBLE INPUT): ")
	password, err := terminal.ReadPassword(0)

	return string(password), err
}

func extract_recipient_emails_from_argument(command_line_argument_for_recipient string) []string {
    recipients := strings.Split(command_line_argument_for_recipient, ";")
	return recipients
}

func main() {
	username, topic, message_body, recipient := parse_all_command_line_arguments()
	recipients := extract_recipient_emails_from_argument(recipient)

	service_info := ServiceAddress{"smtp.gmail.com", SMTP_PORT}

	email := ask_for_user_email()
	password, password_err := ask_for_user_password()

	if password_err != nil {
		log.Fatal(password_err)
	}

	user := User{username, email}
	message := Message{topic, message_body}

	email_sender := EmailSender{
		service_address: service_info,
		user:            user,
		message:         message,
	}

	auth := email_sender.authenticate_host(strings.TrimSpace(password))

	email_send_err := email_sender.send_email(auth, recipients)

	if email_send_err != nil {
		log.Fatal(email_send_err)
	}

	fmt.Printf("\nE-mail was sent successfully!\n")
}
