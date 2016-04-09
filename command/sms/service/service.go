package service

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"

	gatherbean "github.com/banerwai/gather/command/bean"
	gatherredis "github.com/banerwai/gather/common/redis"
	"github.com/banerwai/gommon/net/sms"
)

type SmsService struct {
}

func (self *SmsService) SendSms(json string) bool {
	var _sms gatherbean.Sms
	_err := self.Unmarshal(json, &_sms)
	if _err != nil {
		return false
	}

	self.goSendSms(_sms)

	return true
}

func (self *SmsService) goSendSms(_sms gatherbean.Sms) {
	var _api_service sms.SmsApiBean
	_api_service.Server(_sms.Name, _sms.Pwd, _sms.Content, _sms.Mobile, _sms.Sign, _sms.Extno)
	_api_service.SendSms("http://web.cr6868.com/asmx/smsservice.aspx")
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
