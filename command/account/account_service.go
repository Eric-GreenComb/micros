package main

import (
	"encoding/json"
	"errors"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/micros/command/account/service"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type inmemService struct {
}

func newInmemService() service.AccountService {
	return &inmemService{}
}

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) CreateAccount(jsonAccount string) (r string) {
	var _account bean.Account
	err := json.Unmarshal([]byte(jsonAccount), &_account)
	if err != nil {
		return err.Error()
	}
	_account.ID = bson.NewObjectId()

	_account.CreatedTime = time.Now().UnixNano()

	_err := AccountCollection.Insert(_account)
	if _err != nil {
		return _err.Error()
	}
	return _account.ID.Hex()
}

func (ims *inmemService) CreateBilling(jsonBilling string) (r string) {
	var _billing bean.Billing
	err := json.Unmarshal([]byte(jsonBilling), &_billing)
	if err != nil {
		return err.Error()
	}
	_billing.ID = bson.NewObjectId()

	_billing.CreatedTime = time.Now().UnixNano()
	_billing.Status = constant.BillingStatusCreate

	_err := BillingCollection.Insert(_billing)
	if _err != nil {
		return _err.Error()
	}
	return _billing.ID.Hex()
}

func (ims *inmemService) DealBilling(billingID string) (r string) {
	if !bson.IsObjectIdHex(billingID) {
		return ""
	}

	_mongoM := bson.M{}
	_mongoM["status"] = constant.BillingStatusDeal

	_err := BillingCollection.Update(bson.M{"_id": bson.ObjectIdHex(billingID)}, bson.M{"$set": _mongoM})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) CancelBilling(billingID string) (r string) {
	if !bson.IsObjectIdHex(billingID) {
		return ""
	}

	_mongoM := bson.M{}
	_mongoM["status"] = constant.BillingStatusCancel

	_err := BillingCollection.Update(bson.M{"_id": bson.ObjectIdHex(billingID)}, bson.M{"$set": _mongoM})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) GenAccount(userID string) (r string) {
	if !bson.IsObjectIdHex(userID) {
		return ""
	}
	var _billings []bean.Billing

	query := bson.M{}
	query["status"] = constant.BillingStatusDeal
	query["user_id"] = bson.ObjectIdHex(userID)

	_err := BillingCollection.Find(query).All(&_billings)
	if _err != nil {
		return _err.Error()
	}

	_err = ims.updateMultiCurrencyAccount(userID, ims.genMultiCurrency(_billings))

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (ims *inmemService) genMultiCurrency(billing []bean.Billing) []bean.MultiCurrencyAccount {
	var _lsMultiCurrencyAccount []bean.MultiCurrencyAccount

	var _beanMultiCurrencyAccount01 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount01.Currency = constant.CurrencyUSD
	_beanMultiCurrencyAccount01.Amount = ims.genMultiCurrencyAmount(constant.CurrencyUSD, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount01)

	var _beanMultiCurrencyAccount02 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount02.Currency = constant.CurrencyCNY
	_beanMultiCurrencyAccount02.Amount = ims.genMultiCurrencyAmount(constant.CurrencyCNY, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount02)

	var _beanMultiCurrencyAccount03 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount03.Currency = constant.CurrencyEUR
	_beanMultiCurrencyAccount03.Amount = ims.genMultiCurrencyAmount(constant.CurrencyEUR, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount03)

	var _beanMultiCurrencyAccount04 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount04.Currency = constant.CurrencyJPY
	_beanMultiCurrencyAccount04.Amount = ims.genMultiCurrencyAmount(constant.CurrencyJPY, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount04)

	var _beanMultiCurrencyAccount05 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount05.Currency = constant.CurrencyGBP
	_beanMultiCurrencyAccount05.Amount = ims.genMultiCurrencyAmount(constant.CurrencyGBP, billing)
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount05)

	return _lsMultiCurrencyAccount
}

func (ims *inmemService) genMultiCurrencyAmount(currencyCode string, billings []bean.Billing) int64 {
	var ret int64
	ret = 0
	for _, _billing := range billings {
		if _billing.Currency == currencyCode {
			ret += int64(_billing.Operate) * _billing.Amount
		}
	}
	return ret
}

func (ims *inmemService) updateMultiCurrencyAccount(userID string, multiCurreency []bean.MultiCurrencyAccount) error {
	if !bson.IsObjectIdHex(userID) {
		return errors.New("user_id is not ObjectIdHex")
	}

	_, _err := AccountCollection.Upsert(bson.M{"user_id": bson.ObjectIdHex(userID)}, bson.M{"$set": bson.M{"multi_curreency": multiCurreency}})
	return _err
}
