
service ProfileService {
    string Ping()
	string AddProfile(1: string jsonProfile)    
	string UpdateProfile(1: string profileID,2: string jsonProfile)
	string UpdateProfileStatus(1: string profileID,2: bool status)	
	string UpdateProfileBase(1: string profileID, 2: map<string,string> mmap)	
	string UpdateProfileAgencyMembers(1: string profileID, 2: string agencyMembers)
}
