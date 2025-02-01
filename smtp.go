package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":2525")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	log.Println("Server started on port 2525")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	conn.Write([]byte("220 Welcome to my awesome mail server\n"))

	var from string
	var data string
	var to []string
	var authed bool

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from connection: ", err)
			break
		}
		line = strings.TrimSpace(line)
		log.Println("Received:", line)

		switch {
		case strings.HasPrefix(strings.ToUpper(line), "HELO"):
			conn.Write([]byte("250 Hello\n"))

		case strings.HasPrefix(strings.ToUpper(line), "MAIL FROM:"):
			if !authed {
				conn.Write([]byte("530 Authentication required\r\n"))
				continue
			}
			from = strings.TrimPrefix(line, "MAIL FROM:")
			to = []string{}
			conn.Write([]byte("250 OK\r\n"))

		case strings.HasPrefix(strings.ToUpper(line), "RCPT TO:"):
			if !authed {
				conn.Write([]byte("530 Authentication required\r\n"))
				continue
			}
			toEmail := strings.TrimPrefix(line, "RCPT TO:")
			to = append(to, toEmail)
			conn.Write([]byte("250 OK\r\n"))

		case strings.HasPrefix(strings.ToUpper(line), "DATA"):
			if !authed {
				conn.Write([]byte("530 Authentication required\r\n"))
				continue
			}
			conn.Write([]byte("354 Enter message, ending with \".\" on a line by itself\r\n"))

			var dataLines []string
			for {
				dataLine, err := reader.ReadString('\n')
				if err != nil {
					log.Println("Error reading mail body: ", err)
					break
				}
				trimmed := strings.TrimSpace(dataLine)
				if trimmed == "." {
					break
				}
				dataLines = append(dataLines, trimmed)
			}
			data = strings.Join(dataLines, "\n")
			log.Println("Received mail from:", from, "to:", to, "data:", data)
			conn.Write([]byte("250 OK Message accepted\r\n"))

			// Forward the email by performing MX lookups and sending to destination servers.
			err := forwardEmail(from, to, []byte(data))
			if err != nil {
				log.Printf("Error forwarding email: %v", err)
				fmt.Fprintf(conn, "550 Failed to forward mail\r\n")
			} else {
				fmt.Fprintf(conn, "250 OK: Message forwarded\r\n")
			}

		case strings.HasPrefix(strings.ToUpper(line), "QUIT"):
			fmt.Fprintf(conn, "221 Bye\r\n")
			return

		default:
			fmt.Fprintf(conn, "502 Command not implemented\r\n")
		}
	}
}

func forwardEmail(from string, to []string, subject string, data []byte) error {

	smtpHost := os.Getenv("smtpHost")
	smtpPort := 587
	smtpUsername := os.Getenv("smtpUsername")
	smtpPassword := os.Getenv("smtpPassword")

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	serverAddr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	formattedData := formatEmailData(from, to, "Forwarded Email", string(data))

	// Send the email
	err := smtp.SendMail(serverAddr, auth, from, to, formattedData)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func formatEmailData(from string, to []string, subject string, body string) []byte {
	headers := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=utf-8\r\n"+
		"\r\n"+
		"%s", from, strings.Join(to, ", "), subject, body)
	return []byte(headers)
}
