
service AccountService {
    string Ping()
	string CreateAccount(1: string jsonAccount) 

	string CreateBilling(1: string jsonBilling)
	string DealBilling(1: string billingID) 
	string CancelBilling(1: string billingID)
	
	string GenAccount(1: string userID)		  
}
