
service ProfileService {
    string Ping()
	string GetProfile(1: string profile_id)
	string GetProfilesByUserId(1: string user_id)	
	string SearchProfiles(1: map<string,i64> option_mmap 2: map<string,string> key_mmap, 3: i64 timestamp, 4: i64 pagesize)		
}
