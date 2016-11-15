
service TokenService {
    string Ping()
	string CreateToken(1: string key,2: i64 ttype)
	bool DeleteToken(1: string key,2: i64 ttype)	
}
