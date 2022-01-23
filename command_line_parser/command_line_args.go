package command_line_parser

import (
	"flag"
	"log"
)

func Parse_all_command_line_arguments() (string, string, string, string) {
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
