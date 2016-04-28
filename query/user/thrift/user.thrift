
service UserService {
    string Ping()
	string GetUser(1: string email)
	i64 CountUser()
}
