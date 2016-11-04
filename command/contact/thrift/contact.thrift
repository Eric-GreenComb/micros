
service ContactService {
    string Ping()
	string CreateContact(1: string jsonContact)    
	string ClientSignContact(1: string contactID,2: bool status)
	string FreelancerSignContact(1: string contactID,2: bool status)	
	string DealContact(1: string contactID,2: bool status)	
	string UpdateContact(1: string contactID, 2: map<string,string> mmap)	
}
