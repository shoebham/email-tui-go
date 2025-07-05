package main

import (
	auth2 "email-client/auth"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
	"os"
)

//func xoauth2(accessToken, email string) sasl.Client {
//	var encodedString string
//
//	body := "user=" + email + "^Aauth=Bearer " + accessToken
//	encodedString = base64.StdEncoding.EncodeToString([]byte(body))
//	return &Xoauth2Client{Identity: encodedString}
//}

func ConnectToImapWithOauth(accessToken, email string) (*client.Client, error) {
	connection, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Login
	err = connection.Authenticate(auth2.NewXoauth2Client(email, accessToken))
	if err != nil {
		log.Fatal(err)
	}
	return connection, err
}

func fetchEmails(connection *client.Client) ([]string, error) {
	// Connect to server
	// Select INBOX
	mbox, err := connection.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	if mbox.Messages == 0 {
		log.Println("No messages")
		return nil, nil
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
		err = connection.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages)
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
	return nil, nil
}

func LoginAndFetch(accessToken string) {
	username := os.Getenv("IMAP_USERNAME")
	password := os.Getenv("IMAP_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("IMAP credentials not set in environment variables")
	}

	connection, err := ConnectToImapWithOauth(accessToken, username)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Logout()
	fetchEmails(connection)

}

//func main() {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatalf("Error loading .env file: %s", err)
//	}
//	username := os.Getenv("IMAP_USERNAME")
//	password := os.Getenv("IMAP_PASSWORD")
//	if username == "" || password == "" {
//		log.Fatal("IMAP credentials not set in environment variables")
//	}
//
//	connection, err := ConnectToImapWithOauth(username, password)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer connection.Logout()
//	fetchEmails(connection)
//
//}
