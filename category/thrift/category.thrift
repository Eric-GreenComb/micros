struct Category{
	1: string id
	2: i32 serialnumber
	3: string name
	4: string desc
	5: list<SubCategory> subcategories
}

struct SubCategory{
	1: string id
	2: i32 serialnumber
	3: string name
	4: string desc
}

service CategoryService {
	string SayHi(1: string name)
	SubCategory GetDemoSubCategory(1: string id)	
	list<SubCategory> GetDemoSubCategories(1: string category_id)

	bool LoadCategory(1: string path)
	list<Category> GetCategories()		
	list<SubCategory> GetSubCategories(1: i32 serialnumber)			
}
