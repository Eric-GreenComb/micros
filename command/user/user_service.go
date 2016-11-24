package main

import (
	"github.com/banerwai/global/bean"
	global "github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/crypto"
	"github.com/banerwai/micros/command/user/service"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type inmemService struct {
}

func newInmemService() service.UserService {
	return &inmemService{}
}

func (ims *inmemService) Ping() (r string) {
	return "pong"
}

func (ims *inmemService) CreateUser(mmap map[string]string) (r string) {
	// var _user bean.User
	// _user.Email = email
	// _user.Pwd = pwd
	// if bson.IsObjectIdHex(invited) {
	// 	_user.Invited = bson.ObjectIdHex(invited)
	// } else {
	// 	_user.Invited = bson.ObjectIdHex(DefaultUserObjectId)
	// }
	// _email := mmap["email"]
	// var _temp bson.M
	// err := UsersCollection.Find(bson.M{"email": email}).One(&_temp)
	// if err != nil {
	// 	return err.Error()
	// }

	// email is a index, if email has ,insert is err
	_mongoM := bson.M{}

	// if _, ok := mmap["pwd"]; ok {
	// 	_b, _ := crypto.GenerateHash(mmap["pwd"])
	// 	mmap["pwd"] = string(_b)
	// }

	for k, v := range mmap {
		switch k {
		case "pwd":
			_b, _ := crypto.GenerateHash(v)
			_mongoM[k] = string(_b)
		case "invited":
			_mongoM[k] = bson.ObjectIdHex(v)
		default:
			_mongoM[k] = v
		}
	}

	_time := time.Now().Unix()

	_id := bson.NewObjectId()
	_mongoM["_id"] = _id
	_mongoM["createdtime"] = _time
	_mongoM["actived"] = false

	err := UsersCollection.Insert(_mongoM)
	if err != nil {
		return err.Error()
	}

	ims.createAccount(_id, mmap["email"])

	return _id.Hex()
}

func (ims *inmemService) createAccount(userID bson.ObjectId, email string) error {
	var _account bean.Account
	_account.UserID = userID
	_account.Email = email
	_account.CreatedTime = time.Now().Unix()

	var _lsMultiCurrencyAccount []bean.MultiCurrencyAccount

	var _beanMultiCurrencyAccount01 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount01.Currency = global.CurrencyUSD
	_beanMultiCurrencyAccount01.Amount = 0
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount01)

	var _beanMultiCurrencyAccount02 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount02.Currency = global.CurrencyCNY
	_beanMultiCurrencyAccount02.Amount = 0
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount02)

	var _beanMultiCurrencyAccount03 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount03.Currency = global.CurrencyEUR
	_beanMultiCurrencyAccount03.Amount = 0
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount03)

	var _beanMultiCurrencyAccount04 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount04.Currency = global.CurrencyJPY
	_beanMultiCurrencyAccount04.Amount = 0
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount04)

	var _beanMultiCurrencyAccount05 bean.MultiCurrencyAccount
	_beanMultiCurrencyAccount05.Currency = global.CurrencyGBP
	_beanMultiCurrencyAccount05.Amount = 0
	_lsMultiCurrencyAccount = append(_lsMultiCurrencyAccount, _beanMultiCurrencyAccount05)

	_account.MultiCurrency = _lsMultiCurrencyAccount

	_err := AccountCollection.Insert(_account)
	return _err
}

func (ims *inmemService) ResetPwd(email string, newpwd string) (r bool) {
	r = true
	_b, _ := crypto.GenerateHash(newpwd)
	err := UsersCollection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"pwd": string(_b)}})
	if nil != err {
		r = false
	}

	return
}

func (ims *inmemService) ActiveUser(email string) (r bool) {
	r = true
	err := UsersCollection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{"actived": true}})
	if nil != err {
		r = false
	}
	return
}
