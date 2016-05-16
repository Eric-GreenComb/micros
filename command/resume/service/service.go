package service

import ()

// Service is the abstract representation of this service.
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
	UpdateResumeSkillExperience(userid string, experience_levels string) string
	// Parameters:
	//  - Userid
	//  - ToolArchs
	UpdateResumeToolandArchs(userid string, tool_archs string) string
	// Parameters:
	//  - Userid
	//  - Portfolioes
	UpdateResumePortfolioes(userid string, portfolioes string) string
	// Parameters:
	//  - Userid
	//  - EmploymentHistories
	UpdateResumeEmploymentHistories(userid string, employment_histories string) string
	// Parameters:
	//  - Userid
	//  - Educations
	UpdateResumeEducations(userid string, educations string) string
	// Parameters:
	//  - Userid
	//  - OtherExperiences
	UpdateResumeOtherExperiences(userid string, other_experiences string) string
}
