package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/banerwai/gommon/net/smtp"
	"text/template"

	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/etcd"
)

// Email email struct
type Email struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Mailtype string `json:"type"`
}

// EmailExtra email extra(tpl) struct
type EmailExtra struct {
	Email    Email             `json:"email"`
	TempName string            `json:"tempname"`
	Parse    map[string]string `json:"parse"`
}

// EmailService email service struct
type EmailService struct {
}

// SendEmail send email service
func (es *EmailService) SendEmail(json string) bool {
	var _email Email
	_err := es.Unmarshal(json, &_email)
	if _err != nil {
		return false
	}

	es.SendEmailBean(_email)

	return true
}

// SendTpl send email by tpl
func (es *EmailService) SendTpl(json string) bool {
	var _emailExtra EmailExtra
	_err := es.Unmarshal(json, &_emailExtra)
	if _err != nil {
		return false
	}

	var _email Email
	_email.Host = _emailExtra.Email.Host
	_email.User = _emailExtra.Email.User
	_email.Password = _emailExtra.Email.Password
	_email.To = _emailExtra.Email.To
	_email.Subject = _emailExtra.Email.Subject
	_email.Mailtype = _emailExtra.Email.Mailtype

	_email.Body = es.GenBodyByTpl(_emailExtra.TempName, _emailExtra.Parse)
	if len(_email.Body) == 0 {
		fmt.Println("gen body error")
		return false
	}

	_err = es.SendEmailBean(_email)
	if _err != nil {
		fmt.Println(_err.Error())
		return false
	}

	return true
}

// GenBodyByTpl gen email body by tpl
func (es *EmailService) GenBodyByTpl(tplname string, parse map[string]string) string {
	_tpl, _err := es.getTplFromEtcd(tplname)
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

func (es *EmailService) getTplFromEtcd(tplname string) (string, error) {
	_key := constant.EtcdKeyTplEmail + tplname
	_tpl, _err := etcd.GetValue(_key)
	if _err != nil {
		return "", _err
	}
	return _tpl, nil
}

// SendEmailBean send email bean
func (es *EmailService) SendEmailBean(_email Email) error {
	var _server smtp.Email
	_server.Server(_email.Host, _email.User, _email.Password)
	return _server.Send(_email.To, _email.Subject, _email.Body, _email.Mailtype)
}

// Unmarshal unmarshal bean
func (es *EmailService) Unmarshal(_json string, bean interface{}) error {
	err := json.Unmarshal([]byte(_json), &bean)
	if err != nil {
		return err
	}
	return nil
}
