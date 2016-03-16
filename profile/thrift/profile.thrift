
service ProfileService {
	string GetProfile(1: string profile_id)
	string GetProfileByCat(1: string name)	
	string GetProfileBySubCat(1: string name)		
}
