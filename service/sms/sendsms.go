package main

import (
	"flag"
	"log"
	"os"

	"github.com/mikespook/golib/signal"

	"github.com/banerwai/micros/service/sms/service"
	"github.com/nats-io/nats"
)

func SendSms(sms string) {
	var _sms_service service.SmsService
	_sms_service.SendSms(email)
}

func SendEmailWorker(subject string, m *nats.Msg) {
	switch subject {
	case "sms":
		SendSms(string(m.Data))
	}
}

func usage() {
	log.Fatalf("Usage: sendsms [-u server] [-s sms/tpl] [-g queue-group]\n")
}

func main() {
	log.Println("Start SendEmail Worker in NATS QueueSubscribe...")
	defer log.Println("Shutdown complete!")

	var urls = flag.String("u", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var subject = flag.String("s", "sms", "The mode subject for sms send")
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
