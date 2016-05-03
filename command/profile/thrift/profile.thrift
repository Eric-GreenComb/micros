
service ProfileService {
    string Ping()
	string AddProfile(1: string json_profile)
	string UpdateProfile(1: string json_profile)
	string DeleteProfile(1: string id)				
}
