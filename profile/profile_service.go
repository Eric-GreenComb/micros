package main

import (
	"github.com/banerwai/micros/profile/service"
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
