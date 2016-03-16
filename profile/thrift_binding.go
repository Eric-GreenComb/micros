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

func (tb thriftBinding) GetProfileByCat(name string) (string, error) {
	r := tb.ProfileService.GetProfileByCat(name)
	return r, nil
}

func (tb thriftBinding) GetProfileBySubCat(name string) (string, error) {
	r := tb.ProfileService.GetProfileBySubCat(name)
	return r, nil
}
