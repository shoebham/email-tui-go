package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
)

func TestSMTPServer(t *testing.T) {
	// Connect to the SMTP server
	conn, err := net.Dial("tcp", "localhost:2525")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	expectResponse(t, reader, "220") // Expect welcome message

	// Send HELO command
	sendCommand(t, conn, "HELO localhost")
	expectResponse(t, reader, "250") // Expect OK response

	// Send MAIL FROM command
	sendCommand(t, conn, "MAIL FROM:<test@example.com>")
	expectResponse(t, reader, "250") // Expect OK response

	// Send RCPT TO command
	sendCommand(t, conn, "RCPT TO:<recipient@example.com>")
	expectResponse(t, reader, "250") // Expect OK response

	// Send DATA command
	sendCommand(t, conn, "DATA")
	expectResponse(t, reader, "354") // Expect ready for data

	// Send email body and end with "."
	sendCommand(t, conn, "Subject: Test Email\r\n\r\nThis is a test email.\r\n.")
	expectResponse(t, reader, "250") // Expect message accepted

	// Send QUIT command
	sendCommand(t, conn, "QUIT")
	expectResponse(t, reader, "221") // Expect goodbye message
}

func sendCommand(t *testing.T, conn net.Conn, command string) {
	_, err := conn.Write([]byte(command + "\r\n"))
	if err != nil {
		t.Fatalf("Failed to send command %q: %v", command, err)
	}
}

func expectResponse(t *testing.T, reader *bufio.Reader, expectedCode string) {
	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}
	if !strings.HasPrefix(response, expectedCode) {
		t.Fatalf("Expected response code %q, got %q", expectedCode, response)
	}
}
