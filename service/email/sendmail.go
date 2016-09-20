package main

import (
	"flag"
	"log"
	"os"

	"github.com/mikespook/golib/signal"

	"github.com/banerwai/micros/service/email/service"
	"github.com/nats-io/nats"
)

// SendEmail send email
func SendEmail(email string) {
	var _emailService service.EmailService
	_emailService.SendEmail(email)
}

// SendTpl send email by tpl
func SendTpl(emailextra string) {
	var _emailService service.EmailService
	_emailService.SendTpl(emailextra)
}

// SendEmailWorker send email nats worker
func SendEmailWorker(subject string, m *nats.Msg) {
	switch subject {
	case "mail":
		SendEmail(string(m.Data))
	case "tpl":
		SendTpl(string(m.Data))
	}
}

func usage() {
	log.Fatalf("Usage: sendmail [-u server] [-s mail/tpl] [-g queue-group]\n")
}

func main() {
	log.Println("Start SendEmail Worker in NATS QueueSubscribe...")
	defer log.Println("Shutdown complete!")

	var urls = flag.String("u", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var subject = flag.String("s", "tpl", "The mode subject for email send")
	var group = flag.String("g", "default", "The nats server queue-group")

	flag.Usage = usage
	flag.Parse()

	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	nc.QueueSubscribe(*subject, *group, func(msg *nats.Msg) {
		go SendEmailWorker(*subject, msg)
	})

	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()
}
