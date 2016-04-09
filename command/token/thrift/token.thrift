
service TokenService {
	string NewToken(1: string key,2: i64 ttype)
	bool DeleteToken(1: string key,2: i64 ttype)	
	i64 VerifyToken(1: string token,2: i64 ttype)
}
