
service ResumeService {
    string Ping()
	string AddResume(1: string json_resume)
	string UpdateResume(1: string json_resume)
}
