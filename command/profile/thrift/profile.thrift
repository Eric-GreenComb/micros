
service ProfileService {
	string AddProfile(1: string json_profile)
	string UpdateProfile(1: string json_profile)
	string DeleteProfile(1: string id)				
}
