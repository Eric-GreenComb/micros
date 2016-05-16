package main

import (
	"github.com/banerwai/micros/command/resume/service"
)

type thriftBinding struct {
	service.ResumeService
}

func (tb thriftBinding) Ping() (string, error) {
	r := tb.ResumeService.Ping()
	return r, nil
}

func (tb thriftBinding) AddResume(resume string) (string, error) {
	r := tb.ResumeService.AddResume(resume)
	return r, nil
}

func (tb thriftBinding) UpdateResume(userid string, resume string) (string, error) {
	r := tb.ResumeService.UpdateResume(userid, resume)
	return r, nil
}

func (tb thriftBinding) UpdateResumeBase(userid string, mmap map[string]string) (string, error) {
	r := tb.ResumeService.UpdateResumeBase(userid, mmap)
	return r, nil
}

func (tb thriftBinding) UpdateResumeSkillExperience(userid string, experience_levels string) (string, error) {
	r := tb.ResumeService.UpdateResumeSkillExperience(userid, experience_levels)
	return r, nil
}

func (tb thriftBinding) UpdateResumeToolandArchs(userid string, tool_archs string) (string, error) {
	r := tb.ResumeService.UpdateResumeToolandArchs(userid, tool_archs)
	return r, nil
}

func (tb thriftBinding) UpdateResumePortfolioes(userid string, portfolioes string) (string, error) {
	r := tb.ResumeService.UpdateResumePortfolioes(userid, portfolioes)
	return r, nil
}

func (tb thriftBinding) UpdateResumeEmploymentHistories(userid string, employment_histories string) (string, error) {
	r := tb.ResumeService.UpdateResumeEmploymentHistories(userid, employment_histories)
	return r, nil
}

func (tb thriftBinding) UpdateResumeEducations(userid string, educations string) (string, error) {
	r := tb.ResumeService.UpdateResumeEducations(userid, educations)
	return r, nil
}

func (tb thriftBinding) UpdateResumeOtherExperiences(userid string, other_experiences string) (string, error) {
	r := tb.ResumeService.UpdateResumeOtherExperiences(userid, other_experiences)
	return r, nil
}
