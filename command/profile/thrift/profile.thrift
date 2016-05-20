
service ProfileService {
    string Ping()
	string AddProfile(1: string json_profile)    
	string UpdateProfile(1: string profile_id,2: string json_profile)
	string UpdateProfileStatus(1: string profile_id,2: bool status)	
	string UpdateProfileBase(1: string profile_id, 2: map<string,string> mmap)	
	string UpdateProfileAgencyMembers(1: string profile_id, 2: string agency_members)
}
