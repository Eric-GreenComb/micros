
service CategoryService {
	string Ping()
	bool LoadCategory(1: string path)
	string GetCategories()		
	string GetSubCategories(1: i32 serialnumber)			
}
