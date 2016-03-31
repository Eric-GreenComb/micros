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

func (tb thriftBinding) SearchProfiles(profile_search_condition *thriftprofile.ProfileSearchCondition) (string, error) {
	r := tb.ProfileService.SearchProfiles(profile_search_condition)
	return r, nil
}
