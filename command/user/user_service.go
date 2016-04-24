package main

import (
	"github.com/banerwai/micros/command/user/service"
	"labix.org/v2/mgo/bson"
	"time"
)

type inmemService struct {
}

func newInmemService() service.UserService {
	return &inmemService{}
}

func (self *inmemService) CreateUser(mmap map[string]string) (r string) {
	// var _user bean.User
	// _user.Email = email
	// _user.Pwd = pwd
	// if bson.IsObjectIdHex(invited) {
	// 	_user.Invited = bson.ObjectIdHex(invited)
	// } else {
	// 	_user.Invited = bson.ObjectIdHex(DefaultUserObjectId)
	// }
	// var _temp bson.M
	// err := UsersCollection.Find(bson.M{"email": email}).One(&_temp)
	// if err != nil {
	// 	return err.Error()
	// }

	// email is a index, if email has ,insert is err
	_mongo_m := bson.M{}

	for k, v := range mmap {
		_mongo_m[k] = v
	}

	_time := time.Now()

	_mongo_m["createdtime"] = _time
	_mongo_m["lastactivity"] = _time
	_mongo_m["actived"] = false

	err := UsersCollection.Insert(_mongo_m)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (self *inmemService) ResetPwd(email string, newpwd string) (r bool) {
	r = true
	err := UsersCollection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"pwd": newpwd}})
	if nil != err {
		r = false
	}

	return
}

func (self *inmemService) ActiveUser(email string) (r bool) {
	r = true
	err := UsersCollection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"actived": true}})
	if nil != err {
		r = false
	}
	return
}
