struct Category{
	1: string id
	2: string name
	3: string desc
	4: list<SubCategory> subcategories
}

struct SubCategory{
	1: string id
	2: string name
	3: string desc
}

service CategoryService {
	string SayHi(1: string name)
	SubCategory GetDemoSubCategory(1: string id)	
	list<SubCategory> GetDemoSubCategories(1: string category_id)

	bool LoadCategory(1: string path)
	list<Category> GetCategories()		
	list<SubCategory> GetSubCategories(1: string category_id)			
}
