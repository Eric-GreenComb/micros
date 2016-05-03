
service UserService {
    string Ping()
	string CreateUser(1: map<string,string> mmap)
	bool ResetPwd(1: string email,2: string newpwd)	
	bool ActiveUser(1: string email)
}
