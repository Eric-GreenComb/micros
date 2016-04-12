package main

import (
	"github.com/banerwai/micros/query/profile/service"
)

type thriftBinding struct {
	service.ProfileService
}

func (tb thriftBinding) GetProfile(profile_id string) (string, error) {
	r := tb.ProfileService.GetProfile(profile_id)
	return r, nil
}

func (tb thriftBinding) SearchProfiles(json_search string, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.SearchProfiles(json_search, timestamp, pagesize)
	return r, nil
}
