
service AuthService {
    string Ping()
	string Login(1: string emailOrUsername,2: string pwd)
}
