
service TokenService {
    string Ping()
	i64 VerifyToken(1: string token,2: i64 ttype,3: i64 overhour)
}
