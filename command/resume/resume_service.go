package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/command/resume/service"
	"labix.org/v2/mgo/bson"
)

type inmemService struct {
}

func newInmemService() service.ResumeService {
	return &inmemService{}
}

func (self *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (self *inmemService) AddResume(resume string) (r string) {
	var _resume bean.Resume
	err := json.Unmarshal([]byte(resume), &_resume)
	if err != nil {
		return err.Error()
	}
	_resume.Id = ""
	_err := ResumeCollection.Insert(_resume)
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateResume(userid string, resume string) (r string) {
	var _resume bean.Resume
	err := json.Unmarshal([]byte(resume), &_resume)
	if err != nil {
		return err.Error()
	}
	_resume.Id = ""
	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": _resume})
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateResumeBase(userid string, mmap map[string]string) (r string) {
	_mongo_m := bson.M{}

	for k, v := range mmap {
		_mongo_m[k] = v
	}

	_, _err := ResumeCollection.Upsert(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": _mongo_m})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (self *inmemService) UpdateResumeSkillExperience(userid string, experience_levels string) (r string) {
	var _beans []bean.SkillExperience
	err := json.Unmarshal([]byte(experience_levels), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"experience_levels": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateResumeToolandArchs(userid string, tool_archs string) (r string) {
	var _beans []bean.ToolandArch
	err := json.Unmarshal([]byte(tool_archs), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"tool_archs": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateResumePortfolioes(userid string, portfolioes string) (r string) {
	var _beans []bean.Portfolio
	err := json.Unmarshal([]byte(portfolioes), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"portfolioes": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateResumeEmploymentHistories(userid string, employment_histories string) (r string) {
	var _beans []bean.EmploymentHistory
	err := json.Unmarshal([]byte(employment_histories), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"employment_histories": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateResumeEducations(userid string, educations string) (r string) {
	var _beans []bean.Education
	err := json.Unmarshal([]byte(educations), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"educations": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (self *inmemService) UpdateResumeOtherExperiences(userid string, other_experiences string) (r string) {
	var _beans []bean.OtherExperience
	err := json.Unmarshal([]byte(other_experiences), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"other_experiences": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}
