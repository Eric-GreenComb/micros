package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/resume/service"
)

type instrumentingMiddleware struct {
	service.ResumeService
	requestDuration metrics.TimeHistogram
}

func (m instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "Ping"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.Ping()
	return
}

func (m instrumentingMiddleware) AddResume(resume string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "AddResume"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.AddResume(resume)
	return
}

func (m instrumentingMiddleware) UpdateResume(userid string, resume string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResume"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResume(userid, resume)
	return
}

func (m instrumentingMiddleware) UpdateResumeBase(userid string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResumeBase"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResumeBase(userid, mmap)
	return
}

func (m instrumentingMiddleware) UpdateResumeSkillExperience(userid string, experience_levels string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResumeSkillExperience"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResumeSkillExperience(userid, experience_levels)
	return
}

func (m instrumentingMiddleware) UpdateResumeToolandArchs(userid string, tool_archs string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResumeToolandArchs"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResumeToolandArchs(userid, tool_archs)
	return
}

func (m instrumentingMiddleware) UpdateResumePortfolioes(userid string, portfolioes string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResumePortfolioes"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResumePortfolioes(userid, portfolioes)
	return
}

func (m instrumentingMiddleware) UpdateResumeEmploymentHistories(userid string, employment_histories string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResumeEmploymentHistories"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResumeEmploymentHistories(userid, employment_histories)
	return
}

func (m instrumentingMiddleware) UpdateResumeEducations(userid string, educations string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResumeEducations"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResumeEducations(userid, educations)
	return
}

func (m instrumentingMiddleware) UpdateResumeOtherExperiences(userid string, other_experiences string) (r string) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "UpdateResumeOtherExperiences"}
		m.requestDuration.With(methodField).Observe(time.Since(begin))
	}(time.Now())
	r = m.ResumeService.UpdateResumeOtherExperiences(userid, other_experiences)
	return
}
