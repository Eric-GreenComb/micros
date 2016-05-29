package main

import (
	"encoding/json"
	"flag"
	"github.com/banerwai/micros/service/email/service"
	"github.com/nats-io/nats"
	"testing"
)

func TestSendEmail(t *testing.T) {

	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	flag.Parse()

	nc, _ := nats.Connect(*urls)

	defer nc.Close()

	_email_extra := service.EmailExtra{}
	_email_extra.Email.Host = "smtp.126.com:25"
	_email_extra.Email.User = "ministor@126.com"
	_email_extra.Email.Password = "xxxxxx"
	_email_extra.Email.To = "ministor@126.com"
	_email_extra.Email.Subject = "this is a tpl email"
	_email_extra.Email.Mailtype = "html"

	_email_extra.TempName = "hello"

	_map_parse := make(map[string]string)
	_map_parse["Hi"] = "Hello"
	_map_parse["Name"] = "Eric"
	_email_extra.Parse = _map_parse

	b, err := json.Marshal(_email_extra)
	if err != nil {
		t.Errorf("TestUnmarshal error")
	}

	nc.Publish("tpl", b)
	nc.Flush()

}
