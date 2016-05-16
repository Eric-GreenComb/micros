package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/command/resume/service"
	thriftresume "github.com/banerwai/micros/command/resume/thrift/gen-go/resume"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftresume.ResumeServiceClient, logger log.Logger) service.ResumeService {
	return &client{cli, logger}
}

type client struct {
	*thriftresume.ResumeServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.ResumeServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) AddResume(resume string) string {
	reply, err := c.ResumeServiceClient.AddResume(resume)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResume(userid string, resume string) string {
	reply, err := c.ResumeServiceClient.UpdateResume(userid, resume)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResumeBase(userid string, mmap map[string]string) string {
	reply, err := c.ResumeServiceClient.UpdateResumeBase(userid, mmap)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResumeSkillExperience(userid string, experience_levels string) string {
	reply, err := c.ResumeServiceClient.UpdateResumeSkillExperience(userid, experience_levels)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResumeToolandArchs(userid string, tool_archs string) string {
	reply, err := c.ResumeServiceClient.UpdateResumeToolandArchs(userid, tool_archs)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResumePortfolioes(userid string, portfolioes string) string {
	reply, err := c.ResumeServiceClient.UpdateResumePortfolioes(userid, portfolioes)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResumeEmploymentHistories(userid string, employment_histories string) string {
	reply, err := c.ResumeServiceClient.UpdateResumeEmploymentHistories(userid, employment_histories)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResumeEducations(userid string, educations string) string {
	reply, err := c.ResumeServiceClient.UpdateResumeEducations(userid, educations)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) UpdateResumeOtherExperiences(userid string, other_experiences string) string {
	reply, err := c.ResumeServiceClient.UpdateResumeOtherExperiences(userid, other_experiences)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}
