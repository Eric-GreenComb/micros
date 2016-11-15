package main

import (
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/banerwai/micros/command/resume/service"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           service.ResumeService
}

func (mw instrumentingMiddleware) Ping() (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Ping", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.Ping()
	return
}

func (mw instrumentingMiddleware) AddResume(resume string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddResume", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.AddResume(resume)
	return
}

func (mw instrumentingMiddleware) UpdateResume(userID string, resume string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResume", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResume(userID, resume)
	return
}

func (mw instrumentingMiddleware) UpdateResumeBase(userID string, mmap map[string]string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResumeBase", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResumeBase(userID, mmap)
	return
}

func (mw instrumentingMiddleware) UpdateResumeSkillExperience(userID string, experienceLevels string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResumeSkillExperience", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResumeSkillExperience(userID, experienceLevels)
	return
}

func (mw instrumentingMiddleware) UpdateResumeToolandArchs(userID string, toolArchs string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResumeSkillExperience", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResumeToolandArchs(userID, toolArchs)
	return
}

func (mw instrumentingMiddleware) UpdateResumePortfolioes(userID string, portfolioes string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResumeSkillExperience", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResumePortfolioes(userID, portfolioes)
	return
}

func (mw instrumentingMiddleware) UpdateResumeEmploymentHistories(userID string, employmentHistories string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResumeEmploymentHistories", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResumeEmploymentHistories(userID, employmentHistories)
	return
}

func (mw instrumentingMiddleware) UpdateResumeEducations(userID string, educations string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResumeEducations", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResumeEducations(userID, educations)
	return
}

func (mw instrumentingMiddleware) UpdateResumeOtherExperiences(userID string, otherExperiences string) (r string) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateResumeOtherExperiences", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	r = mw.next.UpdateResumeOtherExperiences(userID, otherExperiences)
	return
}
