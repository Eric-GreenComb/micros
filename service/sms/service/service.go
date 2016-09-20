package service

import (
	"encoding/json"

	"github.com/banerwai/gommon/openapi/alidayu"
)

// Sms sms struct
type Sms struct {
	RecNum          string `json:"rec_num"`
	SmsFreeSignName string `json:"sms_free_sign_name"`
	SmsTemplateCode string `json:"sms_template_code"`
	SmsParam        string `json:"sms_param"`
}

// SmsService smsservice struct
type SmsService struct {
}

// SendSms smsservice sendsms
func (smsService *SmsService) SendSms(json string) bool {
	var _sms Sms
	_err := smsService.Unmarshal(json, &_sms)
	if _err != nil {
		return false
	}

	smsService.SendSmsBean(_sms)

	return true
}

// SendSmsBean send sms bean
// success, resp := alidayu.SendSMS("18888888888", "身份验证", "SMS_4000328", `{"code":"1234","product":"alidayu"}`)
func (smsService *SmsService) SendSmsBean(sms Sms) (success bool, response string) {
	return alidayu.SendSMS(sms.RecNum, sms.SmsFreeSignName, sms.SmsTemplateCode, sms.SmsParam)
}

// Unmarshal unmarshal
func (smsService *SmsService) Unmarshal(_json string, bean interface{}) error {
	err := json.Unmarshal([]byte(_json), &bean)
	if err != nil {
		return err
	}
	return nil
}
