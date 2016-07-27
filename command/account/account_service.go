package main

import (
	"encoding/json"
	"errors"
	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/command/account/service"
	"labix.org/v2/mgo/bson"
	"time"
)

type inmemService struct {
}

func newInmemService() service.AccountService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) CreateAccount(json_account string) (r string) {
	var _account bean.Account
	err := json.Unmarshal([]byte(json_account), &_account)
	if err != nil {
		return err.Error()
	}
	_account.Id = bson.NewObjectId()

	_time := time.Now()

	_account.CreatedTime = _time

	_err := AccountCollection.Insert(_account)
	if _err != nil {
		return _err.Error()
	}
	return _account.Id.Hex()
}

func (self *inmemService) CreateBilling(json_billing string) (r string) {
	var _billing bean.Billing
	err := json.Unmarshal([]byte(json_billing), &_billing)
	if err != nil {
		return err.Error()
	}
	_billing.Id = bson.NewObjectId()

	_time := time.Now()

	_billing.CreatedTime = _time
	_billing.Status = global.BillingStatus_Create

	_err := BillingCollection.Insert(_billing)
	if _err != nil {
		return _err.Error()
	}
	return _billing.Id.Hex()
}

func (self *inmemService) DealBilling(billing_id string) (r string) {
	if !bson.IsObjectIdHex(billing_id) {
		return ""
	}

	_mongo_m := bson.M{}
	_mongo_m["status"] = global.BillingStatus_Deal

	_err := BillingCollection.Update(bson.M{"_id": bson.ObjectIdHex(billing_id)}, bson.M{"$set": _mongo_m})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) CancelBilling(billing_id string) (r string) {
	if !bson.IsObjectIdHex(billing_id) {
		return ""
	}

	_mongo_m := bson.M{}
	_mongo_m["status"] = global.BillingStatus_Cancel

	_err := BillingCollection.Update(bson.M{"_id": bson.ObjectIdHex(billing_id)}, bson.M{"$set": _mongo_m})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) GenAccount(user_id string) (r string) {
	if !bson.IsObjectIdHex(user_id) {
		return ""
	}
	var _billings []bean.Billing

	query := bson.M{}
	query["status"] = global.BillingStatus_Deal
	query["user_id"] = bson.ObjectIdHex(user_id)

	_err := BillingCollection.Find(query).All(&_billings)
	if _err != nil {
		return _err.Error()
	}

	_err = self.updateMultiCurrencyAccount(user_id, self.genMultiCurrency(_billings))

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) genMultiCurrency(billing []bean.Billing) []bean.MultiCurrencyAccount {
	var _lsMultiCurrencyAccount []bean.MultiCurrencyAccount

	var _beanMultiCurrencyAccount01 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount01.Currency = global.CURRENCY_USD
	_beanMultiCurrencyAccount01.Amount = self.genMultiCurrencyAmount(global.CURRENCY_USD, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount01)

	var _beanMultiCurrencyAccount02 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount02.Currency = global.CURRENCY_CNY
	_beanMultiCurrencyAccount02.Amount = self.genMultiCurrencyAmount(global.CURRENCY_CNY, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount02)

	var _beanMultiCurrencyAccount03 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount03.Currency = global.CURRENCY_EUR
	_beanMultiCurrencyAccount03.Amount = self.genMultiCurrencyAmount(global.CURRENCY_EUR, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount03)

	var _beanMultiCurrencyAccount04 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount04.Currency = global.CURRENCY_JPY
	_beanMultiCurrencyAccount04.Amount = self.genMultiCurrencyAmount(global.CURRENCY_JPY, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount04)

	var _beanMultiCurrencyAccount05 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount05.Currency = global.CURRENCY_GBP
	_beanMultiCurrencyAccount05.Amount = self.genMultiCurrencyAmount(global.CURRENCY_GBP, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount05)

	return _lsMultiCurrencyAccount
}

func (self *inmemService) genMultiCurrencyAmount(currency_code string, billings []bean.Billing) int64 {
	var ret int64
	ret = 0
	for _, _billing := range billings {
		if _billing.Currency == currency_code {
			ret += int64(_billing.Operate) * _billing.Amount
		}
	}
	return ret
}

func (self *inmemService) updateMultiCurrencyAccount(user_id string, multi_curreency []bean.MultiCurrencyAccount) error {
	if !bson.IsObjectIdHex(user_id) {
		return errors.New("user_id is not ObjectIdHex")
	}

	_, _err := AccountCollection.Upsert(bson.M{"user_id": bson.ObjectIdHex(user_id)}, bson.M{"$set": bson.M{"multi_curreency": multi_curreency}})
	return _err
}
