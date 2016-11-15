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

func (tb thriftBinding) UpdateResumeSkillExperience(userid string, experienceLevels string) (string, error) {
	r := tb.ResumeService.UpdateResumeSkillExperience(userid, experienceLevels)
	return r, nil
}

func (tb thriftBinding) UpdateResumeToolandArchs(userid string, toolArchs string) (string, error) {
	r := tb.ResumeService.UpdateResumeToolandArchs(userid, toolArchs)
	return r, nil
}

func (tb thriftBinding) UpdateResumePortfolioes(userid string, portfolioes string) (string, error) {
	r := tb.ResumeService.UpdateResumePortfolioes(userid, portfolioes)
	return r, nil
}

func (tb thriftBinding) UpdateResumeEmploymentHistories(userid string, employmentHistories string) (string, error) {
	r := tb.ResumeService.UpdateResumeEmploymentHistories(userid, employmentHistories)
	return r, nil
}

func (tb thriftBinding) UpdateResumeEducations(userid string, educations string) (string, error) {
	r := tb.ResumeService.UpdateResumeEducations(userid, educations)
	return r, nil
}

func (tb thriftBinding) UpdateResumeOtherExperiences(userid string, otherExperiences string) (string, error) {
	r := tb.ResumeService.UpdateResumeOtherExperiences(userid, otherExperiences)
	return r, nil
}
