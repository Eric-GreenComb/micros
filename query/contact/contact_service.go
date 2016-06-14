package main

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"

	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/gommon/etcd"
	"github.com/banerwai/micros/query/contact/service"
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

func (self *inmemService) GetContactTpl(tplname string) (r string) {
	_tpl, _err := self.getTplFromEtcd(tplname)
	if _err != nil {
		return ""
	}

	return _tpl
}

func (self *inmemService) GetContact(contactid string) (r string) {
	if !bson.IsObjectIdHex(contactid) {
		return ""
	}

	var _contact bean.Contact

	err := ContactCollection.Find(bson.M{"_id": bson.ObjectIdHex(contactid)}).One(&_contact)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_contact)
	r = string(b)
	return
}

func (self *inmemService) GetContactSignStatus(contactid string) (r string) {
	if !bson.IsObjectIdHex(contactid) {
		return ""
	}

	var _bson_m bson.M
	_selector := bson.M{}
	_selector["client_signup"] = 1
	_selector["freelancer_signup"] = 1

	err := ContactCollection.Find(bson.M{"_id": bson.ObjectIdHex(contactid)}).Select(_selector).One(&_bson_m)
	if err != nil {
		return err.Error()
	}

	_data, _err := bson.Marshal(_bson_m)
	if _err != nil {
		return _err.Error()
	}

	r = string(_data)
	return
}

func (self *inmemService) GetClientContact(clientemail string) (r string) {
	var _contacts []bean.Contact

	query := bson.M{}
	query["dealed"] = true
	query["client_email"] = clientemail

	err := ContactCollection.Find(query).Sort("-createdtime").All(&_contacts)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_contacts)
	r = string(b)
	return
}

func (self *inmemService) GetFreelancerContact(freelanceremail string) (r string) {
	var _contacts []bean.Contact

	query := bson.M{}
	query["dealed"] = true
	query["freelancer_email"] = freelanceremail

	err := ContactCollection.Find(query).Sort("-createdtime").All(&_contacts)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_contacts)
	r = string(b)
	return
}

func (self *inmemService) getTplFromEtcd(tplname string) (string, error) {
	_key := global.ETCD_KEY_TPL_CONTACT + tplname
	_tpl, _err := etcd.GetValue(_key)
	if _err != nil {
		return "", _err
	}
	return _tpl, nil
}
