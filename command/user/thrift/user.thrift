
service UserService {
	string CreateUser(1: string email,2: string usernameraw,3: string pwd)
	bool UpdatePwd(1: string email,2: string oldpwd,3: string newpwd)
	bool ActiveUser(1: string token)
	i64 CountUser()
}
