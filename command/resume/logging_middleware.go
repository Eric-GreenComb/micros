package main

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/resume/service"
)

type loggingMiddleware struct {
	service.ResumeService
	log.Logger
}

func (m loggingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "Ping",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.Ping()
	return
}

func (m loggingMiddleware) AddResume(resume string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "AddResume",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.AddResume(resume)
	return
}

func (m loggingMiddleware) UpdateResume(userid string, resume string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResume",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResume(userid, resume)
	return
}

func (m loggingMiddleware) UpdateResumeBase(userid string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResumeBase",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResumeBase(userid, mmap)
	return
}

func (m loggingMiddleware) UpdateResumeSkillExperience(userid string, experience_levels string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResumeSkillExperience",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResumeSkillExperience(userid, experience_levels)
	return
}

func (m loggingMiddleware) UpdateResumeToolandArchs(userid string, tool_archs string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResumeToolandArchs",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResumeToolandArchs(userid, tool_archs)
	return
}

func (m loggingMiddleware) UpdateResumePortfolioes(userid string, portfolioes string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResumePortfolioes",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResumePortfolioes(userid, portfolioes)
	return
}

func (m loggingMiddleware) UpdateResumeEmploymentHistories(userid string, employment_histories string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResumeEmploymentHistories",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResumeEmploymentHistories(userid, employment_histories)
	return
}

func (m loggingMiddleware) UpdateResumeEducations(userid string, educations string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResumeEducations",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResumeEducations(userid, educations)
	return
}

func (m loggingMiddleware) UpdateResumeOtherExperiences(userid string, other_experiences string) (r string) {
	defer func(begin time.Time) {
		m.Logger.Log(
			"method", "UpdateResume",
			"r", r,
			"took", time.Since(begin),
		)
	}(time.Now())
	r = m.ResumeService.UpdateResumeOtherExperiences(userid, other_experiences)
	return
}
