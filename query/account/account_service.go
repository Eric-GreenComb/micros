package main

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"time"

	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/query/account/service"
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

func (self *inmemService) GetAccountByUserId(user_id string) (r string) {
	if !bson.IsObjectIdHex(user_id) {
		return ""
	}
	var _obj bean.Account
	err := AccountCollection.Find(bson.M{"user_id": bson.ObjectIdHex(user_id)}).One(&_obj)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_obj)
	r = string(b)
	return
}

func (self *inmemService) GetBillingById(id string) (r string) {
	if !bson.IsObjectIdHex(id) {
		return ""
	}

	var _obj bean.Billing

	err := BillingCollection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&_obj)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_obj)
	r = string(b)
	return
}

func (self *inmemService) GetDealBillingByUserId(user_id string, timestamp int64, pagesize int64) (r string) {
	var _objs []bean.Billing

	query := bson.M{"createdtime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = global.BillingStatus_Deal

	err := BillingCollection.Find(query).Sort("-createdtime").Limit(int(pagesize)).All(&_objs)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_objs)
	r = string(b)
	return
}

func (self *inmemService) GetBillingByUserId(user_id string, timestamp int64, pagesize int64) (r string) {
	var _objs []bean.Billing

	query := bson.M{"createdtime": bson.M{"$lt": time.Unix(timestamp, 0)}}

	err := BillingCollection.Find(query).Sort("-createdtime").Limit(int(pagesize)).All(&_objs)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_objs)
	r = string(b)
	return
}
