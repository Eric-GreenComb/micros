package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"

	"github.com/banerwai/micros/query/profile/service"
)

type inmemService struct {
}

func newInmemService() service.ProfileService {
	return &inmemService{}
}

func (self *inmemService) GetProfile(id string) (r string) {
	r = id
	return
}

func (self *inmemService) GetProfilesByEmail(email string) (r string) {
	r = email
	return
}

//  db.profile.find({createdtime:{$gt:timestamp}}).sort({"createdtime":1}).limit(10)
func (self *inmemService) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (r string) {
	// db.profile.find(self.genQuery(profile_search_condition, timestamp)).sort({"createdtime":1}).limit(pagesize)
	_query := self.genQuery(option_mmap, key_mmap, timestamp)
	fmt.Println("%q", _query)
	r = "OK"
	return
}

//	query := bson.M{"serial_number": ...}
func (self *inmemService) genQuery(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64) interface{} {
	query := bson.M{"createdtime": bson.M{"$gt": timestamp}}

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

	// search overview by key
	if overview, ok := key_mmap["overview"]; ok {
		query["overview"] = overview
	}

	return query
}
