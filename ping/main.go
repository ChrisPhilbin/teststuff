package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
	"net/smtp"
	"log"
)

func sendit() {
		auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
		to := []string{"recipient@example.net"}
		msg := []byte("To: recipient@example.net\r\n" +
			"Subject: Website down!\r\n" +
			"\r\n" +
			"Your website is down!\r\n")
		err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
		if err != nil {
			log.Fatal(err)
		}
 }

func main() {

	for{
		result, _ := exec.Command("ping", "google.com").Output()
		if strings.Contains(string(result), "Request timed out.") {
			fmt.Println("Error: Unable to connect")
			sendit()
		} else {
			fmt.Println("Success: Machine is connected")
		}
		time.Sleep(1 * time.Hour)
	}
}