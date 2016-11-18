package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
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

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) GetProfile(profileID string) (r string) {
	if !bson.IsObjectIdHex(profileID) {
		return ""
	}
	var _profile bean.Profile
	err := ProfileCollection.Find(bson.M{"_id": bson.ObjectIdHex(profileID)}).One(&_profile)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profile)
	r = string(b)
	return
}

func (ims *inmemService) GetProfilesByUserID(userID string) (r string) {
	if !bson.IsObjectIdHex(userID) {
		return ""
	}

	var _profiles []bean.Profile

	err := ProfileCollection.Find(bson.M{"user_id": bson.ObjectIdHex(userID)}).All(&_profiles)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profiles)
	r = string(b)
	return
}

func (ims *inmemService) GetProfilesByCategory(categoryID int64, timestamp int64, pagesize int64) (r string) {
	var _profiles []bean.Profile

	query := bson.M{"last_activetime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = true
	query["category_number"] = categoryID

	err := ProfileCollection.Find(query).Sort("-last_activetime").Limit(int(pagesize)).All(&_profiles)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profiles)
	r = string(b)
	return
}

func (ims *inmemService) GetProfilesBySubCategory(subcategoryID int64, timestamp int64, pagesize int64) (r string) {
	var _profiles []bean.Profile

	query := bson.M{"last_activetime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = true
	query["serial_number"] = subcategoryID

	err := ProfileCollection.Find(query).Sort("-last_activetime").Limit(int(pagesize)).All(&_profiles)

	if err != nil {
		return ""
	}

	b, _ := json.Marshal(_profiles)
	r = string(b)
	return
}

//  db.profile.find({createdtime:{$gt:timestamp}}).sort({"createdtime":1}).limit(10)
func (ims *inmemService) SearchProfiles(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) (r string) {
	// db.profile.find(self.genQuery(profile_search_condition, timestamp)).sort({"createdtime":1}).limit(pagesize)
	_query := ims.genQuery(optionMap, keyMap, timestamp)
	r = ims.searchProfile(_query, pagesize)
	return
}

func (ims *inmemService) searchProfile(q interface{}, pagesize int64) (r string) {
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
func (ims *inmemService) genQuery(optionMap map[string]int64, keyMap map[string]string, timestamp int64) interface{} {
	query := bson.M{"last_activetime": bson.M{"$lt": time.Unix(timestamp, 0)}}
	query["status"] = true
	// query := bson.M{}

	// search option by search options
	if serialNumber, ok := optionMap["serial_number"]; ok {
		query["serial_number"] = serialNumber
	}

	if hoursBilled, ok := optionMap["hours_billed"]; ok {
		query["hours_billed"] = hoursBilled
	}

	if availableHours, ok := optionMap["available_hours"]; ok {
		query["available_hours"] = availableHours
	}

	if jobSuccess, ok := optionMap["job_success"]; ok {
		query["job_success"] = jobSuccess
	}

	if lastActivity, ok := optionMap["last_activity"]; ok {
		query["last_activity"] = lastActivity
	}

	if freelancerType, ok := optionMap["freelancer_type"]; ok {
		query["freelancer_type"] = freelancerType
	}

	if hourlyRate, ok := optionMap["hourly_rate"]; ok {
		query["hourly_rate"] = hourlyRate
	}

	if regionID, ok := optionMap["region_id"]; ok {
		query["region_id"] = regionID
	}

	if len(keyMap) == 0 {
		return query
	}

	// search overview by key
	var _bsons []bson.M

	for k, v := range keyMap {
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
