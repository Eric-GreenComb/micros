package main

import (
	"github.com/banerwai/micros/command/profile/service"
)

type thriftBinding struct {
	service.ProfileService
}

func (tb thriftBinding) AddProfile(json_profile string) (string, error) {
	r := tb.ProfileService.AddProfile(json_profile)
	return r, nil
}

func (tb thriftBinding) UpdateProfile(json_profile string) (string, error) {
	r := tb.ProfileService.UpdateProfile(json_profile)
	return r, nil
}

func (tb thriftBinding) DeleteProfile(id string) (string, error) {
	r := tb.ProfileService.DeleteProfile(id)
	return r, nil
}
