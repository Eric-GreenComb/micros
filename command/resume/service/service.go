package service

import ()

// ResumeService is the abstract representation of this service.
type ResumeService interface {
	Ping() string
	// Parameters:
	//  - Resume
	AddResume(resume string) string
	// Parameters:
	//  - Userid
	//  - Resume
	UpdateResume(userid string, resume string) string
	// Parameters:
	//  - Userid
	//  - Mmap
	UpdateResumeBase(userid string, mmap map[string]string) string
	// Parameters:
	//  - Userid
	//  - ExperienceLevels
	UpdateResumeSkillExperience(userid string, experienceLevels string) string
	// Parameters:
	//  - Userid
	//  - ToolArchs
	UpdateResumeToolandArchs(userid string, toolArchs string) string
	// Parameters:
	//  - Userid
	//  - Portfolioes
	UpdateResumePortfolioes(userid string, portfolioes string) string
	// Parameters:
	//  - Userid
	//  - EmploymentHistories
	UpdateResumeEmploymentHistories(userid string, employmentHistories string) string
	// Parameters:
	//  - Userid
	//  - Educations
	UpdateResumeEducations(userid string, educations string) string
	// Parameters:
	//  - Userid
	//  - OtherExperiences
	UpdateResumeOtherExperiences(userid string, otherExperiences string) string
}
