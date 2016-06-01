package main

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"

	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/query/profile/service"
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

func (self *inmemService) GetProfile(profile_id string) (r string) {
	if !bson.IsObjectIdHex(profile_id) {
		return ""
	}
	var _profile bean.Profile
	err := ProfileCollection.Find(bson.M{"_id": bson.ObjectIdHex(profile_id)}).One(&_profile)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profile)
	r = string(b)
	return
}

func (self *inmemService) GetProfilesByUserId(user_id string) (r string) {
	if !bson.IsObjectIdHex(user_id) {
		return ""
	}

	var _profiles []bean.Profile

	err := ProfileCollection.Find(bson.M{"user_id": bson.ObjectIdHex(user_id)}).All(&_profiles)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profiles)
	r = string(b)
	return
}

func (self *inmemService) GetProfilesByCategory(category_id int64, timestamp int64, pagesize int64) (r string) {
	var _profiles []bean.Profile

	query := bson.M{"last_activetime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = true
	query["category_number"] = category_id

	err := ProfileCollection.Find(query).Sort("-last_activetime").Limit(int(pagesize)).All(&_profiles)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profiles)
	r = string(b)
	return
}

func (self *inmemService) GetProfilesBySubCategory(subcategory_id int64, timestamp int64, pagesize int64) (r string) {
	var _profiles []bean.Profile

	query := bson.M{"last_activetime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = true
	query["serial_number"] = subcategory_id

	err := ProfileCollection.Find(query).Sort("-last_activetime").Limit(int(pagesize)).All(&_profiles)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profiles)
	r = string(b)
	return
}

//  db.profile.find({createdtime:{$gt:timestamp}}).sort({"createdtime":1}).limit(10)
func (self *inmemService) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (r string) {
	// db.profile.find(self.genQuery(profile_search_condition, timestamp)).sort({"createdtime":1}).limit(pagesize)
	_query := self.genQuery(option_mmap, key_mmap, timestamp)
	r = self.searchProfile(_query, pagesize)
	return
}

func (self *inmemService) searchProfile(q interface{}, pagesize int64) (r string) {
	var _profiles []bean.Profile
	err := ProfileCollection.Find(q).Sort("-last_activetime").Limit(int(pagesize)).All(&_profiles)
	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profiles)
	r = string(b)
	return
}

//	query := bson.M{"serial_number": ...}
func (self *inmemService) genQuery(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64) interface{} {
	query := bson.M{"last_activetime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = true
	// query := bson.M{}

	// search option by search options
	if serial_number, ok := option_mmap["serial_number"]; ok {
		query["serial_number"] = serial_number
	}

	if hours_billed, ok := option_mmap["hours_billed"]; ok {
		query["hours_billed"] = hours_billed
	}

	if available_hours, ok := option_mmap["available_hours"]; ok {
		query["available_hours"] = available_hours
	}

	if job_success, ok := option_mmap["job_success"]; ok {
		query["job_success"] = job_success
	}

	if last_activity, ok := option_mmap["last_activity"]; ok {
		query["last_activity"] = last_activity
	}

	if freelancer_type, ok := option_mmap["freelancer_type"]; ok {
		query["freelancer_type"] = freelancer_type
	}

	if hourly_rate, ok := option_mmap["hourly_rate"]; ok {
		query["hourly_rate"] = hourly_rate
	}

	if region_id, ok := option_mmap["region_id"]; ok {
		query["region_id"] = region_id
	}

	if len(key_mmap) == 0 {
		return query
	}

	// search overview by key
	var _bsons []bson.M

	for k, v := range key_mmap {
		_bson := bson.M{}
		_regex := bson.RegEx{v, "i"}
		_split := strings.Split(k, "/")
		if len(_split) > 1 {
			_bson[_split[0]] = _regex
		} else {
			_bson[k] = _regex
		}

		_bsons = append(_bsons, _bson)
	}
	query["$and"] = _bsons

	return query
}
