package service

import ()

// Service is the abstract representation of this service.
type ResumeService interface {
	Ping() string
	// Parameters:
	//  - ID
	GetResume(id string) string
}
