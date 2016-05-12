package service

import ()

// Service is the abstract representation of this service.
type ResumeService interface {
	Ping() string
	// Parameters:
	//  - JSONResume
	AddResume(json_resume string) string
	// Parameters:
	//  - JSONResume
	UpdateResume(json_resume string) string
}
