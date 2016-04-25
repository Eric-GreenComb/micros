
service TokenService {
	string NewToken(1: string key,2: i64 ttype)
	bool DeleteToken(1: string key,2: i64 ttype)	
}
