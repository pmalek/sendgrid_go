package mailer

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func helloEmail(subject, fromAddress, toAddress string) []byte {
	from := mail.NewEmail("from Patryk", fromAddress)
	to := mail.NewEmail("to Patryk", toAddress)

	content := mail.NewContent("text/plain", "some text here")
	m := mail.NewV3MailInit(from, subject, to, content)

	/*
	 *address = "test2@example.com"
	 *name = "Example User"
	 *email := mail.NewEmail(name, address)
	 *m.Personalizations[0].AddTos(email)
	 */

	return mail.GetRequestBody(m)
}

func SendHelloEmail(sendgrid_api_key, subject, fromAddress, toAddress string) {
	request := sendgrid.GetRequest(sendgrid_api_key, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	log.Printf("fromAddress: %q", fromAddress)
	log.Printf("toAddress: %q", toAddress)

	var body = helloEmail(subject, fromAddress, toAddress)
	log.Printf("body : %s", body)
	request.Body = body

	if response, err := sendgrid.API(request); err != nil {
		log.Printf("Error sending e-mail: %q", err)
	} else {
		log.Printf("status code: %v", response.StatusCode)
		log.Printf("body: %s", response.Body)
		log.Printf("headers: %s", response.Headers)
	}
}
