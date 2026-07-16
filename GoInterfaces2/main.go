package main

import (
	"fmt"
)

type NotificationSender interface {
	Send(message string) error
}

type EmailSend struct{}

type SmsSender struct{}

func (emial *EmailSend) Send(message string) {
	fmt.Println(message)
}

func (sms *SmsSender) Send(message string) {
	fmt.Println(message)
}

type NotificationService struct {
	sender NotificationSender
}

func main() {

}
