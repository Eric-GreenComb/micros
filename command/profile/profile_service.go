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

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) AddProfile(json_profile string) (r string) {
	var _profile bean.Profile
	err := json.Unmarshal([]byte(json_profile), &_profile)
	if err != nil {
		return err.Error()
	}
	_profile.Id = ""

	_time := time.Now()

	_profile.CreatedTime = _time
	_profile.LastActiveTime = _time
	_profile.Status = true

	_err := ProfileCollection.Insert(_profile)
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateProfile(profile_id string, json_profile string) (r string) {
	var _profile bean.Profile
	err := json.Unmarshal([]byte(json_profile), &_profile)
	if err != nil {
		return err.Error()
	}

	_time := time.Now()
	_profile.LastActiveTime = _time

	_profile.Id = ""
	_err := ProfileCollection.Update(bson.M{"_id": bson.ObjectIdHex(profile_id)}, bson.M{"$set": _profile})
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateProfileStatus(profile_id string, status bool) (r string) {
	_err := ProfileCollection.Update(bson.M{"_id": bson.ObjectIdHex(profile_id)}, bson.M{"$set": bson.M{"status": status}})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) UpdateProfileBase(profile_id string, mmap map[string]string) (r string) {
	_mongo_m := bson.M{}

	for k, v := range mmap {
		switch k {
		case "hour_rate":
			_i, _ := strconv.Atoi(v)
			_mongo_m[k] = _i
		case "work_hours":
			_i, _ := strconv.Atoi(v)
			_mongo_m[k] = _i
		case "portfolio_nums":
			_i, _ := strconv.Atoi(v)
			_mongo_m[k] = _i
		default:
			_mongo_m[k] = v
		}
	}

	_time := time.Now()
	_mongo_m["last_activetime"] = _time

	_err := ProfileCollection.Update(bson.M{"_id": bson.ObjectIdHex(profile_id)}, bson.M{"$set": _mongo_m})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) UpdateProfileAgencyMembers(profile_id string, agency_members string) (r string) {
	var _beans []bean.AgencyMember
	err := json.Unmarshal([]byte(agency_members), &_beans)
	if err != nil {
		return err.Error()
	}

	_, _err := ProfileCollection.Upsert(bson.M{"_id": bson.ObjectIdHex(profile_id)}, bson.M{"$set": bson.M{"agency_members": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}
