package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	bstrings "github.com/banerwai/gommon/strings"
	"github.com/banerwai/micros/command/contact/service"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type inmemService struct {
}

func newInmemService() service.ContactService {
	return &inmemService{}
}

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) CreateContact(jsonContact string) (r string) {
	var _contact bean.Contact
	err := json.Unmarshal([]byte(jsonContact), &_contact)
	if err != nil {
		return err.Error()
	}
	_contact.ID = bson.NewObjectId()

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
	return _contact.ID.Hex()
}

func (ims *inmemService) ClientSignContact(contactID string, status bool) (r string) {
	if !bson.IsObjectIdHex(contactID) {
		return ""
	}
	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contactID)}, bson.M{"$set": bson.M{"client_signup": status}})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) FreelancerSignContact(contactID string, status bool) (r string) {
	if !bson.IsObjectIdHex(contactID) {
		return ""
	}
	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contactID)}, bson.M{"$set": bson.M{"freelancer_signup": status}})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) DealContact(contactID string, status bool) (r string) {
	if !bson.IsObjectIdHex(contactID) {
		return ""
	}

	var _contact bean.Contact

	err := ContactCollection.Find(bson.M{"_id": bson.ObjectIdHex(contactID)}).One(&_contact)

	if err != nil {
		return ""
	}

	_mmap := make(map[string]string)
	json.Unmarshal([]byte(_contact.TplParam), &_mmap)
	_mmap["ContactNumber"] = _contact.ID.Hex()

	_content := bstrings.ParseTpl("default", _contact.ContactTpl, _mmap)

	_mongoM := bson.M{}
	_mongoM["contact_content"] = _content
	_mongoM["client_signup"] = true
	_mongoM["freelancer_signup"] = true
	_mongoM["dealed"] = true
	_mongoM["dealedtime"] = time.Now()

	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contactID)}, bson.M{"$set": _mongoM})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) UpdateContact(contactID string, mmap map[string]string) (r string) {
	if !bson.IsObjectIdHex(contactID) {
		return ""
	}

	_mongoM := bson.M{}

	for k, v := range mmap {
		_mongoM[k] = v
	}

	_err := ContactCollection.Update(bson.M{"_id": bson.ObjectIdHex(contactID)}, bson.M{"$set": _mongoM})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}
