package service

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"

	gatherredis "github.com/banerwai/gather/common/redis"
	"github.com/banerwai/gommon/openapi/alidayu"
)

type Sms struct {
	RecNum          string `json:"rec_num"`
	SmsFreeSignName string `json:"sms_free_sign_name"`
	SmsTemplateCode string `json:"sms_template_code"`
	SmsParam        string `json:"sms_param"`
}

type SmsService struct {
}

func (self *SmsService) SendSms(json string) bool {
	var _sms Sms
	_err := self.Unmarshal(json, &_sms)
	if _err != nil {
		return false
	}

	self.SendSmsBean(_sms)

	return true
}

// success, resp := alidayu.SendSMS("18888888888", "身份验证", "SMS_4000328", `{"code":"1234","product":"alidayu"}`)
func (self *SmsService) SendSmsBean(sms Sms) (success bool, response string) {
	return alidayu.SendSMS(sms.RecNum, sms.SmsFreeSignName, sms.SmsTemplateCode, sms.SmsParam)
}

func (self *SmsService) Unmarshal(_json string, bean interface{}) error {
	err := json.Unmarshal([]byte(_json), &bean)
	if err != nil {
		return err
	}
	return nil
}
