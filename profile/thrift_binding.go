package main

import (
	"github.com/banerwai/micros/profile/service"
)

type thriftBinding struct {
	service.ProfileService
}

func (tb thriftBinding) GetProfile(profile_id string) (string, error) {
	r := tb.ProfileService.GetProfile(profile_id)
	return r, nil
}
