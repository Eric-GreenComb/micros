package service

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/banerwai/gather/bean"
)

func TestUnmarshal(t *testing.T) {

	_email_bean := bean.Email{}
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

	var _email_out bean.Email
	var _service EmailService
	_service.Unmarshal(string(b), &_email_out)
	fmt.Println(_email_out.Host)
	fmt.Println(_email_out.User)
}

func TestLPOP4Redis(t *testing.T) {

	var _service EmailService
	_service.LPOP4Redis("banerwai:email:activeuser")
}
