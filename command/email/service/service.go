package service

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"

	gatherredis "github.com/banerwai/gather/common/redis"
	"github.com/banerwai/gommon/net/smtp"
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

type EmailService struct {
}

func (self *EmailService) SendEmail(json string) bool {
	var _email Email
	_err := self.Unmarshal(json, &_email)
	if _err != nil {
		return false
	}

	self.goSendEmail(_email)

	return true
}

func (self *EmailService) goSendEmail(_email Email) {
	var _server smtp.Email
	_server.Server(_email.Host, _email.User, _email.Password)
	_server.Send(_email.To, _email.Subject, _email.Body, _email.Mailtype)
}

func (self *EmailService) LPOP4Redis(key string) error {

	conn := gatherredis.RedisPool.Get()
	defer conn.Close()

	var _pop_err error
	var _email string
	for _pop_err == nil {
		_email, _pop_err = redis.String(conn.Do("LPOP", key))
		if _pop_err != nil {
			continue
		}
		self.SendEmail(_email)
	}

	return nil
}

func (self *EmailService) Unmarshal(_json string, bean interface{}) error {
	err := json.Unmarshal([]byte(_json), &bean)
	if err != nil {
		return err
	}
	return nil
}
