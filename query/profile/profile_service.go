package main

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo/bson"

	"github.com/banerwai/gather/query/dto"
	"github.com/banerwai/micros/query/profile/service"
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
func (self *inmemService) SearchProfiles(json_search string, timestamp int64, pagesize int64) (r string) {
	// db.profile.find(self.genQuery(profile_search_condition, timestamp)).sort({"createdtime":1}).limit(pagesize)

	prfile_search_dto, err := self.genProfileSearchDto(json_search)
	if err != nil {
		r = err.Error()
		return
	}
	_query := self.genQuery(prfile_search_dto, timestamp)
	fmt.Println("%q", _query)
	r = "OK"
	return
}

func (self *inmemService) genProfileSearchDto(json_search string) (dto.ProfileSearchDto, error) {
	var prfile_search_dto dto.ProfileSearchDto
	err := json.Unmarshal([]byte(json_search), &prfile_search_dto)
	return prfile_search_dto, err
}

//	query := bson.M{"serial_number": ...}
func (self *inmemService) genQuery(prfile_search_dto dto.ProfileSearchDto, timestamp int64) interface{} {
	query := bson.M{"createdtime": bson.M{"$gt": timestamp}}

	if prfile_search_dto.SerialNumber != -1 {
		query["serial_number"] = prfile_search_dto.SerialNumber
	}
	if prfile_search_dto.HoursBilled != -1 {
		query["hours_billed"] = prfile_search_dto.HoursBilled
	}
	if prfile_search_dto.AvailableHours != -1 {
		query["available_hours"] = prfile_search_dto.AvailableHours
	}
	if prfile_search_dto.JobSuccess != -1 {
		query["job_success"] = prfile_search_dto.JobSuccess
	}
	if prfile_search_dto.LastActivity != -1 {
		query["last_activity"] = prfile_search_dto.LastActivity
	}
	if prfile_search_dto.FreelancerType != -1 {
		query["freelancer_type"] = prfile_search_dto.FreelancerType
	}
	if prfile_search_dto.HourlyRate != -1 {
		query["hourly_rate"] = prfile_search_dto.HourlyRate
	}
	if prfile_search_dto.RegionId != -1 {
		query["region_id"] = prfile_search_dto.RegionId
	}
	return query
}
