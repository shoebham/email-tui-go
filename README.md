# Simple SMTP Server

This is a SMTP server that uses net/smtp package to authenticate and send 
emails from a email id.


To use your email id, add a .env file and enter your authentication details 
like this 
```text
smtpHost=smtp.example.com
smtpPort=587
smtpUsername=<your-email-id>
smtpPassword=<your-email-password>
```

then run ```smtp.go``` and then open a terminal and run the following command
```bash
 telnet localhost 2525
```

After the above command, you will see a prompt like this
```text
220 Welcome to my awesome mail server
```
Then you can enter the following client commands  to send an email
```text
Client: HELO localhost
Server: 250 Hello

Client: MAIL FROM:<test@example.com>
Server: 250 OK

Client: RCPT TO:<recipient@example.com>
Server: 250 OK

Client: DATA
Server: 354 Enter message, ending with "." on a line by itself

Client: Subject: Test Email
Client: 
Client: This is a test email.
Client: .
Server: 250 OK Message accepted

Client: QUIT
Server: 221 Bye
```
