package service

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"

	gatherredis "github.com/banerwai/gather/common/redis"
	"github.com/banerwai/gommon/openapi/alidayu"
)

type SMS struct {
	RecNum          string `json:"rec_num"`
	SmsFreeSignName string `json:"sms_free_sign_name"`
	SmsTemplateCode string `json:"sms_template_code"`
	SmsParam        string `json:"sms_param"`
}

type SmsService struct {
}

func (self *SmsService) SendSms(json string) bool {
	var _sms SMS
	_err := self.Unmarshal(json, &_sms)
	if _err != nil {
		return false
	}

	self.goSendSms(_sms)

	return true
}

func (self *SmsService) goSendSms(sms SMS) {
	alidayu.SendSMS(sms.RecNum, sms.SmsFreeSignName, sms.SmsTemplateCode, sms.SmsParam)
}

func (self *SmsService) LPOP4Redis(key string) error {

	conn := gatherredis.RedisPool.Get()
	defer conn.Close()

	var _pop_err error
	var _sms string
	for _pop_err == nil {
		_sms, _pop_err = redis.String(conn.Do("LPOP", key))
		if _pop_err != nil {
			continue
		}
		self.SendSms(_sms)
	}

	return nil
}

func (self *SmsService) Unmarshal(_json string, bean interface{}) error {
	err := json.Unmarshal([]byte(_json), &bean)
	if err != nil {
		return err
	}
	return nil
}
