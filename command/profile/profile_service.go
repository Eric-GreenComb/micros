package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/command/profile/service"
	"labix.org/v2/mgo/bson"
	"strconv"
	"time"
)

type inmemService struct {
}

func newInmemService() service.ProfileService {
	return &inmemService{}
}

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) AddProfile(jsonProfile string) (r string) {
	var _profile bean.Profile
	err := json.Unmarshal([]byte(jsonProfile), &_profile)
	if err != nil {
		return err.Error()
	}
	_profile.ID = bson.NewObjectId()

	_time := time.Now()

	_profile.CreatedTime = _time
	_profile.LastActiveTime = _time
	_profile.Status = true

	_err := ProfileCollection.Insert(_profile)
	if _err != nil {
		return _err.Error()
	}
	return _profile.ID.Hex()
}

func (ims *inmemService) UpdateProfile(profileID string, jsonProfile string) (r string) {
	var _profile bean.Profile
	err := json.Unmarshal([]byte(jsonProfile), &_profile)
	if err != nil {
		return err.Error()
	}

	_profile.ID = ""
	_err := ProfileCollection.Update(bson.M{"_id": bson.ObjectIdHex(profileID)}, bson.M{"$set": _profile})
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (ims *inmemService) UpdateProfileStatus(profileID string, status bool) (r string) {
	_err := ProfileCollection.Update(bson.M{"_id": bson.ObjectIdHex(profileID)}, bson.M{"$set": bson.M{"status": status}})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) UpdateProfileBase(profileID string, mmap map[string]string) (r string) {
	_mongoM := bson.M{}

	for k, v := range mmap {
		switch k {
		case "hour_rate":
			_i, _ := strconv.Atoi(v)
			_mongoM[k] = _i
		case "work_hours":
			_i, _ := strconv.Atoi(v)
			_mongoM[k] = _i
		case "portfolio_nums":
			_i, _ := strconv.Atoi(v)
			_mongoM[k] = _i
		default:
			_mongoM[k] = v
		}
	}

	_err := ProfileCollection.Update(bson.M{"_id": bson.ObjectIdHex(profileID)}, bson.M{"$set": _mongoM})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) UpdateProfileAgencyMembers(profileID string, agencyMembers string) (r string) {
	var _beans []bean.AgencyMember
	err := json.Unmarshal([]byte(agencyMembers), &_beans)
	if err != nil {
		return err.Error()
	}

	_, _err := ProfileCollection.Upsert(bson.M{"_id": bson.ObjectIdHex(profileID)}, bson.M{"$set": bson.M{"agency_members": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}
