package main

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/micros/command/resume/service"
	"gopkg.in/mgo.v2/bson"
)

type inmemService struct {
}

func newInmemService() service.ResumeService {
	return &inmemService{}
}

func (ims *inmemService) Ping() (r string) {
	r = "pong"
	return
}

func (ims *inmemService) AddResume(resume string) (r string) {
	var _resume bean.Resume
	err := json.Unmarshal([]byte(resume), &_resume)
	if err != nil {
		return err.Error()
	}
	_resume.ID = bson.NewObjectId()
	_err := ResumeCollection.Insert(_resume)
	if _err != nil {
		return _err.Error()
	}
	return _resume.ID.Hex()
}

func (ims *inmemService) UpdateResume(userid string, resume string) (r string) {
	var _resume bean.Resume
	err := json.Unmarshal([]byte(resume), &_resume)
	if err != nil {
		return err.Error()
	}
	_resume.ID = ""
	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": _resume})
	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (ims *inmemService) UpdateResumeBase(userid string, mmap map[string]string) (r string) {
	_mongoM := bson.M{}

	for k, v := range mmap {
		_mongoM[k] = v
	}

	_, _err := ResumeCollection.Upsert(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": _mongoM})
	if nil != _err {
		r = _err.Error()
	}

	return "OK"
}

func (ims *inmemService) UpdateResumeSkillExperience(userid string, experienceLevels string) (r string) {
	var _beans []bean.SkillExperience
	err := json.Unmarshal([]byte(experienceLevels), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"experience_levels": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (ims *inmemService) UpdateResumeToolandArchs(userid string, toolArchs string) (r string) {
	var _beans []bean.ToolandArch
	err := json.Unmarshal([]byte(toolArchs), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"tool_archs": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (ims *inmemService) UpdateResumePortfolioes(userid string, portfolioes string) (r string) {
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

func (ims *inmemService) UpdateResumeEmploymentHistories(userid string, employmentHistories string) (r string) {
	var _beans []bean.EmploymentHistory
	err := json.Unmarshal([]byte(employmentHistories), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"employment_histories": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}

func (ims *inmemService) UpdateResumeEducations(userid string, educations string) (r string) {
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

func (ims *inmemService) UpdateResumeOtherExperiences(userid string, otherExperiences string) (r string) {
	var _beans []bean.OtherExperience
	err := json.Unmarshal([]byte(otherExperiences), &_beans)
	if err != nil {
		return err.Error()
	}

	_err := ResumeCollection.Update(bson.M{"userid": bson.ObjectIdHex(userid)}, bson.M{"$set": bson.M{"other_experiences": _beans}})

	if _err != nil {
		return _err.Error()
	}
	return "OK"
}
