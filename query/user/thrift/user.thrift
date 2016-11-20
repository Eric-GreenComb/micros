
service UserService {
    string Ping()
	string GetUserByEmail(1: string email)
	string GetUserByID(1: string ID)	
	i64 CountUser()
}
