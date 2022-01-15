package command_line_parser

import (
    "strings"
    "flag"
)

func is_empty(flag_value string) bool {
	return strings.TrimSpace(flag_value) == ""
}

func Parse_all_command_line_arguments() (string, string, string, string) {
	username := flag.String("from", "", "Your username")
	topic := flag.String("topic", "", "The topic of the e-mail")
	message_body := flag.String("send", "", "The actual message that you want to send")
	recipient := flag.String("to", "", "The recipient's e-mail")

	flag.Parse()

	return username, topic, message_body, recipient
}
