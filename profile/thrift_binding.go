package main

import (
	"github.com/banerwai/micros/profile/service"

	thriftprofile "github.com/banerwai/micros/profile/thrift/gen-go/profile"
)

type thriftBinding struct {
	service.ProfileService
}

func (tb thriftBinding) GetProfile(profile_id string) (string, error) {
	r := tb.ProfileService.GetProfile(profile_id)
	return r, nil
}

func (tb thriftBinding) SearchProfiles(profile_search_condition *thriftprofile.ProfileSearchCondition, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.SearchProfiles(profile_search_condition, timestamp, pagesize)
	return r, nil
}
