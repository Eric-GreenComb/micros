
service ProfileService {
	string GetProfile(1: string id)
	string SearchProfiles(1: string json_search, 2: i64 timestamp, 3: i64 pagesize)		
}
