
service ProfileService {
    string Ping()
	string GetProfile(1: string profileID)
	string GetProfilesByUserID(1: string userID)
	string GetProfilesByCategory(1: i64 categoryID, 2: i64 timestamp, 3: i64 pagesize)
	string GetProfilesBySubCategory(1: i64 subcategoryID, 2: i64 timestamp, 3: i64 pagesize)
	string SearchProfiles(1: map<string,i64> optionMap, 2: map<string,string> keyMap, 3: i64 timestamp, 4: i64 pagesize)		
}
