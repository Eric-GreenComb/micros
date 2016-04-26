
service ProfileService {
	string GetProfile(1: string id)
	string GetProfilesByEmail(1: string email)	
	string SearchProfiles(1: map<string,i64> option_mmap 2: map<string,string> key_mmap, 3: i64 timestamp, 4: i64 pagesize)		
}
