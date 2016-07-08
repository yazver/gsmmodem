package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/yazver/gsmmodem"
	"github.com/yazver/gsmmodem/sms"
)

func main() {
	// Connect to modem and initialize
	dev := &gsmmodem.Device{
		CommandPort: "com6",
		NotifyPort:  "com7",
	}
	if err := dev.Open(); err != nil {
		log.Println(err)
		return
	}
	if err := dev.Init(gsmmodem.DeviceE173()); err != nil {
		log.Println(err)
		return
	}
	defer dev.Close()

	fmt.Print("Enter phone number: ")
	phoneNumber := ""
	if _, err := fmt.Scanln(&phoneNumber); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter message:")
	message := ""
	if _, err := fmt.Scanln(&message); err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("Send message \"%s\" to phone number %s\n", phoneNumber, message)
	err := dev.SendLongSMS(message, sms.PhoneNumber(phoneNumber))
	if err != nil {
		color.Set(color.FgHiRed)
		log.Printf("Ошибка отправки сообщения: %s\n", err.Error())
	} else {
		color.Set(color.FgHiGreen)
		log.Printf("Сообщение отправлено на номер: " + phoneNumber)
	}

}
