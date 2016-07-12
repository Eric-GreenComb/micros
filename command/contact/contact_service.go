package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	bstrings "github.com/banerwai/gommon/strings"
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

	var _contact bean.Contact

	err := ContactCollection.Find(bson.M{"_id": bson.ObjectIdHex(contact_id)}).One(&_contact)

	if err != nil {
		return ""
	}

	_mmap := make(map[string]string)
	json.Unmarshal([]byte(_contact.TplParam), &_mmap)
	_mmap["ContactNumber"] = _contact.Id.Hex()

	_content := bstrings.ParseTpl("default", _contact.ContactTpl, _mmap)

	_mongo_m := bson.M{}
	_mongo_m["contact_content"] = _content
	_mongo_m["client_signup"] = true
	_mongo_m["freelancer_signup"] = true
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
