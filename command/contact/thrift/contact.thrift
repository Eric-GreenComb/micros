
service ContactService {
    string Ping()
	string CreateContact(1: string json_contact)    
	string ClientSignContact(1: string contact_id,2: bool status)
	string FreelancerSignContact(1: string contact_id,2: bool status)	
	string DealContact(1: string contact_id,2: bool status)	
	string UpdateContact(1: string contact_id, 2: map<string,string> mmap)	
}
