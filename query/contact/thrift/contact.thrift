
service ContactService {
	string Ping()
	string GetContactTpl(1: string tplname)
	string GetContact(1: string contactid)
	string GetContactSignStatus(1: string contactid)		
	string GetClientContact(1: string clientemail)
	string GetFreelancerContact(1: string freelanceremail)
}
