
service ResumeService {
    string Ping()
	string AddResume(1: string resume)
	string UpdateResume(1: string userid, 2: string resume)
	string UpdateResumeBase(1: string userid, 2: map<string,string> mmap)
	string UpdateResumeSkillExperience(1: string userid, 2: string experienceLevels)
	string UpdateResumeToolandArchs(1: string userid, 2: string toolArchs)
	string UpdateResumePortfolioes(1: string userid, 2: string portfolioes)
	string UpdateResumeEmploymentHistories(1: string userid, 2: string employmentHistories)
	string UpdateResumeEducations(1: string userid, 2: string educations)
	string UpdateResumeOtherExperiences(1: string userid, 2: string otherExperiences)
}
