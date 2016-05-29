package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/banerwai/gommon/net/smtp"
	"text/template"

	"github.com/banerwai/global"
	"github.com/banerwai/gommon/etcd"
)

type Email struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Mailtype string `json:"type"`
}

type EmailExtra struct {
	Email    Email             `json:"email"`
	TempName string            `json:"tempname"`
	Parse    map[string]string `json:"parse"`
}

type EmailService struct {
}

func (self *EmailService) SendEmail(json string) bool {
	var _email Email
	_err := self.Unmarshal(json, &_email)
	if _err != nil {
		return false
	}

	self.SendEmailBean(_email)

	return true
}

func (self *EmailService) SendTpl(json string) bool {
	var _email_extra EmailExtra
	_err := self.Unmarshal(json, &_email_extra)
	if _err != nil {
		return false
	}

	var _email Email
	_email.Host = _email_extra.Email.Host
	_email.User = _email_extra.Email.User
	_email.Password = _email_extra.Email.Password
	_email.To = _email_extra.Email.To
	_email.Subject = _email_extra.Email.Subject
	_email.Mailtype = _email_extra.Email.Mailtype

	_email.Body = self.GenBodyByTpl(_email_extra.TempName, _email_extra.Parse)
	if len(_email.Body) == 0 {
		fmt.Println("gen body error")
		return false
	}

	_err = self.SendEmailBean(_email)
	if _err != nil {
		fmt.Println(_err.Error())
		return false
	}

	return true
}

func (self *EmailService) GenBodyByTpl(tplname string, parse map[string]string) string {
	_tpl, _err := self.getTplFromEtcd(tplname)
	if _err != nil || len(_tpl) == 0 {
		return ""
	}

	tpl, err := template.New(tplname).Parse(_tpl)
	if err != nil {
		return ""
	}

	b := bytes.NewBuffer(make([]byte, 0))

	err = tpl.Execute(b, parse)
	if err != nil {
		return ""
	}
	return b.String()
}

func (self *EmailService) getTplFromEtcd(tplname string) (string, error) {
	_key := global.ETCD_KEY_TPL_EMAIL + tplname
	fmt.Println(_key)
	_tpl, _err := etcd.GetValue(_key)
	if _err != nil {
		return "", _err
	}
	return _tpl, nil
}

func (self *EmailService) SendEmailBean(_email Email) error {
	var _server smtp.Email
	_server.Server(_email.Host, _email.User, _email.Password)
	return _server.Send(_email.To, _email.Subject, _email.Body, _email.Mailtype)
}

func (self *EmailService) Unmarshal(_json string, bean interface{}) error {
	err := json.Unmarshal([]byte(_json), &bean)
	if err != nil {
		return err
	}
	return nil
}
