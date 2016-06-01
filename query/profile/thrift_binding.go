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

func (tb thriftBinding) GetProfile(profile_id string) (string, error) {
	r := tb.ProfileService.GetProfile(profile_id)
	return r, nil
}

func (tb thriftBinding) GetProfilesByUserId(user_id string) (string, error) {
	r := tb.ProfileService.GetProfilesByUserId(user_id)
	return r, nil
}

func (tb thriftBinding) GetProfilesByCategory(category_id int64, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.GetProfilesByCategory(category_id, timestamp, pagesize)
	return r, nil
}

func (tb thriftBinding) GetProfilesBySubCategory(subcategory_id int64, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.GetProfilesBySubCategory(subcategory_id, timestamp, pagesize)
	return r, nil
}

func (tb thriftBinding) SearchProfiles(option_mmap map[string]int64, key_mmap map[string]string, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.SearchProfiles(option_mmap, key_mmap, timestamp, pagesize)
	return r, nil
}
