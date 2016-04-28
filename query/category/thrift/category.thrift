
service CategoryService {
	string Ping()
	string SayHi(1: string name)
	string GetDemoSubCategory(1: string id)	
	string GetDemoSubCategories(1: string category_id)

	bool LoadCategory(1: string path)
	string GetCategories()		
	string GetSubCategories(1: i32 serialnumber)			
}
