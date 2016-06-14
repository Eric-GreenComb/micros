package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/command/contact/service"
	"labix.org/v2/mgo/bson"
	"time"
)

type inmemService struct {
}

func newInmemService() service.ContactService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) CreateContact(json_contact string) (r string) {
	var _contact bean.Contact
	err := json.Unmarshal([]byte(json_contact), &_contact)
	if err != nil {
		return err.Error()
	}
	_contact.Id = bson.NewObjectId()

	_time := time.Now()

	_contact.CreatedTime = _time
	_contact.DealedTime = _time

	_contact.ClientSignup = false
	_contact.FreelancerSignup = false
	_contact.Dealed = false

	_err := ContactCollection.Insert(_contact)
	if _err != nil {
		return _err.Error()
	}
	return _contact.Id.Hex()
}

func (self *inmemService) ClientSignContact(contact_id string, status bool) (r string) {
	if !bson.IsObjectIdHex(contact_id) {
		return ""
	}
	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contact_id)}, bson.M{"$set": bson.M{"client_signup": status}})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) FreelancerSignContact(contact_id string, status bool) (r string) {
	if !bson.IsObjectIdHex(contact_id) {
		return ""
	}
	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contact_id)}, bson.M{"$set": bson.M{"freelancer_signup": status}})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) DealContact(contact_id string, status bool) (r string) {
	if !bson.IsObjectIdHex(contact_id) {
		return ""
	}
	_mongo_m := bson.M{}
	_mongo_m["dealed"] = true
	_mongo_m["dealedtime"] = time.Now()

	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contact_id)}, bson.M{"$set": _mongo_m})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) UpdateContact(contact_id string, mmap map[string]string) (r string) {
	if !bson.IsObjectIdHex(contact_id) {
		return ""
	}

	_mongo_m := bson.M{}

	for k, v := range mmap {
		_mongo_m[k] = v
	}

	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contact_id)}, bson.M{"$set": _mongo_m})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}
