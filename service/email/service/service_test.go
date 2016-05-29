package service

import (
	"encoding/json"
	"testing"
)

func TestUnmarshal(t *testing.T) {

	_email_bean := Email{}
	// change email
	_email_bean.Host = "smtp.126.com:25"
	_email_bean.User = "xxx@126.com"
	_email_bean.Password = "xxx"
	_email_bean.To = "xxx@126.com"
	_email_bean.Subject = "Hi"
	_email_bean.Body = "This is a test email"
	_email_bean.Mailtype = "html"

	b, err := json.Marshal(_email_bean)
	if err != nil {
		t.Errorf("TestUnmarshal error")
	}

	var _email_out Email
	var _service EmailService
	_service.Unmarshal(string(b), &_email_out)
	if _email_out.Host != "smtp.126.com:25" {
		t.Errorf("Unmarshal error")
	}
}
