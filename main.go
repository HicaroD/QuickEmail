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

const GMAIL_SMTP_PORT = "587"

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

func ask_for_user_email() (string, error) {
	var email string

	fmt.Print("Insert your e-mail: ")
	fmt.Scanln(&email)

	if email == "" {
		return "", fmt.Errorf("\nE-mail should not be empty!")
	}

	return email, nil
}

func ask_for_user_password() (string, error) {
	fmt.Print("Insert your password (INVISIBLE INPUT): ")
	password, err := terminal.ReadPassword(0)

	if string(password) == "" {
		return "", fmt.Errorf("Password shouldn't be empty")
	}

	if err != nil {
		log.Fatal(err)
	}

	return string(password), err
}

func extract_recipient_emails_from_argument(recipient_argument string) ([]string, error) {
	if recipient_argument == "" {
		return nil, fmt.Errorf("You should pass at least one recipient!")
	}
	recipients := strings.Split(recipient_argument, ";")
	return recipients, nil
}

func main() {
	username, topic, message_body, recipient := parse_all_command_line_arguments()
	recipients, err := extract_recipient_emails_from_argument(recipient)

	if err != nil {
		log.Fatal(err)
	}

	service_info := ServiceAddress{"smtp.gmail.com", GMAIL_SMTP_PORT}

	email, err := ask_for_user_email()

	if err != nil {
		log.Fatal(err)
	}

	password, err := ask_for_user_password()

	if err != nil {
		log.Fatal(err)
	}

	user := User{username, email}
	message := Message{topic, message_body}

	email_sender := EmailSender{
		service_address: service_info,
		user:            user,
		message:         message,
	}

	auth := email_sender.authenticate_host(strings.TrimSpace(password))

	err = email_sender.send_email(auth, recipients)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nE-mail was sent successfully!\n")
}
