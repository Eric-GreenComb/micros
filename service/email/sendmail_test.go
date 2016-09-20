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

	_emailExtra := service.EmailExtra{}
	_emailExtra.Email.Host = "smtp.126.com:25"
	_emailExtra.Email.User = "ministor@126.com"
	_emailExtra.Email.Password = "xxxxxx"
	_emailExtra.Email.To = "ministor@126.com"
	_emailExtra.Email.Subject = "this is a tpl email"
	_emailExtra.Email.Mailtype = "html"

	_emailExtra.TempName = "hello"

	_mapParse := make(map[string]string)
	_mapParse["Hi"] = "Hello"
	_mapParse["Name"] = "Eric"
	_emailExtra.Parse = _mapParse

	b, err := json.Marshal(_emailExtra)
	if err != nil {
		t.Errorf("TestUnmarshal error")
	}

	nc.Publish("tpl", b)
	nc.Flush()

}
