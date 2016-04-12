package main

import (
	"github.com/banerwai/micros/command/profile/service"
)

type inmemService struct {
}

func newInmemService() service.ProfileService {
	return &inmemService{}
}

func (self *inmemService) AddProfile(json_profile string) (r string) {
	r = json_profile
	return
}

func (self *inmemService) UpdateProfile(json_profile string) (r string) {
	r = json_profile
	return
}

func (self *inmemService) DeleteProfile(id string) (r string) {
	r = id
	return
}
