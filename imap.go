package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	username := os.Getenv("IMAP_USERNAME")
	password := os.Getenv("IMAP_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("IMAP credentials not set in environment variables")
	}

	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	// Login
	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	if mbox.Messages == 0 {
		log.Println("No messages")
		return
	}

	// Fetch the last 10 messages
	from := uint32(1)
	if mbox.Messages > 10 {
		from = mbox.Messages - 9
	}
	to := mbox.Messages

	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	go func() {
		err = c.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages)
		if err != nil {
			log.Fatal(err)
		}
		//close(messages)
	}()

	for msg := range messages {
		fmt.Printf("[Message Time: %s] Message Subject: %s\n",
			msg.Envelope.Date.Format("2006-01-02 15:04:05"),
			msg.Envelope.Subject)
	}
}
