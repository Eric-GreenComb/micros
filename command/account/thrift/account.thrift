
service AccountService {
    string Ping()
	string CreateAccount(1: string json_account) 

	string CreateBilling(1: string json_billing)
	string DealBilling(1: string billing_id) 
	string CancelBilling(1: string billing_id)
	
	string GenAccount(1: string user_id)		  
}
