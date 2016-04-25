package main

import (
	"time"

	"github.com/banerwai/gommon/uuid"
	"github.com/banerwai/micros/command/token/service"

	"labix.org/v2/mgo/bson"
)

type inmemService struct {
}

func newInmemService() service.TokenService {
	return &inmemService{}
}

func (self *inmemService) NewToken_(key string, ttype int64) string {
	_uuid := uuid.UUID()

	_mongo_m := bson.M{}
	_mongo_m["key"] = key
	_mongo_m["token"] = _uuid
	_mongo_m["type"] = ttype
	_mongo_m["createdtime"] = time.Now()

	_, _err := TokenCollection.Upsert(bson.M{"key": key, "type": ttype}, _mongo_m)
	if _err != nil {
		return _err.Error()
	}
	return _uuid
}

func (self *inmemService) DeleteToken(key string, ttype int64) bool {
	TokenCollection.Remove(bson.M{"key": key, "type": ttype})
	return true
}
