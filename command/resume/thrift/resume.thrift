
service ResumeService {
    string Ping()
	string AddResume(1: string resume)
	string UpdateResume(1: string userid, 2: string resume)
	string UpdateResumeBase(1: string userid, 2: map<string,string> mmap)
	string UpdateResumeSkillExperience(1: string userid, 2: string experience_levels)
	string UpdateResumeToolandArchs(1: string userid, 2: string tool_archs)
	string UpdateResumePortfolioes(1: string userid, 2: string portfolioes)
	string UpdateResumeEmploymentHistories(1: string userid, 2: string employment_histories)
	string UpdateResumeEducations(1: string userid, 2: string educations)
	string UpdateResumeOtherExperiences(1: string userid, 2: string other_experiences)
}
