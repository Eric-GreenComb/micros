package main

import (
	"strconv"

	banerwaistrings "github.com/banerwai/gommon/strings"
	"github.com/banerwai/micros/profile/service"
	thriftprofile "github.com/banerwai/micros/profile/thrift/gen-go/profile"
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

func (self *inmemService) SearchProfiles(profile_search_condition *thriftprofile.ProfileSearchCondition) (r string) {
	var SerialNumber, HoursBilled, AvailableHours string

	SerialNumber = strconv.Itoa(int(profile_search_condition.SerialNumber)) + ";"

	if profile_search_condition.HoursBilled == -1 {
		HoursBilled = "All HoursBilled;"
	}
	if profile_search_condition.AvailableHours == -1 {
		AvailableHours = "All AvailableHours;"
	}

	r = banerwaistrings.ConstructString(SerialNumber, HoursBilled, AvailableHours)
	return
}
