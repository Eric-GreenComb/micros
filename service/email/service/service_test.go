package service

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T) {

	_emailBean := Email{}
	// change email
	_emailBean.Host = "smtp.126.com:25"
	_emailBean.User = "xxx@126.com"
	_emailBean.Password = "xxx"
	_emailBean.To = "xxx@126.com"
	_emailBean.Subject = "Hi"
	_emailBean.Body = "This is a test email"
	_emailBean.Mailtype = "html"

	b, err := json.Marshal(_emailBean)
	if err != nil {
		t.Errorf("TestUnmarshal error")
	}

	var _emailOut Email
	var _service EmailService
	_service.Unmarshal(string(b), &_emailOut)
	if _emailOut.Host != "smtp.126.com:25" {
		t.Errorf("Unmarshal error")
	}
}
