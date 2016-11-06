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

func (tb thriftBinding) AddProfile(jsonProfile string) (string, error) {
	r := tb.ProfileService.AddProfile(jsonProfile)
	return r, nil
}

func (tb thriftBinding) UpdateProfile(profileID string, jsonProfile string) (string, error) {
	r := tb.ProfileService.UpdateProfile(profileID, jsonProfile)
	return r, nil
}

func (tb thriftBinding) UpdateProfileStatus(profileID string, status bool) (string, error) {
	r := tb.ProfileService.UpdateProfileStatus(profileID, status)
	return r, nil
}

func (tb thriftBinding) UpdateProfileBase(profileID string, mmap map[string]string) (string, error) {
	r := tb.ProfileService.UpdateProfileBase(profileID, mmap)
	return r, nil
}

func (tb thriftBinding) UpdateProfileAgencyMembers(profileID string, agencyMembers string) (string, error) {
	r := tb.ProfileService.UpdateProfileAgencyMembers(profileID, agencyMembers)
	return r, nil
}
