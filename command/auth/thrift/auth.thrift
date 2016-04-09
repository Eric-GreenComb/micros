
service AuthService {
	string Register(1: string email,2: string pwd,3: string fromUserId)
	string Login(1: string emailOrUsername,2: string pwd)
}
