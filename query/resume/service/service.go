package service

import ()

// Service is the abstract representation of this service.
type ResumeService interface {
	Ping() string
	// Parameters:
	//  - Userid
	GetResume(userID string) string
}
