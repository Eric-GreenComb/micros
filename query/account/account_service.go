package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"

	"github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/micros/query/account/service"
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

func (ims *inmemService) GetAccountByUserID(userID string) (r string) {
	if !bson.IsObjectIdHex(userID) {
		return ""
	}
	var _obj bean.Account
	err := AccountCollection.Find(bson.M{"user_id": bson.ObjectIdHex(userID)}).One(&_obj)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_obj)
	r = string(b)
	return
}

func (ims *inmemService) GetBillingByID(ID string) (r string) {
	if !bson.IsObjectIdHex(ID) {
		return ""
	}

	var _obj bean.Billing

	err := BillingCollection.Find(bson.M{"_id": bson.ObjectIdHex(ID)}).One(&_obj)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_obj)
	r = string(b)
	return
}

func (ims *inmemService) GetDealBillingByUserID(userID string, timestamp int64, pagesize int64) (r string) {
	var _objs []bean.Billing

	query := bson.M{"createdtime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = constant.BillingStatusDeal
	query["user_id"] = userID

	err := BillingCollection.Find(query).Sort("-createdtime").Limit(int(pagesize)).All(&_objs)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_objs)
	r = string(b)
	return
}

func (ims *inmemService) GetBillingByUserID(userID string, timestamp int64, pagesize int64) (r string) {
	var _objs []bean.Billing

	query := bson.M{"createdtime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["user_id"] = userID

	err := BillingCollection.Find(query).Sort("-createdtime").Limit(int(pagesize)).All(&_objs)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_objs)
	r = string(b)
	return
}
