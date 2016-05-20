package main

import (
	"github.com/banerwai/micros/command/profile/service"
)

type thriftBinding struct {
	service.ProfileService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.ProfileService.Ping()
	return r, nil
}

func (tb thriftBinding) AddProfile(json_profile string) (string, error) {
	r := tb.ProfileService.AddProfile(json_profile)
	return r, nil
}

func (tb thriftBinding) UpdateProfile(profile_id string, json_profile string) (string, error) {
	r := tb.ProfileService.UpdateProfile(profile_id, json_profile)
	return r, nil
}

func (tb thriftBinding) UpdateProfileStatus(profile_id string, status bool) (string, error) {
	r := tb.ProfileService.UpdateProfileStatus(profile_id, status)
	return r, nil
}

func (tb thriftBinding) UpdateProfileBase(profile_id string, mmap map[string]string) (string, error) {
	r := tb.ProfileService.UpdateProfileBase(profile_id, mmap)
	return r, nil
}

func (tb thriftBinding) UpdateProfileAgencyMembers(profile_id string, agency_members string) (string, error) {
	r := tb.ProfileService.UpdateProfileAgencyMembers(profile_id, agency_members)
	return r, nil
}
