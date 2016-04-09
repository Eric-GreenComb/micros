package main

import (
	"fmt"
	"labix.org/v2/mgo/bson"

	"github.com/banerwai/micros/command/profile/service"
	thriftprofile "github.com/banerwai/micros/command/profile/thrift/gen-go/profile"
)

type inmemService struct {
}

func newInmemService() service.ProfileService {
	return &inmemService{}
}

func (self *inmemService) GetProfile(profile_id string) (r string) {
	r = profile_id
	return
}

//  db.profile.find({createdtime:{$gt:timestamp}}).sort({"createdtime":1}).limit(10)
func (self *inmemService) SearchProfiles(profile_search_condition *thriftprofile.ProfileSearchCondition, timestamp int64, pagesize int64) (r string) {
	// db.profile.find(self.genQuery(profile_search_condition, timestamp)).sort({"createdtime":1}).limit(pagesize)
	fmt.Println(self.genQuery(profile_search_condition, timestamp))
	r = "OK"
	return
}

//	query := bson.M{"serial_number": ...}
func (self *inmemService) genQuery(profile_search_condition *thriftprofile.ProfileSearchCondition, timestamp int64) interface{} {
	query := bson.M{"createdtime": bson.M{"$gt": timestamp}}

	if profile_search_condition.SerialNumber != -1 {
		query["serial_number"] = profile_search_condition.SerialNumber
	}
	if profile_search_condition.HoursBilled != -1 {
		query["hours_billed"] = profile_search_condition.HoursBilled
	}
	if profile_search_condition.AvailableHours == -1 {
		query["available_hours"] = profile_search_condition.AvailableHours
	}
	if profile_search_condition.JobSuccess != -1 {
		query["job_success"] = profile_search_condition.JobSuccess
	}
	if profile_search_condition.LastActivity != -1 {
		query["last_activity"] = profile_search_condition.LastActivity
	}
	if profile_search_condition.FreelancerType != -1 {
		query["freelancer_type"] = profile_search_condition.FreelancerType
	}
	if profile_search_condition.HourlyRate != -1 {
		query["hourly_rate"] = profile_search_condition.HourlyRate
	}
	if profile_search_condition.RegionID != -1 {
		query["region_id"] = profile_search_condition.RegionID
	}
	return query
}
