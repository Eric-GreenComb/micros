package main

import (
	"github.com/banerwai/micros/query/profile/service"
)

type thriftBinding struct {
	service.ProfileService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.ProfileService.Ping()
	return r, nil
}

func (tb thriftBinding) GetProfile(id string) (string, error) {
	r := tb.ProfileService.GetProfile(id)
	return r, nil
}

func (tb thriftBinding) GetProfilesByEmail(email string) (string, error) {
	r := tb.ProfileService.GetProfilesByEmail(email)
	return r, nil
}

func (tb thriftBinding) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	return r, nil
}
