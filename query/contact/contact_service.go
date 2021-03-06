package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"

	"github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/etcd"
	"github.com/banerwai/micros/query/contact/service"
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

func (ims *inmemService) GetContactTpl(tplname string) (r string) {
	_tpl, _err := ims.getTplFromEtcd(tplname)
	if _err != nil {
		return ""
	}

	return _tpl
}

func (ims *inmemService) GetContact(contactid string) (r string) {
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

func (ims *inmemService) GetContactSignStatus(contactid string) (r string) {
	if !bson.IsObjectIdHex(contactid) {
		return ""
	}

	var _bsonM bson.M
	_selector := bson.M{}
	_selector["client_signup"] = 1
	_selector["freelancer_signup"] = 1

	err := ContactCollection.Find(bson.M{"_id": bson.ObjectIdHex(contactid)}).Select(_selector).One(&_bsonM)
	if err != nil {
		return err.Error()
	}

	_data, _err := bson.Marshal(_bsonM)
	if _err != nil {
		return _err.Error()
	}

	r = string(_data)
	return
}

func (ims *inmemService) GetClientContact(clientemail string) (r string) {
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

func (ims *inmemService) GetFreelancerContact(freelanceremail string) (r string) {
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

func (ims *inmemService) getTplFromEtcd(tplname string) (string, error) {
	_key := constant.EtcdKeyTplContact + tplname
	_tpl, _err := etcd.GetValue(_key)

	if _err == nil {
		return _tpl, nil
	}

	_key = constant.EtcdKeyTplContact + "default"
	_tpl, _err = etcd.GetValue(_key)
	if _err != nil {
		return "", _err
	}
	return _tpl, nil
}
