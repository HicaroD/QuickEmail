package main

import (
	"reflect"
	"testing"
)

func TestGetFullServiceAddress(t *testing.T) {
	service_address := ServiceAddress{
		"stmp.example.com",
		"587",
	}

	full_service_address := service_address.get_full_service_address()

	if full_service_address != "stmp.example.com:587" {
		t.Errorf("Service address is not valid!")
	}
}

func TestHasMoreThanOneRecipient(t *testing.T) {
	recipients := "example@service.com;example2@service.com"

	if !has_more_than_one_recipient_email(recipients) {
		t.Errorf("It has more than recipient email")
	}
}

func TestExtractMultipleRecipients(t *testing.T) {
	recipients := "example@service.com;example2@service.com"

	extracted_recipients_email := extract_recipient_emails_from_argument(recipients)
	expected_output := []string{"example@service.com", "example2@service.com"}

	if !reflect.DeepEqual(extracted_recipients_email, expected_output) {
		t.Errorf("Invalid recipients. See: %v", extracted_recipients_email)
	}
}

func TestExtractSingleRecipient(t *testing.T) {
	recipients := "example@service.com"

	extracted_recipients_email := extract_recipient_emails_from_argument(recipients)
	expected_output := []string{"example@service.com"}

	if !reflect.DeepEqual(extracted_recipients_email, expected_output) {
		t.errorf("invalid recipient. see: %v", extracted_recipients_email)
	}
}
