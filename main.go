package main

import (
    "net/smtp"
    "fmt"
    "log"
    "strings"
    "flag"
)

type ServiceAddress struct {
    host string
    port string
}

func (service_address ServiceAddress) get_full_service_address() string {
    return service_address.host + ":" +  service_address.port
}


type User struct {
    email string
}


type Message struct {
    subject string
    message_body string
}


type EmailSender struct {
    service_address ServiceAddress
    user User
    message Message
}

func (email_sender EmailSender) get_email_message(recipient []string, message Message) []byte {
    msg := fmt.Sprintf("To: %s\r\n" + "Subject: %s\r\n" + "\r\n" + "%s\r\n", 
                      recipient,
                      message.subject,
                      message.message_body)

    return []byte(msg)
}

func (email_sender EmailSender) send_email(auth smtp.Auth, message Message, recipient []string) error {
    err := smtp.SendMail(
        email_sender.service_address.get_full_service_address(),
        auth,
        email_sender.user.email,
        recipient,
        email_sender.get_email_message(recipient, message),
    )

    return err;
}

func (email_sender EmailSender) authenticate_host(password string) smtp.Auth {
    auth := smtp.PlainAuth("",
                           email_sender.user.email,
                           password,
                           email_sender.service_address.host,
                       )
    return auth
}

func main(){
    email := flag.String("email", "", "Your email")
    topic := flag.String("topic", "", "The topic of the e-mail")
    message_body := flag.String("message", "", "The actual e-mail that you want to send")
    recipient := flag.String("recipient", "", "The recipient e-mail")

    flag.Parse()

    gmail_address := ServiceAddress {"smtp.gmail.com", "587"}
    user := User {*email}
    message := Message {*topic, *message_body}

    email_sender := EmailSender {
        service_address: gmail_address,
        user: user, 
        message: message,
    }

    var password string
    fmt.Print("Insert your password: ")
    fmt.Scanln(&password)

    auth := email_sender.authenticate_host(strings.TrimSpace(password))
    err := email_sender.send_email(auth, message, []string{*recipient})

    if err != nil {
        log.Fatal(err)
    }
}
