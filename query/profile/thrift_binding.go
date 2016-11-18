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

func (tb thriftBinding) GetProfile(profileID string) (string, error) {
	r := tb.ProfileService.GetProfile(profileID)
	return r, nil
}

func (tb thriftBinding) GetProfilesByUserID(userID string) (string, error) {
	r := tb.ProfileService.GetProfilesByUserID(userID)
	return r, nil
}

func (tb thriftBinding) GetProfilesByCategory(categoryID int64, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.GetProfilesByCategory(categoryID, timestamp, pagesize)
	return r, nil
}

func (tb thriftBinding) GetProfilesBySubCategory(subcategoryID int64, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.GetProfilesBySubCategory(subcategoryID, timestamp, pagesize)
	return r, nil
}

func (tb thriftBinding) SearchProfiles(optionMap map[string]int64, keyMap map[string]string, timestamp int64, pagesize int64) (string, error) {
	r := tb.ProfileService.SearchProfiles(optionMap, keyMap, timestamp, pagesize)
	return r, nil
}
