package main

import (
	"time"

	"github.com/banerwai/gommon/uuid"
	"github.com/banerwai/micros/command/token/service"

	"gopkg.in/mgo.v2/bson"
)

type inmemService struct {
}

func newInmemService() service.TokenService {
	return &inmemService{}
}

func (ims *inmemService) Ping() string {
	return "pong"
}

func (ims *inmemService) CreateToken(key string, ttype int64) string {
	_uuid := uuid.UUID()

	_mongoM := bson.M{}
	_mongoM["key"] = key
	_mongoM["token"] = _uuid
	_mongoM["type"] = ttype
	_mongoM["createdtime"] = time.Now()

	_, _err := TokenCollection.Upsert(bson.M{"key": key, "type": ttype}, _mongoM)
	if _err != nil {
		return _err.Error()
	}
	return _uuid
}

func (ims *inmemService) DeleteToken(key string, ttype int64) bool {
	TokenCollection.Remove(bson.M{"key": key, "type": ttype})
	return true
}
